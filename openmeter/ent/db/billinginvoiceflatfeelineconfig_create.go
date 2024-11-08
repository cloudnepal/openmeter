// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/alpacahq/alpacadecimal"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceflatfeelineconfig"
)

// BillingInvoiceFlatFeeLineConfigCreate is the builder for creating a BillingInvoiceFlatFeeLineConfig entity.
type BillingInvoiceFlatFeeLineConfigCreate struct {
	config
	mutation *BillingInvoiceFlatFeeLineConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) SetNamespace(s string) *BillingInvoiceFlatFeeLineConfigCreate {
	bifflcc.mutation.SetNamespace(s)
	return bifflcc
}

// SetAmount sets the "amount" field.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) SetAmount(a alpacadecimal.Decimal) *BillingInvoiceFlatFeeLineConfigCreate {
	bifflcc.mutation.SetAmount(a)
	return bifflcc
}

// SetID sets the "id" field.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) SetID(s string) *BillingInvoiceFlatFeeLineConfigCreate {
	bifflcc.mutation.SetID(s)
	return bifflcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) SetNillableID(s *string) *BillingInvoiceFlatFeeLineConfigCreate {
	if s != nil {
		bifflcc.SetID(*s)
	}
	return bifflcc
}

// Mutation returns the BillingInvoiceFlatFeeLineConfigMutation object of the builder.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) Mutation() *BillingInvoiceFlatFeeLineConfigMutation {
	return bifflcc.mutation
}

// Save creates the BillingInvoiceFlatFeeLineConfig in the database.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) Save(ctx context.Context) (*BillingInvoiceFlatFeeLineConfig, error) {
	bifflcc.defaults()
	return withHooks(ctx, bifflcc.sqlSave, bifflcc.mutation, bifflcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) SaveX(ctx context.Context) *BillingInvoiceFlatFeeLineConfig {
	v, err := bifflcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) Exec(ctx context.Context) error {
	_, err := bifflcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) ExecX(ctx context.Context) {
	if err := bifflcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) defaults() {
	if _, ok := bifflcc.mutation.ID(); !ok {
		v := billinginvoiceflatfeelineconfig.DefaultID()
		bifflcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) check() error {
	if _, ok := bifflcc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "BillingInvoiceFlatFeeLineConfig.namespace"`)}
	}
	if v, ok := bifflcc.mutation.Namespace(); ok {
		if err := billinginvoiceflatfeelineconfig.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "BillingInvoiceFlatFeeLineConfig.namespace": %w`, err)}
		}
	}
	if _, ok := bifflcc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`db: missing required field "BillingInvoiceFlatFeeLineConfig.amount"`)}
	}
	return nil
}

func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) sqlSave(ctx context.Context) (*BillingInvoiceFlatFeeLineConfig, error) {
	if err := bifflcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bifflcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bifflcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BillingInvoiceFlatFeeLineConfig.ID type: %T", _spec.ID.Value)
		}
	}
	bifflcc.mutation.id = &_node.ID
	bifflcc.mutation.done = true
	return _node, nil
}

func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) createSpec() (*BillingInvoiceFlatFeeLineConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &BillingInvoiceFlatFeeLineConfig{config: bifflcc.config}
		_spec = sqlgraph.NewCreateSpec(billinginvoiceflatfeelineconfig.Table, sqlgraph.NewFieldSpec(billinginvoiceflatfeelineconfig.FieldID, field.TypeString))
	)
	_spec.OnConflict = bifflcc.conflict
	if id, ok := bifflcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bifflcc.mutation.Namespace(); ok {
		_spec.SetField(billinginvoiceflatfeelineconfig.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := bifflcc.mutation.Amount(); ok {
		_spec.SetField(billinginvoiceflatfeelineconfig.FieldAmount, field.TypeOther, value)
		_node.Amount = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingInvoiceFlatFeeLineConfigUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) OnConflict(opts ...sql.ConflictOption) *BillingInvoiceFlatFeeLineConfigUpsertOne {
	bifflcc.conflict = opts
	return &BillingInvoiceFlatFeeLineConfigUpsertOne{
		create: bifflcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bifflcc *BillingInvoiceFlatFeeLineConfigCreate) OnConflictColumns(columns ...string) *BillingInvoiceFlatFeeLineConfigUpsertOne {
	bifflcc.conflict = append(bifflcc.conflict, sql.ConflictColumns(columns...))
	return &BillingInvoiceFlatFeeLineConfigUpsertOne{
		create: bifflcc,
	}
}

type (
	// BillingInvoiceFlatFeeLineConfigUpsertOne is the builder for "upsert"-ing
	//  one BillingInvoiceFlatFeeLineConfig node.
	BillingInvoiceFlatFeeLineConfigUpsertOne struct {
		create *BillingInvoiceFlatFeeLineConfigCreate
	}

	// BillingInvoiceFlatFeeLineConfigUpsert is the "OnConflict" setter.
	BillingInvoiceFlatFeeLineConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetAmount sets the "amount" field.
func (u *BillingInvoiceFlatFeeLineConfigUpsert) SetAmount(v alpacadecimal.Decimal) *BillingInvoiceFlatFeeLineConfigUpsert {
	u.Set(billinginvoiceflatfeelineconfig.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *BillingInvoiceFlatFeeLineConfigUpsert) UpdateAmount() *BillingInvoiceFlatFeeLineConfigUpsert {
	u.SetExcluded(billinginvoiceflatfeelineconfig.FieldAmount)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billinginvoiceflatfeelineconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) UpdateNewValues() *BillingInvoiceFlatFeeLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(billinginvoiceflatfeelineconfig.FieldID)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(billinginvoiceflatfeelineconfig.FieldNamespace)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) Ignore() *BillingInvoiceFlatFeeLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) DoNothing() *BillingInvoiceFlatFeeLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingInvoiceFlatFeeLineConfigCreate.OnConflict
// documentation for more info.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) Update(set func(*BillingInvoiceFlatFeeLineConfigUpsert)) *BillingInvoiceFlatFeeLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingInvoiceFlatFeeLineConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) SetAmount(v alpacadecimal.Decimal) *BillingInvoiceFlatFeeLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceFlatFeeLineConfigUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) UpdateAmount() *BillingInvoiceFlatFeeLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceFlatFeeLineConfigUpsert) {
		s.UpdateAmount()
	})
}

