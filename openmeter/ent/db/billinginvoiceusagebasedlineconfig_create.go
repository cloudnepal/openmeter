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
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceusagebasedlineconfig"
	"github.com/openmeterio/openmeter/openmeter/productcatalog/plan"
)

// BillingInvoiceUsageBasedLineConfigCreate is the builder for creating a BillingInvoiceUsageBasedLineConfig entity.
type BillingInvoiceUsageBasedLineConfigCreate struct {
	config
	mutation *BillingInvoiceUsageBasedLineConfigMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetNamespace(s string) *BillingInvoiceUsageBasedLineConfigCreate {
	biublcc.mutation.SetNamespace(s)
	return biublcc
}

// SetPriceType sets the "price_type" field.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetPriceType(pt plan.PriceType) *BillingInvoiceUsageBasedLineConfigCreate {
	biublcc.mutation.SetPriceType(pt)
	return biublcc
}

// SetFeatureKey sets the "feature_key" field.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetFeatureKey(s string) *BillingInvoiceUsageBasedLineConfigCreate {
	biublcc.mutation.SetFeatureKey(s)
	return biublcc
}

// SetPrice sets the "price" field.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetPrice(pl *plan.Price) *BillingInvoiceUsageBasedLineConfigCreate {
	biublcc.mutation.SetPrice(pl)
	return biublcc
}

// SetID sets the "id" field.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetID(s string) *BillingInvoiceUsageBasedLineConfigCreate {
	biublcc.mutation.SetID(s)
	return biublcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SetNillableID(s *string) *BillingInvoiceUsageBasedLineConfigCreate {
	if s != nil {
		biublcc.SetID(*s)
	}
	return biublcc
}

// Mutation returns the BillingInvoiceUsageBasedLineConfigMutation object of the builder.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) Mutation() *BillingInvoiceUsageBasedLineConfigMutation {
	return biublcc.mutation
}

// Save creates the BillingInvoiceUsageBasedLineConfig in the database.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) Save(ctx context.Context) (*BillingInvoiceUsageBasedLineConfig, error) {
	biublcc.defaults()
	return withHooks(ctx, biublcc.sqlSave, biublcc.mutation, biublcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) SaveX(ctx context.Context) *BillingInvoiceUsageBasedLineConfig {
	v, err := biublcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) Exec(ctx context.Context) error {
	_, err := biublcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) ExecX(ctx context.Context) {
	if err := biublcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) defaults() {
	if _, ok := biublcc.mutation.ID(); !ok {
		v := billinginvoiceusagebasedlineconfig.DefaultID()
		biublcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) check() error {
	if _, ok := biublcc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "BillingInvoiceUsageBasedLineConfig.namespace"`)}
	}
	if v, ok := biublcc.mutation.Namespace(); ok {
		if err := billinginvoiceusagebasedlineconfig.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "BillingInvoiceUsageBasedLineConfig.namespace": %w`, err)}
		}
	}
	if _, ok := biublcc.mutation.PriceType(); !ok {
		return &ValidationError{Name: "price_type", err: errors.New(`db: missing required field "BillingInvoiceUsageBasedLineConfig.price_type"`)}
	}
	if v, ok := biublcc.mutation.PriceType(); ok {
		if err := billinginvoiceusagebasedlineconfig.PriceTypeValidator(v); err != nil {
			return &ValidationError{Name: "price_type", err: fmt.Errorf(`db: validator failed for field "BillingInvoiceUsageBasedLineConfig.price_type": %w`, err)}
		}
	}
	if _, ok := biublcc.mutation.FeatureKey(); !ok {
		return &ValidationError{Name: "feature_key", err: errors.New(`db: missing required field "BillingInvoiceUsageBasedLineConfig.feature_key"`)}
	}
	if v, ok := biublcc.mutation.FeatureKey(); ok {
		if err := billinginvoiceusagebasedlineconfig.FeatureKeyValidator(v); err != nil {
			return &ValidationError{Name: "feature_key", err: fmt.Errorf(`db: validator failed for field "BillingInvoiceUsageBasedLineConfig.feature_key": %w`, err)}
		}
	}
	if _, ok := biublcc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`db: missing required field "BillingInvoiceUsageBasedLineConfig.price"`)}
	}
	if v, ok := biublcc.mutation.Price(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`db: validator failed for field "BillingInvoiceUsageBasedLineConfig.price": %w`, err)}
		}
	}
	return nil
}

