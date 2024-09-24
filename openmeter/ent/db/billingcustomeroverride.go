// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/openmeter/billing"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingcustomeroverride"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingprofile"
	"github.com/openmeterio/openmeter/openmeter/ent/db/customer"
)

// BillingCustomerOverride is the model entity for the BillingCustomerOverride schema.
type BillingCustomerOverride struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID string `json:"customer_id,omitempty"`
	// BillingProfileID holds the value of the "billing_profile_id" field.
	BillingProfileID *string `json:"billing_profile_id,omitempty"`
	// CollectionAlignment holds the value of the "collection_alignment" field.
	CollectionAlignment *billing.AlignmentKind `json:"collection_alignment,omitempty"`
	// ItemCollectionPeriodSeconds holds the value of the "item_collection_period_seconds" field.
	ItemCollectionPeriodSeconds *int64 `json:"item_collection_period_seconds,omitempty"`
	// InvoiceAutoAdvance holds the value of the "invoice_auto_advance" field.
	InvoiceAutoAdvance *bool `json:"invoice_auto_advance,omitempty"`
	// InvoiceDraftPeriodSeconds holds the value of the "invoice_draft_period_seconds" field.
	InvoiceDraftPeriodSeconds *int64 `json:"invoice_draft_period_seconds,omitempty"`
	// InvoiceDueAfterSeconds holds the value of the "invoice_due_after_seconds" field.
	InvoiceDueAfterSeconds *int64 `json:"invoice_due_after_seconds,omitempty"`
	// InvoiceCollectionMethod holds the value of the "invoice_collection_method" field.
	InvoiceCollectionMethod *billing.CollectionMethod `json:"invoice_collection_method,omitempty"`
	// InvoiceItemResolution holds the value of the "invoice_item_resolution" field.
	InvoiceItemResolution *billing.GranularityResolution `json:"invoice_item_resolution,omitempty"`
	// InvoiceItemPerSubject holds the value of the "invoice_item_per_subject" field.
	InvoiceItemPerSubject *bool `json:"invoice_item_per_subject,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillingCustomerOverrideQuery when eager-loading is set.
	Edges        BillingCustomerOverrideEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BillingCustomerOverrideEdges holds the relations/edges for other nodes in the graph.
type BillingCustomerOverrideEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// BillingProfile holds the value of the billing_profile edge.
	BillingProfile *BillingProfile `json:"billing_profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingCustomerOverrideEdges) CustomerOrErr() (*Customer, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: customer.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// BillingProfileOrErr returns the BillingProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingCustomerOverrideEdges) BillingProfileOrErr() (*BillingProfile, error) {
	if e.BillingProfile != nil {
		return e.BillingProfile, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: billingprofile.Label}
	}
	return nil, &NotLoadedError{edge: "billing_profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingCustomerOverride) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingcustomeroverride.FieldInvoiceAutoAdvance, billingcustomeroverride.FieldInvoiceItemPerSubject:
			values[i] = new(sql.NullBool)
		case billingcustomeroverride.FieldItemCollectionPeriodSeconds, billingcustomeroverride.FieldInvoiceDraftPeriodSeconds, billingcustomeroverride.FieldInvoiceDueAfterSeconds:
			values[i] = new(sql.NullInt64)
		case billingcustomeroverride.FieldID, billingcustomeroverride.FieldNamespace, billingcustomeroverride.FieldCustomerID, billingcustomeroverride.FieldBillingProfileID, billingcustomeroverride.FieldCollectionAlignment, billingcustomeroverride.FieldInvoiceCollectionMethod, billingcustomeroverride.FieldInvoiceItemResolution:
			values[i] = new(sql.NullString)
		case billingcustomeroverride.FieldCreatedAt, billingcustomeroverride.FieldUpdatedAt, billingcustomeroverride.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingCustomerOverride fields.
func (bco *BillingCustomerOverride) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingcustomeroverride.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bco.ID = value.String
			}
		case billingcustomeroverride.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				bco.Namespace = value.String
			}
		case billingcustomeroverride.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bco.CreatedAt = value.Time
			}
		case billingcustomeroverride.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bco.UpdatedAt = value.Time
			}
		case billingcustomeroverride.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				bco.DeletedAt = new(time.Time)
				*bco.DeletedAt = value.Time
			}
		case billingcustomeroverride.FieldCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value.Valid {
				bco.CustomerID = value.String
			}
		case billingcustomeroverride.FieldBillingProfileID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field billing_profile_id", values[i])
			} else if value.Valid {
				bco.BillingProfileID = new(string)
				*bco.BillingProfileID = value.String
			}
		case billingcustomeroverride.FieldCollectionAlignment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collection_alignment", values[i])
			} else if value.Valid {
				bco.CollectionAlignment = new(billing.AlignmentKind)
				*bco.CollectionAlignment = billing.AlignmentKind(value.String)
			}
		case billingcustomeroverride.FieldItemCollectionPeriodSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field item_collection_period_seconds", values[i])
			} else if value.Valid {
				bco.ItemCollectionPeriodSeconds = new(int64)
				*bco.ItemCollectionPeriodSeconds = value.Int64
			}
		case billingcustomeroverride.FieldInvoiceAutoAdvance:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_auto_advance", values[i])
			} else if value.Valid {
				bco.InvoiceAutoAdvance = new(bool)
				*bco.InvoiceAutoAdvance = value.Bool
			}
		case billingcustomeroverride.FieldInvoiceDraftPeriodSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_draft_period_seconds", values[i])
			} else if value.Valid {
				bco.InvoiceDraftPeriodSeconds = new(int64)
				*bco.InvoiceDraftPeriodSeconds = value.Int64
			}
		case billingcustomeroverride.FieldInvoiceDueAfterSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_due_after_seconds", values[i])
			} else if value.Valid {
				bco.InvoiceDueAfterSeconds = new(int64)
				*bco.InvoiceDueAfterSeconds = value.Int64
			}
		case billingcustomeroverride.FieldInvoiceCollectionMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_collection_method", values[i])
			} else if value.Valid {
				bco.InvoiceCollectionMethod = new(billing.CollectionMethod)
				*bco.InvoiceCollectionMethod = billing.CollectionMethod(value.String)
			}
		case billingcustomeroverride.FieldInvoiceItemResolution:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_item_resolution", values[i])
			} else if value.Valid {
				bco.InvoiceItemResolution = new(billing.GranularityResolution)
				*bco.InvoiceItemResolution = billing.GranularityResolution(value.String)
			}
		case billingcustomeroverride.FieldInvoiceItemPerSubject:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_item_per_subject", values[i])
			} else if value.Valid {
				bco.InvoiceItemPerSubject = new(bool)
				*bco.InvoiceItemPerSubject = value.Bool
			}
		default:
			bco.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingCustomerOverride.
