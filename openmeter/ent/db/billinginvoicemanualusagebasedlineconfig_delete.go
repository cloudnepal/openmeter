// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoicemanualusagebasedlineconfig"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceManualUsageBasedLineConfigDelete is the builder for deleting a BillingInvoiceManualUsageBasedLineConfig entity.
type BillingInvoiceManualUsageBasedLineConfigDelete struct {
	config
	hooks    []Hook
	mutation *BillingInvoiceManualUsageBasedLineConfigMutation
}

// Where appends a list predicates to the BillingInvoiceManualUsageBasedLineConfigDelete builder.
func (bimublcd *BillingInvoiceManualUsageBasedLineConfigDelete) Where(ps ...predicate.BillingInvoiceManualUsageBasedLineConfig) *BillingInvoiceManualUsageBasedLineConfigDelete {
	bimublcd.mutation.Where(ps...)
	return bimublcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bimublcd *BillingInvoiceManualUsageBasedLineConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bimublcd.sqlExec, bimublcd.mutation, bimublcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bimublcd *BillingInvoiceManualUsageBasedLineConfigDelete) ExecX(ctx context.Context) int {
	n, err := bimublcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bimublcd *BillingInvoiceManualUsageBasedLineConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billinginvoicemanualusagebasedlineconfig.Table, sqlgraph.NewFieldSpec(billinginvoicemanualusagebasedlineconfig.FieldID, field.TypeString))
	if ps := bimublcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bimublcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bimublcd.mutation.done = true
	return affected, err
}

// BillingInvoiceManualUsageBasedLineConfigDeleteOne is the builder for deleting a single BillingInvoiceManualUsageBasedLineConfig entity.
type BillingInvoiceManualUsageBasedLineConfigDeleteOne struct {
	bimublcd *BillingInvoiceManualUsageBasedLineConfigDelete
}

// Where appends a list predicates to the BillingInvoiceManualUsageBasedLineConfigDelete builder.
func (bimublcdo *BillingInvoiceManualUsageBasedLineConfigDeleteOne) Where(ps ...predicate.BillingInvoiceManualUsageBasedLineConfig) *BillingInvoiceManualUsageBasedLineConfigDeleteOne {
	bimublcdo.bimublcd.mutation.Where(ps...)
	return bimublcdo
}

// Exec executes the deletion query.
func (bimublcdo *BillingInvoiceManualUsageBasedLineConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := bimublcdo.bimublcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billinginvoicemanualusagebasedlineconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bimublcdo *BillingInvoiceManualUsageBasedLineConfigDeleteOne) ExecX(ctx context.Context) {
	if err := bimublcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