func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) sqlSave(ctx context.Context) (*BillingInvoiceUsageBasedLineConfig, error) {
	if err := biublcc.check(); err != nil {
		return nil, err
	}
	_node, _spec, err := biublcc.createSpec()
	if err != nil {
		return nil, err
	}
	if err := sqlgraph.CreateNode(ctx, biublcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BillingInvoiceUsageBasedLineConfig.ID type: %T", _spec.ID.Value)
		}
	}
	biublcc.mutation.id = &_node.ID
	biublcc.mutation.done = true
	return _node, nil
}

func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) createSpec() (*BillingInvoiceUsageBasedLineConfig, *sqlgraph.CreateSpec, error) {
	var (
		_node = &BillingInvoiceUsageBasedLineConfig{config: biublcc.config}
		_spec = sqlgraph.NewCreateSpec(billinginvoiceusagebasedlineconfig.Table, sqlgraph.NewFieldSpec(billinginvoiceusagebasedlineconfig.FieldID, field.TypeString))
	)
	_spec.OnConflict = biublcc.conflict
	if id, ok := biublcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := biublcc.mutation.Namespace(); ok {
		_spec.SetField(billinginvoiceusagebasedlineconfig.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := biublcc.mutation.PriceType(); ok {
		_spec.SetField(billinginvoiceusagebasedlineconfig.FieldPriceType, field.TypeEnum, value)
		_node.PriceType = value
	}
	if value, ok := biublcc.mutation.FeatureKey(); ok {
		_spec.SetField(billinginvoiceusagebasedlineconfig.FieldFeatureKey, field.TypeString, value)
		_node.FeatureKey = value
	}
	if value, ok := biublcc.mutation.Price(); ok {
		vv, err := billinginvoiceusagebasedlineconfig.ValueScanner.Price.Value(value)
		if err != nil {
			return nil, nil, err
		}
		_spec.SetField(billinginvoiceusagebasedlineconfig.FieldPrice, field.TypeString, vv)
		_node.Price = value
	}
	return _node, _spec, nil
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingInvoiceUsageBasedLineConfigUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) OnConflict(opts ...sql.ConflictOption) *BillingInvoiceUsageBasedLineConfigUpsertOne {
	biublcc.conflict = opts
	return &BillingInvoiceUsageBasedLineConfigUpsertOne{
		create: biublcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (biublcc *BillingInvoiceUsageBasedLineConfigCreate) OnConflictColumns(columns ...string) *BillingInvoiceUsageBasedLineConfigUpsertOne {
	biublcc.conflict = append(biublcc.conflict, sql.ConflictColumns(columns...))
	return &BillingInvoiceUsageBasedLineConfigUpsertOne{
		create: biublcc,
	}
}

type (
	// BillingInvoiceUsageBasedLineConfigUpsertOne is the builder for "upsert"-ing
	//  one BillingInvoiceUsageBasedLineConfig node.
	BillingInvoiceUsageBasedLineConfigUpsertOne struct {
		create *BillingInvoiceUsageBasedLineConfigCreate
	}

	// BillingInvoiceUsageBasedLineConfigUpsert is the "OnConflict" setter.
	BillingInvoiceUsageBasedLineConfigUpsert struct {
		*sql.UpdateSet
	}
)

// SetPriceType sets the "price_type" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsert) SetPriceType(v plan.PriceType) *BillingInvoiceUsageBasedLineConfigUpsert {
	u.Set(billinginvoiceusagebasedlineconfig.FieldPriceType, v)
	return u
}

// UpdatePriceType sets the "price_type" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsert) UpdatePriceType() *BillingInvoiceUsageBasedLineConfigUpsert {
	u.SetExcluded(billinginvoiceusagebasedlineconfig.FieldPriceType)
	return u
}

// SetPrice sets the "price" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsert) SetPrice(v *plan.Price) *BillingInvoiceUsageBasedLineConfigUpsert {
	u.Set(billinginvoiceusagebasedlineconfig.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsert) UpdatePrice() *BillingInvoiceUsageBasedLineConfigUpsert {
	u.SetExcluded(billinginvoiceusagebasedlineconfig.FieldPrice)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billinginvoiceusagebasedlineconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) UpdateNewValues() *BillingInvoiceUsageBasedLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldID)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldNamespace)
		}
		if _, exists := u.create.mutation.FeatureKey(); exists {
			s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldFeatureKey)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) Ignore() *BillingInvoiceUsageBasedLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) DoNothing() *BillingInvoiceUsageBasedLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingInvoiceUsageBasedLineConfigCreate.OnConflict
