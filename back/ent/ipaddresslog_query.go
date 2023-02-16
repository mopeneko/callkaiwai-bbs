// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/ipaddresslog"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/post"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/predicate"
)

// IPAddressLogQuery is the builder for querying IPAddressLog entities.
type IPAddressLogQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.IPAddressLog
	withPostID *PostQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IPAddressLogQuery builder.
func (ialq *IPAddressLogQuery) Where(ps ...predicate.IPAddressLog) *IPAddressLogQuery {
	ialq.predicates = append(ialq.predicates, ps...)
	return ialq
}

// Limit the number of records to be returned by this query.
func (ialq *IPAddressLogQuery) Limit(limit int) *IPAddressLogQuery {
	ialq.ctx.Limit = &limit
	return ialq
}

// Offset to start from.
func (ialq *IPAddressLogQuery) Offset(offset int) *IPAddressLogQuery {
	ialq.ctx.Offset = &offset
	return ialq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ialq *IPAddressLogQuery) Unique(unique bool) *IPAddressLogQuery {
	ialq.ctx.Unique = &unique
	return ialq
}

// Order specifies how the records should be ordered.
func (ialq *IPAddressLogQuery) Order(o ...OrderFunc) *IPAddressLogQuery {
	ialq.order = append(ialq.order, o...)
	return ialq
}