// Exec executes the query.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingInvoiceFlatFeeLineConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: BillingInvoiceFlatFeeLineConfigUpsertOne.ID is not supported by MySQL driver. Use BillingInvoiceFlatFeeLineConfigUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BillingInvoiceFlatFeeLineConfigUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BillingInvoiceFlatFeeLineConfigCreateBulk is the builder for creating many BillingInvoiceFlatFeeLineConfig entities in bulk.
type BillingInvoiceFlatFeeLineConfigCreateBulk struct {
	config
	err      error
	builders []*BillingInvoiceFlatFeeLineConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the BillingInvoiceFlatFeeLineConfig entities in the database.
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) Save(ctx context.Context) ([]*BillingInvoiceFlatFeeLineConfig, error) {
	if bifflccb.err != nil {
		return nil, bifflccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bifflccb.builders))
	nodes := make([]*BillingInvoiceFlatFeeLineConfig, len(bifflccb.builders))
	mutators := make([]Mutator, len(bifflccb.builders))
	for i := range bifflccb.builders {
		func(i int, root context.Context) {
			builder := bifflccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingInvoiceFlatFeeLineConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bifflccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bifflccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bifflccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bifflccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) SaveX(ctx context.Context) []*BillingInvoiceFlatFeeLineConfig {
	v, err := bifflccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := bifflccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) ExecX(ctx context.Context) {
	if err := bifflccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingInvoiceFlatFeeLineConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingInvoiceFlatFeeLineConfigUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	bifflccb.conflict = opts
	return &BillingInvoiceFlatFeeLineConfigUpsertBulk{
		create: bifflccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bifflccb *BillingInvoiceFlatFeeLineConfigCreateBulk) OnConflictColumns(columns ...string) *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	bifflccb.conflict = append(bifflccb.conflict, sql.ConflictColumns(columns...))
	return &BillingInvoiceFlatFeeLineConfigUpsertBulk{
		create: bifflccb,
	}
}

// BillingInvoiceFlatFeeLineConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of BillingInvoiceFlatFeeLineConfig nodes.
type BillingInvoiceFlatFeeLineConfigUpsertBulk struct {
	create *BillingInvoiceFlatFeeLineConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billinginvoiceflatfeelineconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) UpdateNewValues() *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(billinginvoiceflatfeelineconfig.FieldID)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(billinginvoiceflatfeelineconfig.FieldNamespace)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceFlatFeeLineConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) Ignore() *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) DoNothing() *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingInvoiceFlatFeeLineConfigCreateBulk.OnConflict
// documentation for more info.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) Update(set func(*BillingInvoiceFlatFeeLineConfigUpsert)) *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingInvoiceFlatFeeLineConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) SetAmount(v alpacadecimal.Decimal) *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceFlatFeeLineConfigUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) UpdateAmount() *BillingInvoiceFlatFeeLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceFlatFeeLineConfigUpsert) {
		s.UpdateAmount()
	})
}

// Exec executes the query.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the BillingInvoiceFlatFeeLineConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingInvoiceFlatFeeLineConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingInvoiceFlatFeeLineConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