// documentation for more info.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) Update(set func(*BillingInvoiceUsageBasedLineConfigUpsert)) *BillingInvoiceUsageBasedLineConfigUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingInvoiceUsageBasedLineConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriceType sets the "price_type" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) SetPriceType(v plan.PriceType) *BillingInvoiceUsageBasedLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.SetPriceType(v)
	})
}

// UpdatePriceType sets the "price_type" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) UpdatePriceType() *BillingInvoiceUsageBasedLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.UpdatePriceType()
	})
}

// SetPrice sets the "price" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) SetPrice(v *plan.Price) *BillingInvoiceUsageBasedLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.SetPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) UpdatePrice() *BillingInvoiceUsageBasedLineConfigUpsertOne {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingInvoiceUsageBasedLineConfigCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: BillingInvoiceUsageBasedLineConfigUpsertOne.ID is not supported by MySQL driver. Use BillingInvoiceUsageBasedLineConfigUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BillingInvoiceUsageBasedLineConfigUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BillingInvoiceUsageBasedLineConfigCreateBulk is the builder for creating many BillingInvoiceUsageBasedLineConfig entities in bulk.
type BillingInvoiceUsageBasedLineConfigCreateBulk struct {
	config
	err      error
	builders []*BillingInvoiceUsageBasedLineConfigCreate
	conflict []sql.ConflictOption
}

// Save creates the BillingInvoiceUsageBasedLineConfig entities in the database.
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) Save(ctx context.Context) ([]*BillingInvoiceUsageBasedLineConfig, error) {
	if biublccb.err != nil {
		return nil, biublccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(biublccb.builders))
	nodes := make([]*BillingInvoiceUsageBasedLineConfig, len(biublccb.builders))
	mutators := make([]Mutator, len(biublccb.builders))
	for i := range biublccb.builders {
		func(i int, root context.Context) {
			builder := biublccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingInvoiceUsageBasedLineConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i], err = builder.createSpec()
				if err != nil {
					return nil, err
				}
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, biublccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = biublccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, biublccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, biublccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) SaveX(ctx context.Context) []*BillingInvoiceUsageBasedLineConfig {
	v, err := biublccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := biublccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) ExecX(ctx context.Context) {
	if err := biublccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingInvoiceUsageBasedLineConfig.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingInvoiceUsageBasedLineConfigUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) OnConflict(opts ...sql.ConflictOption) *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	biublccb.conflict = opts
	return &BillingInvoiceUsageBasedLineConfigUpsertBulk{
		create: biublccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (biublccb *BillingInvoiceUsageBasedLineConfigCreateBulk) OnConflictColumns(columns ...string) *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	biublccb.conflict = append(biublccb.conflict, sql.ConflictColumns(columns...))
	return &BillingInvoiceUsageBasedLineConfigUpsertBulk{
		create: biublccb,
	}
}

// BillingInvoiceUsageBasedLineConfigUpsertBulk is the builder for "upsert"-ing
// a bulk of BillingInvoiceUsageBasedLineConfig nodes.
type BillingInvoiceUsageBasedLineConfigUpsertBulk struct {
	create *BillingInvoiceUsageBasedLineConfigCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billinginvoiceusagebasedlineconfig.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) UpdateNewValues() *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldID)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldNamespace)
			}
			if _, exists := b.mutation.FeatureKey(); exists {
				s.SetIgnore(billinginvoiceusagebasedlineconfig.FieldFeatureKey)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingInvoiceUsageBasedLineConfig.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) Ignore() *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) DoNothing() *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingInvoiceUsageBasedLineConfigCreateBulk.OnConflict
// documentation for more info.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) Update(set func(*BillingInvoiceUsageBasedLineConfigUpsert)) *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingInvoiceUsageBasedLineConfigUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriceType sets the "price_type" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) SetPriceType(v plan.PriceType) *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.SetPriceType(v)
	})
}

// UpdatePriceType sets the "price_type" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) UpdatePriceType() *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.UpdatePriceType()
	})
}

// SetPrice sets the "price" field.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) SetPrice(v *plan.Price) *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.SetPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) UpdatePrice() *BillingInvoiceUsageBasedLineConfigUpsertBulk {
	return u.Update(func(s *BillingInvoiceUsageBasedLineConfigUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the BillingInvoiceUsageBasedLineConfigCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingInvoiceUsageBasedLineConfigCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingInvoiceUsageBasedLineConfigUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}