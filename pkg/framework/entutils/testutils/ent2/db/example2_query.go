// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent2/db/example2"
	"github.com/openmeterio/openmeter/pkg/framework/entutils/testutils/ent2/db/predicate"
)

// Example2Query is the builder for querying Example2 entities.
type Example2Query struct {
	config
	ctx        *QueryContext
	order      []example2.OrderOption
	inters     []Interceptor
	predicates []predicate.Example2
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the Example2Query builder.
func (e *Example2Query) Where(ps ...predicate.Example2) *Example2Query {
	e.predicates = append(e.predicates, ps...)
	return e
}

// Limit the number of records to be returned by this query.
func (e *Example2Query) Limit(limit int) *Example2Query {
	e.ctx.Limit = &limit
	return e
}

// Offset to start from.
func (e *Example2Query) Offset(offset int) *Example2Query {
	e.ctx.Offset = &offset
	return e
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (e *Example2Query) Unique(unique bool) *Example2Query {
	e.ctx.Unique = &unique
	return e
}

// Order specifies how the records should be ordered.
func (e *Example2Query) Order(o ...example2.OrderOption) *Example2Query {
	e.order = append(e.order, o...)
	return e
}

// First returns the first Example2 entity from the query.
// Returns a *NotFoundError when no Example2 was found.
func (e *Example2Query) First(ctx context.Context) (*Example2, error) {
	nodes, err := e.Limit(1).All(setContextOp(ctx, e.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{example2.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (e *Example2Query) FirstX(ctx context.Context) *Example2 {
	node, err := e.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Example2 ID from the query.
// Returns a *NotFoundError when no Example2 ID was found.
func (e *Example2Query) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = e.Limit(1).IDs(setContextOp(ctx, e.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{example2.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (e *Example2Query) FirstIDX(ctx context.Context) string {
	id, err := e.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Example2 entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Example2 entity is found.
// Returns a *NotFoundError when no Example2 entities are found.
func (e *Example2Query) Only(ctx context.Context) (*Example2, error) {
	nodes, err := e.Limit(2).All(setContextOp(ctx, e.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{example2.Label}
	default:
		return nil, &NotSingularError{example2.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (e *Example2Query) OnlyX(ctx context.Context) *Example2 {
	node, err := e.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Example2 ID in the query.
// Returns a *NotSingularError when more than one Example2 ID is found.
// Returns a *NotFoundError when no entities are found.
func (e *Example2Query) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = e.Limit(2).IDs(setContextOp(ctx, e.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{example2.Label}
	default:
		err = &NotSingularError{example2.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (e *Example2Query) OnlyIDX(ctx context.Context) string {
	id, err := e.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Example2s.
func (e *Example2Query) All(ctx context.Context) ([]*Example2, error) {
	ctx = setContextOp(ctx, e.ctx, ent.OpQueryAll)
	if err := e.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Example2, *Example2Query]()
	return withInterceptors[[]*Example2](ctx, e, qr, e.inters)
}

// AllX is like All, but panics if an error occurs.
func (e *Example2Query) AllX(ctx context.Context) []*Example2 {
	nodes, err := e.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Example2 IDs.
func (e *Example2Query) IDs(ctx context.Context) (ids []string, err error) {
	if e.ctx.Unique == nil && e.path != nil {
		e.Unique(true)
	}
	ctx = setContextOp(ctx, e.ctx, ent.OpQueryIDs)
	if err = e.Select(example2.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (e *Example2Query) IDsX(ctx context.Context) []string {
	ids, err := e.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (e *Example2Query) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, e.ctx, ent.OpQueryCount)
	if err := e.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, e, querierCount[*Example2Query](), e.inters)
}

// CountX is like Count, but panics if an error occurs.
func (e *Example2Query) CountX(ctx context.Context) int {
	count, err := e.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (e *Example2Query) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, e.ctx, ent.OpQueryExist)
	switch _, err := e.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (e *Example2Query) ExistX(ctx context.Context) bool {
	exist, err := e.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the Example2Query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (e *Example2Query) Clone() *Example2Query {
	if e == nil {
		return nil
	}
	return &Example2Query{
		config:     e.config,
		ctx:        e.ctx.Clone(),
		order:      append([]example2.OrderOption{}, e.order...),
		inters:     append([]Interceptor{}, e.inters...),
		predicates: append([]predicate.Example2{}, e.predicates...),
		// clone intermediate query.
		sql:  e.sql.Clone(),
		path: e.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Example2.Query().
//		GroupBy(example2.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (e *Example2Query) GroupBy(field string, fields ...string) *Example2GroupBy {
	e.ctx.Fields = append([]string{field}, fields...)
	grbuild := &Example2GroupBy{build: e}
	grbuild.flds = &e.ctx.Fields
	grbuild.label = example2.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Example2.Query().
//		Select(example2.FieldCreatedAt).
//		Scan(ctx, &v)
func (e *Example2Query) Select(fields ...string) *Example2Select {
	e.ctx.Fields = append(e.ctx.Fields, fields...)
	sbuild := &Example2Select{Example2Query: e}
	sbuild.label = example2.Label
	sbuild.flds, sbuild.scan = &e.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a Example2Select configured with the given aggregations.
func (e *Example2Query) Aggregate(fns ...AggregateFunc) *Example2Select {
	return e.Select().Aggregate(fns...)
}

func (e *Example2Query) prepareQuery(ctx context.Context) error {
	for _, inter := range e.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, e); err != nil {
				return err
			}
		}
	}
	for _, f := range e.ctx.Fields {
		if !example2.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if e.path != nil {
		prev, err := e.path(ctx)
		if err != nil {
			return err
		}
		e.sql = prev
	}
	return nil
}

func (e *Example2Query) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Example2, error) {
	var (
		nodes = []*Example2{}
		_spec = e.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Example2).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Example2{config: e.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(e.modifiers) > 0 {
		_spec.Modifiers = e.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, e.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (e *Example2Query) sqlCount(ctx context.Context) (int, error) {
	_spec := e.querySpec()
	if len(e.modifiers) > 0 {
		_spec.Modifiers = e.modifiers
	}
	_spec.Node.Columns = e.ctx.Fields
	if len(e.ctx.Fields) > 0 {
		_spec.Unique = e.ctx.Unique != nil && *e.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, e.driver, _spec)
}

func (e *Example2Query) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(example2.Table, example2.Columns, sqlgraph.NewFieldSpec(example2.FieldID, field.TypeString))
	_spec.From = e.sql
	if unique := e.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if e.path != nil {
		_spec.Unique = true
	}
	if fields := e.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, example2.FieldID)
		for i := range fields {
			if fields[i] != example2.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := e.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := e.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := e.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := e.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (e *Example2Query) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(e.driver.Dialect())
	t1 := builder.Table(example2.Table)
	columns := e.ctx.Fields
	if len(columns) == 0 {
		columns = example2.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if e.sql != nil {
		selector = e.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if e.ctx.Unique != nil && *e.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range e.modifiers {
		m(selector)
	}
	for _, p := range e.predicates {
		p(selector)
	}
	for _, p := range e.order {
		p(selector)
	}
	if offset := e.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := e.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (e *Example2Query) ForUpdate(opts ...sql.LockOption) *Example2Query {
	if e.driver.Dialect() == dialect.Postgres {
		e.Unique(false)
	}
	e.modifiers = append(e.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return e
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (e *Example2Query) ForShare(opts ...sql.LockOption) *Example2Query {
	if e.driver.Dialect() == dialect.Postgres {
		e.Unique(false)
	}
	e.modifiers = append(e.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return e
}

// Example2GroupBy is the group-by builder for Example2 entities.
type Example2GroupBy struct {
	selector
	build *Example2Query
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eb *Example2GroupBy) Aggregate(fns ...AggregateFunc) *Example2GroupBy {
	eb.fns = append(eb.fns, fns...)
	return eb
}

// Scan applies the selector query and scans the result into the given value.
func (eb *Example2GroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eb.build.ctx, ent.OpQueryGroupBy)
	if err := eb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Example2Query, *Example2GroupBy](ctx, eb.build, eb, eb.build.inters, v)
}

func (eb *Example2GroupBy) sqlScan(ctx context.Context, root *Example2Query, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(eb.fns))
	for _, fn := range eb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*eb.flds)+len(eb.fns))
		for _, f := range *eb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*eb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Example2Select is the builder for selecting fields of Example2 entities.
type Example2Select struct {
	*Example2Query
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (e *Example2Select) Aggregate(fns ...AggregateFunc) *Example2Select {
	e.fns = append(e.fns, fns...)
	return e
}

// Scan applies the selector query and scans the result into the given value.
func (e *Example2Select) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, e.ctx, ent.OpQuerySelect)
	if err := e.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Example2Query, *Example2Select](ctx, e.Example2Query, e, e.inters, v)
}

func (e *Example2Select) sqlScan(ctx context.Context, root *Example2Query, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(e.fns))
	for _, fn := range e.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*e.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := e.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
