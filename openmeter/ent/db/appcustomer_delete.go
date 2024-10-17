// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/appcustomer"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// AppCustomerDelete is the builder for deleting a AppCustomer entity.
type AppCustomerDelete struct {
	config
	hooks    []Hook
	mutation *AppCustomerMutation
}

// Where appends a list predicates to the AppCustomerDelete builder.
func (acd *AppCustomerDelete) Where(ps ...predicate.AppCustomer) *AppCustomerDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AppCustomerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AppCustomerDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AppCustomerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appcustomer.Table, sqlgraph.NewFieldSpec(appcustomer.FieldID, field.TypeInt))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AppCustomerDeleteOne is the builder for deleting a single AppCustomer entity.
type AppCustomerDeleteOne struct {
	acd *AppCustomerDelete
}

// Where appends a list predicates to the AppCustomerDelete builder.
func (acdo *AppCustomerDeleteOne) Where(ps ...predicate.AppCustomer) *AppCustomerDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AppCustomerDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appcustomer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AppCustomerDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}