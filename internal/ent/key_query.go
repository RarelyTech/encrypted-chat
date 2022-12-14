// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// KeyQuery is the builder for querying Key entities.
type KeyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Key
	withMember *MemberQuery
	withGroup  *GroupQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeyQuery builder.
func (kq *KeyQuery) Where(ps ...predicate.Key) *KeyQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit adds a limit step to the query.
func (kq *KeyQuery) Limit(limit int) *KeyQuery {
	kq.limit = &limit
	return kq
}

// Offset adds an offset step to the query.
func (kq *KeyQuery) Offset(offset int) *KeyQuery {
	kq.offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KeyQuery) Unique(unique bool) *KeyQuery {
	kq.unique = &unique
	return kq
}

// Order adds an order step to the query.
func (kq *KeyQuery) Order(o ...OrderFunc) *KeyQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// QueryMember chains the current query on the "member" edge.
func (kq *KeyQuery) QueryMember() *MemberQuery {
	query := &MemberQuery{config: kq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(key.Table, key.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, key.MemberTable, key.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(kq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (kq *KeyQuery) QueryGroup() *GroupQuery {
	query := &GroupQuery{config: kq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(key.Table, key.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, key.GroupTable, key.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(kq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Key entity from the query.
// Returns a *NotFoundError when no Key was found.
func (kq *KeyQuery) First(ctx context.Context) (*Key, error) {
	nodes, err := kq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{key.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KeyQuery) FirstX(ctx context.Context) *Key {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Key ID from the query.
// Returns a *NotFoundError when no Key ID was found.
func (kq *KeyQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{key.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KeyQuery) FirstIDX(ctx context.Context) string {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Key entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Key entity is found.
// Returns a *NotFoundError when no Key entities are found.
func (kq *KeyQuery) Only(ctx context.Context) (*Key, error) {
	nodes, err := kq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{key.Label}
	default:
		return nil, &NotSingularError{key.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KeyQuery) OnlyX(ctx context.Context) *Key {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Key ID in the query.
// Returns a *NotSingularError when more than one Key ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KeyQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{key.Label}
	default:
		err = &NotSingularError{key.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KeyQuery) OnlyIDX(ctx context.Context) string {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Keys.
func (kq *KeyQuery) All(ctx context.Context) ([]*Key, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kq *KeyQuery) AllX(ctx context.Context) []*Key {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Key IDs.
func (kq *KeyQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := kq.Select(key.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KeyQuery) IDsX(ctx context.Context) []string {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KeyQuery) Count(ctx context.Context) (int, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KeyQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KeyQuery) Exist(ctx context.Context) (bool, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KeyQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KeyQuery) Clone() *KeyQuery {
	if kq == nil {
		return nil
	}
	return &KeyQuery{
		config:     kq.config,
		limit:      kq.limit,
		offset:     kq.offset,
		order:      append([]OrderFunc{}, kq.order...),
		predicates: append([]predicate.Key{}, kq.predicates...),
		withMember: kq.withMember.Clone(),
		withGroup:  kq.withGroup.Clone(),
		// clone intermediate query.
		sql:    kq.sql.Clone(),
		path:   kq.path,
		unique: kq.unique,
	}
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (kq *KeyQuery) WithMember(opts ...func(*MemberQuery)) *KeyQuery {
	query := &MemberQuery{config: kq.config}
	for _, opt := range opts {
		opt(query)
	}
	kq.withMember = query
	return kq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (kq *KeyQuery) WithGroup(opts ...func(*GroupQuery)) *KeyQuery {
	query := &GroupQuery{config: kq.config}
	for _, opt := range opts {
		opt(query)
	}
	kq.withGroup = query
	return kq
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
//	client.Key.Query().
//		GroupBy(key.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kq *KeyQuery) GroupBy(field string, fields ...string) *KeyGroupBy {
	grbuild := &KeyGroupBy{config: kq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kq.sqlQuery(ctx), nil
	}
	grbuild.label = key.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Key.Query().
//		Select(key.FieldCreatedAt).
//		Scan(ctx, &v)
func (kq *KeyQuery) Select(fields ...string) *KeySelect {
	kq.fields = append(kq.fields, fields...)
	selbuild := &KeySelect{KeyQuery: kq}
	selbuild.label = key.Label
	selbuild.flds, selbuild.scan = &kq.fields, selbuild.Scan
	return selbuild
}

func (kq *KeyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kq.fields {
		if !key.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	return nil
}

func (kq *KeyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Key, error) {
	var (
		nodes       = []*Key{}
		_spec       = kq.querySpec()
		loadedTypes = [2]bool{
			kq.withMember != nil,
			kq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Key).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Key{config: kq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := kq.withMember; query != nil {
		if err := kq.loadMember(ctx, query, nodes, nil,
			func(n *Key, e *Member) { n.Edges.Member = e }); err != nil {
			return nil, err
		}
	}
	if query := kq.withGroup; query != nil {
		if err := kq.loadGroup(ctx, query, nodes, nil,
			func(n *Key, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (kq *KeyQuery) loadMember(ctx context.Context, query *MemberQuery, nodes []*Key, init func(*Key), assign func(*Key, *Member)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Key)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (kq *KeyQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Key, init func(*Key), assign func(*Key, *Group)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Key)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (kq *KeyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	_spec.Node.Columns = kq.fields
	if len(kq.fields) > 0 {
		_spec.Unique = kq.unique != nil && *kq.unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KeyQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := kq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (kq *KeyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   key.Table,
			Columns: key.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: key.FieldID,
			},
		},
		From:   kq.sql,
		Unique: true,
	}
	if unique := kq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, key.FieldID)
		for i := range fields {
			if fields[i] != key.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KeyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(key.Table)
	columns := kq.fields
	if len(columns) == 0 {
		columns = key.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.unique != nil && *kq.unique {
		selector.Distinct()
	}
	for _, m := range kq.modifiers {
		m(selector)
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (kq *KeyQuery) Modify(modifiers ...func(s *sql.Selector)) *KeySelect {
	kq.modifiers = append(kq.modifiers, modifiers...)
	return kq.Select()
}

// KeyGroupBy is the group-by builder for Key entities.
type KeyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KeyGroupBy) Aggregate(fns ...AggregateFunc) *KeyGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kgb *KeyGroupBy) Scan(ctx context.Context, v any) error {
	query, err := kgb.path(ctx)
	if err != nil {
		return err
	}
	kgb.sql = query
	return kgb.sqlScan(ctx, v)
}

func (kgb *KeyGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range kgb.fields {
		if !key.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kgb *KeyGroupBy) sqlQuery() *sql.Selector {
	selector := kgb.sql.Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kgb.fields)+len(kgb.fns))
		for _, f := range kgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kgb.fields...)...)
}

// KeySelect is the builder for selecting fields of Key entities.
type KeySelect struct {
	*KeyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KeySelect) Scan(ctx context.Context, v any) error {
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	ks.sql = ks.KeyQuery.sqlQuery(ctx)
	return ks.sqlScan(ctx, v)
}

func (ks *KeySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ks.sql.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ks *KeySelect) Modify(modifiers ...func(s *sql.Selector)) *KeySelect {
	ks.modifiers = append(ks.modifiers, modifiers...)
	return ks
}
