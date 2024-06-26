// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent1/db/example1"
	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent1/db/predicate"
)

// Example1Delete is the builder for deleting a Example1 entity.
type Example1Delete struct {
	config
	hooks    []Hook
	mutation *Example1Mutation
}

// Where appends a list predicates to the Example1Delete builder.
func (e *Example1Delete) Where(ps ...predicate.Example1) *Example1Delete {
	e.mutation.Where(ps...)
	return e
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (e *Example1Delete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, e.sqlExec, e.mutation, e.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (e *Example1Delete) ExecX(ctx context.Context) int {
	n, err := e.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (e *Example1Delete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(example1.Table, sqlgraph.NewFieldSpec(example1.FieldID, field.TypeString))
	if ps := e.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, e.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	e.mutation.done = true
	return affected, err
}

// Example1DeleteOne is the builder for deleting a single Example1 entity.
type Example1DeleteOne struct {
	e *Example1Delete
}

// Where appends a list predicates to the Example1Delete builder.
func (eo *Example1DeleteOne) Where(ps ...predicate.Example1) *Example1DeleteOne {
	eo.e.mutation.Where(ps...)
	return eo
}

// Exec executes the deletion query.
func (eo *Example1DeleteOne) Exec(ctx context.Context) error {
	n, err := eo.e.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{example1.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eo *Example1DeleteOne) ExecX(ctx context.Context) {
	if err := eo.Exec(ctx); err != nil {
		panic(err)
	}
}