// This includes values selected through modifiers, order, etc.
func (bco *BillingCustomerOverride) Value(name string) (ent.Value, error) {
	return bco.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the BillingCustomerOverride entity.
func (bco *BillingCustomerOverride) QueryCustomer() *CustomerQuery {
	return NewBillingCustomerOverrideClient(bco.config).QueryCustomer(bco)
}

// QueryBillingProfile queries the "billing_profile" edge of the BillingCustomerOverride entity.
func (bco *BillingCustomerOverride) QueryBillingProfile() *BillingProfileQuery {
	return NewBillingCustomerOverrideClient(bco.config).QueryBillingProfile(bco)
}

// Update returns a builder for updating this BillingCustomerOverride.
// Note that you need to call BillingCustomerOverride.Unwrap() before calling this method if this BillingCustomerOverride
// was returned from a transaction, and the transaction was committed or rolled back.
func (bco *BillingCustomerOverride) Update() *BillingCustomerOverrideUpdateOne {
	return NewBillingCustomerOverrideClient(bco.config).UpdateOne(bco)
}

// Unwrap unwraps the BillingCustomerOverride entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bco *BillingCustomerOverride) Unwrap() *BillingCustomerOverride {
	_tx, ok := bco.config.driver.(*txDriver)
	if !ok {
		panic("db: BillingCustomerOverride is not a transactional entity")
	}
	bco.config.driver = _tx.drv
	return bco
}

// String implements the fmt.Stringer.
func (bco *BillingCustomerOverride) String() string {
	var builder strings.Builder
	builder.WriteString("BillingCustomerOverride(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bco.ID))
	builder.WriteString("namespace=")
	builder.WriteString(bco.Namespace)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bco.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bco.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := bco.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(bco.CustomerID)
	builder.WriteString(", ")
	if v := bco.BillingProfileID; v != nil {
		builder.WriteString("billing_profile_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bco.CollectionAlignment; v != nil {
		builder.WriteString("collection_alignment=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.ItemCollectionPeriodSeconds; v != nil {
		builder.WriteString("item_collection_period_seconds=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceAutoAdvance; v != nil {
		builder.WriteString("invoice_auto_advance=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceDraftPeriodSeconds; v != nil {
		builder.WriteString("invoice_draft_period_seconds=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceDueAfterSeconds; v != nil {
		builder.WriteString("invoice_due_after_seconds=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceCollectionMethod; v != nil {
		builder.WriteString("invoice_collection_method=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceItemResolution; v != nil {
		builder.WriteString("invoice_item_resolution=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := bco.InvoiceItemPerSubject; v != nil {
		builder.WriteString("invoice_item_per_subject=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// BillingCustomerOverrides is a parsable slice of BillingCustomerOverride.
type BillingCustomerOverrides []*BillingCustomerOverride