// QueryPostID chains the current query on the "post_id" edge.
func (ialq *IPAddressLogQuery) QueryPostID() *PostQuery {
	query := (&PostClient{config: ialq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ialq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ialq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipaddresslog.Table, ipaddresslog.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ipaddresslog.PostIDTable, ipaddresslog.PostIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(ialq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IPAddressLog entity from the query.
// Returns a *NotFoundError when no IPAddressLog was found.
func (ialq *IPAddressLogQuery) First(ctx context.Context) (*IPAddressLog, error) {
	nodes, err := ialq.Limit(1).All(setContextOp(ctx, ialq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ipaddresslog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ialq *IPAddressLogQuery) FirstX(ctx context.Context) *IPAddressLog {
	node, err := ialq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IPAddressLog ID from the query.
// Returns a *NotFoundError when no IPAddressLog ID was found.
func (ialq *IPAddressLogQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ialq.Limit(1).IDs(setContextOp(ctx, ialq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ipaddresslog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ialq *IPAddressLogQuery) FirstIDX(ctx context.Context) string {
	id, err := ialq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IPAddressLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IPAddressLog entity is found.
// Returns a *NotFoundError when no IPAddressLog entities are found.
func (ialq *IPAddressLogQuery) Only(ctx context.Context) (*IPAddressLog, error) {
	nodes, err := ialq.Limit(2).All(setContextOp(ctx, ialq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ipaddresslog.Label}
	default:
		return nil, &NotSingularError{ipaddresslog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ialq *IPAddressLogQuery) OnlyX(ctx context.Context) *IPAddressLog {
	node, err := ialq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IPAddressLog ID in the query.
// Returns a *NotSingularError when more than one IPAddressLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (ialq *IPAddressLogQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ialq.Limit(2).IDs(setContextOp(ctx, ialq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ipaddresslog.Label}
	default:
		err = &NotSingularError{ipaddresslog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ialq *IPAddressLogQuery) OnlyIDX(ctx context.Context) string {
	id, err := ialq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IPAddressLogs.
func (ialq *IPAddressLogQuery) All(ctx context.Context) ([]*IPAddressLog, error) {
	ctx = setContextOp(ctx, ialq.ctx, "All")
	if err := ialq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IPAddressLog, *IPAddressLogQuery]()
	return withInterceptors[[]*IPAddressLog](ctx, ialq, qr, ialq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ialq *IPAddressLogQuery) AllX(ctx context.Context) []*IPAddressLog {
	nodes, err := ialq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IPAddressLog IDs.
func (ialq *IPAddressLogQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ialq.ctx.Unique == nil && ialq.path != nil {
		ialq.Unique(true)
	}
	ctx = setContextOp(ctx, ialq.ctx, "IDs")
	if err = ialq.Select(ipaddresslog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ialq *IPAddressLogQuery) IDsX(ctx context.Context) []string {
	ids, err := ialq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ialq *IPAddressLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ialq.ctx, "Count")
	if err := ialq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ialq, querierCount[*IPAddressLogQuery](), ialq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ialq *IPAddressLogQuery) CountX(ctx context.Context) int {
	count, err := ialq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ialq *IPAddressLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ialq.ctx, "Exist")
	switch _, err := ialq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ialq *IPAddressLogQuery) ExistX(ctx context.Context) bool {
	exist, err := ialq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IPAddressLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ialq *IPAddressLogQuery) Clone() *IPAddressLogQuery {
	if ialq == nil {
		return nil
	}
	return &IPAddressLogQuery{
		config:     ialq.config,
		ctx:        ialq.ctx.Clone(),
		order:      append([]OrderFunc{}, ialq.order...),
		inters:     append([]Interceptor{}, ialq.inters...),
		predicates: append([]predicate.IPAddressLog{}, ialq.predicates...),
		withPostID: ialq.withPostID.Clone(),
		// clone intermediate query.
		sql:  ialq.sql.Clone(),
		path: ialq.path,
	}
}

// WithPostID tells the query-builder to eager-load the nodes that are connected to
// the "post_id" edge. The optional arguments are used to configure the query builder of the edge.
func (ialq *IPAddressLogQuery) WithPostID(opts ...func(*PostQuery)) *IPAddressLogQuery {
	query := (&PostClient{config: ialq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ialq.withPostID = query
	return ialq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IPAddress string `json:"ip_address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IPAddressLog.Query().
//		GroupBy(ipaddresslog.FieldIPAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ialq *IPAddressLogQuery) GroupBy(field string, fields ...string) *IPAddressLogGroupBy {
	ialq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IPAddressLogGroupBy{build: ialq}
	grbuild.flds = &ialq.ctx.Fields
	grbuild.label = ipaddresslog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IPAddress string `json:"ip_address,omitempty"`
//	}
//
//	client.IPAddressLog.Query().
//		Select(ipaddresslog.FieldIPAddress).
//		Scan(ctx, &v)
//
func (ialq *IPAddressLogQuery) Select(fields ...string) *IPAddressLogSelect {
	ialq.ctx.Fields = append(ialq.ctx.Fields, fields...)
	sbuild := &IPAddressLogSelect{IPAddressLogQuery: ialq}
	sbuild.label = ipaddresslog.Label
	sbuild.flds, sbuild.scan = &ialq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IPAddressLogSelect configured with the given aggregations.
func (ialq *IPAddressLogQuery) Aggregate(fns ...AggregateFunc) *IPAddressLogSelect {
	return ialq.Select().Aggregate(fns...)
}

func (ialq *IPAddressLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ialq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ialq); err != nil {
				return err
			}
		}
	}
	for _, f := range ialq.ctx.Fields {
		if !ipaddresslog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ialq.path != nil {
		prev, err := ialq.path(ctx)
		if err != nil {
			return err
		}
		ialq.sql = prev
	}
	return nil
}

func (ialq *IPAddressLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IPAddressLog, error) {
	var (
		nodes       = []*IPAddressLog{}
		withFKs     = ialq.withFKs
		_spec       = ialq.querySpec()
		loadedTypes = [1]bool{
			ialq.withPostID != nil,
		}
	)
	if ialq.withPostID != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddresslog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IPAddressLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IPAddressLog{config: ialq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ialq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ialq.withPostID; query != nil {
		if err := ialq.loadPostID(ctx, query, nodes, nil,
			func(n *IPAddressLog, e *Post) { n.Edges.PostID = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ialq *IPAddressLogQuery) loadPostID(ctx context.Context, query *PostQuery, nodes []*IPAddressLog, init func(*IPAddressLog), assign func(*IPAddressLog, *Post)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*IPAddressLog)
	for i := range nodes {
		if nodes[i].post_ip_address_log == nil {
			continue
		}
		fk := *nodes[i].post_ip_address_log
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_ip_address_log" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ialq *IPAddressLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ialq.querySpec()
	_spec.Node.Columns = ialq.ctx.Fields
	if len(ialq.ctx.Fields) > 0 {
		_spec.Unique = ialq.ctx.Unique != nil && *ialq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ialq.driver, _spec)
}

func (ialq *IPAddressLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ipaddresslog.Table, ipaddresslog.Columns, sqlgraph.NewFieldSpec(ipaddresslog.FieldID, field.TypeString))
	_spec.From = ialq.sql
	if unique := ialq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ialq.path != nil {
		_spec.Unique = true
	}
	if fields := ialq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddresslog.FieldID)
		for i := range fields {
			if fields[i] != ipaddresslog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ialq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ialq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ialq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ialq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ialq *IPAddressLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ialq.driver.Dialect())
	t1 := builder.Table(ipaddresslog.Table)
	columns := ialq.ctx.Fields
	if len(columns) == 0 {
		columns = ipaddresslog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ialq.sql != nil {
		selector = ialq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ialq.ctx.Unique != nil && *ialq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ialq.predicates {
		p(selector)
	}
	for _, p := range ialq.order {
		p(selector)
	}
	if offset := ialq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ialq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IPAddressLogGroupBy is the group-by builder for IPAddressLog entities.
type IPAddressLogGroupBy struct {
	selector
	build *IPAddressLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ialgb *IPAddressLogGroupBy) Aggregate(fns ...AggregateFunc) *IPAddressLogGroupBy {
	ialgb.fns = append(ialgb.fns, fns...)
	return ialgb
}

// Scan applies the selector query and scans the result into the given value.
func (ialgb *IPAddressLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ialgb.build.ctx, "GroupBy")
	if err := ialgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPAddressLogQuery, *IPAddressLogGroupBy](ctx, ialgb.build, ialgb, ialgb.build.inters, v)
}

func (ialgb *IPAddressLogGroupBy) sqlScan(ctx context.Context, root *IPAddressLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ialgb.fns))
	for _, fn := range ialgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ialgb.flds)+len(ialgb.fns))
		for _, f := range *ialgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ialgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ialgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IPAddressLogSelect is the builder for selecting fields of IPAddressLog entities.
type IPAddressLogSelect struct {
	*IPAddressLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ials *IPAddressLogSelect) Aggregate(fns ...AggregateFunc) *IPAddressLogSelect {
	ials.fns = append(ials.fns, fns...)
	return ials
}

// Scan applies the selector query and scans the result into the given value.
func (ials *IPAddressLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ials.ctx, "Select")
	if err := ials.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPAddressLogQuery, *IPAddressLogSelect](ctx, ials.IPAddressLogQuery, ials, ials.inters, v)
}

func (ials *IPAddressLogSelect) sqlScan(ctx context.Context, root *IPAddressLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ials.fns))
	for _, fn := range ials.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ials.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ials.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
