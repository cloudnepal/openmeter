package billingadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/alpacahq/alpacadecimal"
	"github.com/samber/lo"

	billingentity "github.com/openmeterio/openmeter/openmeter/billing/entity"
	"github.com/openmeterio/openmeter/openmeter/ent/db"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceline"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
	"github.com/openmeterio/openmeter/pkg/convert"
)

func (a *adapter) mapInvoiceLineFromDB(ctx context.Context, invoiceLines []*db.BillingInvoiceLine) ([]*billingentity.Line, error) {
	pendingParentIDs := make([]string, 0, len(invoiceLines))
	resolvedChildrenOfIDs := make(map[string]struct{}, len(invoiceLines))

	for _, line := range invoiceLines {
		if line.ParentLineID != nil {
			pendingParentIDs = append(pendingParentIDs, *line.ParentLineID)
		}

		if line.Status != billingentity.InvoiceLineStatusDetailed {
			resolvedChildrenOfIDs[line.ID] = struct{}{}
		}
	}

	// NOTE: Given that the invoice lines can be in parent-child relationship we might fetch
	// duplicate lines, so we need to deduplicate them.
	//
	// We cannot get around this limitation, as a parent line might have more children than the ones we have
	// saved.
	references, err := a.fetchInvoiceLineNewReferences(ctx, pendingParentIDs, lo.Keys(resolvedChildrenOfIDs))
	if err != nil {
		return nil, err
	}

	references = append(references, invoiceLines...)

	mappedEntities := make(map[string]*billingentity.Line, len(references))

	for _, dbLine := range references {
		if _, ok := mappedEntities[dbLine.ID]; ok {
			continue
		}

		entity, err := a.mapInvoiceLineWithoutReferences(dbLine)
		if err != nil {
			return nil, err
		}
		mappedEntities[dbLine.ID] = &entity
	}

	for _, entity := range mappedEntities {
		if entity.ParentLineID != nil {
			parent, ok := mappedEntities[*entity.ParentLineID]
			if !ok {
				// We don't care about the parent if it's not loaded as it might be too deep
				continue
			}

			entity.ParentLine = parent
			// We only add children references if we know that those has been properly resolved
			if _, ok := resolvedChildrenOfIDs[parent.ID]; ok {
				parent.Children.Append(entity)
			}
		}
	}

	result := make([]*billingentity.Line, 0, len(mappedEntities))
	for _, dbEntity := range invoiceLines {
		entity, ok := mappedEntities[dbEntity.ID]
		if !ok {
			return nil, fmt.Errorf("missing entity[%s]", dbEntity.ID)
		}

		// Given that we did the one level deep resolution, all the children should be resolved
		// if it's not present it means that the line has no children, but before we only rely on Append
		// to set the Present flag implicitly.
		if !entity.Children.IsPresent() {
			entity.Children = billingentity.NewLineChildren(nil)
		}

		entity.SaveDBSnapshot()

		result = append(result, entity)
	}

	return result, nil
}

func (a *adapter) fetchInvoiceLineNewReferences(ctx context.Context, parentIDs []string, childrenOf []string) ([]*db.BillingInvoiceLine, error) {
	if len(parentIDs) == 0 && len(childrenOf) == 0 {
		return nil, nil
	}

	query := a.db.BillingInvoiceLine.Query()

	query = a.expandLineItems(query)

	predicates := make([]predicate.BillingInvoiceLine, 0, 2)
	if len(parentIDs) > 0 {
		predicates = append(predicates, billinginvoiceline.IDIn(lo.Uniq(parentIDs)...))
	}

	if len(childrenOf) > 0 {
		predicates = append(predicates, billinginvoiceline.ParentLineIDIn(lo.Uniq(childrenOf)...))
	}

	query = query.Where(
		billinginvoiceline.Or(predicates...),
		billinginvoiceline.DeletedAtIsNil(),
	)

	return query.All(ctx)
}

func (a *adapter) mapInvoiceLineWithoutReferences(dbLine *db.BillingInvoiceLine) (billingentity.Line, error) {
	invoiceLine := billingentity.Line{
		LineBase: billingentity.LineBase{
			Namespace: dbLine.Namespace,
			ID:        dbLine.ID,

			CreatedAt: dbLine.CreatedAt.In(time.UTC),
			UpdatedAt: dbLine.UpdatedAt.In(time.UTC),
			DeletedAt: convert.TimePtrIn(dbLine.DeletedAt, time.UTC),

			Metadata:  dbLine.Metadata,
			InvoiceID: dbLine.InvoiceID,
			Status:    dbLine.Status,

			Period: billingentity.Period{
				Start: dbLine.PeriodStart.In(time.UTC),
				End:   dbLine.PeriodEnd.In(time.UTC),
			},

			ParentLineID:           dbLine.ParentLineID,
			ChildUniqueReferenceID: dbLine.ChildUniqueReferenceID,

			InvoiceAt: dbLine.InvoiceAt.In(time.UTC),

			Name:        dbLine.Name,
			Description: dbLine.Description,

			Type:     dbLine.Type,
			Currency: dbLine.Currency,

			TaxConfig: lo.EmptyableToPtr(dbLine.TaxConfig),
		},
	}

	switch dbLine.Type {
	case billingentity.InvoiceLineTypeFee:
		invoiceLine.FlatFee = billingentity.FlatFeeLine{
			ConfigID:      dbLine.Edges.FlatFeeLine.ID,
			PerUnitAmount: dbLine.Edges.FlatFeeLine.PerUnitAmount,
			Quantity:      lo.FromPtrOr(dbLine.Quantity, alpacadecimal.Zero),
		}
	case billingentity.InvoiceLineTypeUsageBased:
		ubpLine := dbLine.Edges.UsageBasedLine
		if ubpLine == nil {
			return invoiceLine, fmt.Errorf("manual usage based line is missing")
		}
		invoiceLine.UsageBased = billingentity.UsageBasedLine{
			ConfigID:   ubpLine.ID,
			FeatureKey: ubpLine.FeatureKey,
			Price:      *ubpLine.Price,
			Quantity:   dbLine.Quantity,
		}
	default:
		return invoiceLine, fmt.Errorf("unsupported line type[%s]: %s", dbLine.ID, dbLine.Type)
	}

	discounts := lo.Map(dbLine.Edges.LineDiscounts, func(discount *db.BillingInvoiceLineDiscount, _ int) billingentity.LineDiscount {
		return billingentity.LineDiscount{
			ID:        discount.ID,
			CreatedAt: discount.CreatedAt.In(time.UTC),
			UpdatedAt: discount.UpdatedAt.In(time.UTC),
			DeletedAt: convert.TimePtrIn(discount.DeletedAt, time.UTC),

			Amount:                 discount.Amount,
			Description:            discount.Description,
			ChildUniqueReferenceID: discount.ChildUniqueReferenceID,
		}
	})
	invoiceLine.Discounts = billingentity.NewLineDiscounts(discounts)

	return invoiceLine, nil
}
