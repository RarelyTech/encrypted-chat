// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// MemberQuery is the builder for querying Member entities.
type MemberQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.Member
	withOwnGroups    *GroupQuery
	withMessages     *MessageQuery
	withGroups       *GroupQuery
	withGroupMembers *GroupMemberQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberQuery builder.
func (mq *MemberQuery) Where(ps ...predicate.Member) *MemberQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MemberQuery) Limit(limit int) *MemberQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MemberQuery) Offset(offset int) *MemberQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MemberQuery) Unique(unique bool) *MemberQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *MemberQuery) Order(o ...OrderFunc) *MemberQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryOwnGroups chains the current query on the "own_groups" edge.
func (mq *MemberQuery) QueryOwnGroups() *GroupQuery {
	query := &GroupQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.OwnGroupsTable, member.OwnGroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMessages chains the current query on the "messages" edge.
func (mq *MemberQuery) QueryMessages() *MessageQuery {
	query := &MessageQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MessagesTable, member.MessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroups chains the current query on the "groups" edge.
func (mq *MemberQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, member.GroupsTable, member.GroupsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroupMembers chains the current query on the "group_members" edge.
func (mq *MemberQuery) QueryGroupMembers() *GroupMemberQuery {
	query := &GroupMemberQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, selector),
			sqlgraph.To(groupmember.Table, groupmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, member.GroupMembersTable, member.GroupMembersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Member entity from the query.
// Returns a *NotFoundError when no Member was found.
func (mq *MemberQuery) First(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{member.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MemberQuery) FirstX(ctx context.Context) *Member {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Member ID from the query.
// Returns a *NotFoundError when no Member ID was found.
func (mq *MemberQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{member.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MemberQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Member entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Member entity is found.
// Returns a *NotFoundError when no Member entities are found.
func (mq *MemberQuery) Only(ctx context.Context) (*Member, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{member.Label}
	default:
		return nil, &NotSingularError{member.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MemberQuery) OnlyX(ctx context.Context) *Member {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Member ID in the query.
// Returns a *NotSingularError when more than one Member ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MemberQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = &NotSingularError{member.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MemberQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Members.
func (mq *MemberQuery) All(ctx context.Context) ([]*Member, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MemberQuery) AllX(ctx context.Context) []*Member {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Member IDs.
func (mq *MemberQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := mq.Select(member.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MemberQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MemberQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MemberQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MemberQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MemberQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MemberQuery) Clone() *MemberQuery {
	if mq == nil {
		return nil
	}
	return &MemberQuery{
		config:           mq.config,
		limit:            mq.limit,
		offset:           mq.offset,
		order:            append([]OrderFunc{}, mq.order...),
		predicates:       append([]predicate.Member{}, mq.predicates...),
		withOwnGroups:    mq.withOwnGroups.Clone(),
		withMessages:     mq.withMessages.Clone(),
		withGroups:       mq.withGroups.Clone(),
		withGroupMembers: mq.withGroupMembers.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithOwnGroups tells the query-builder to eager-load the nodes that are connected to
// the "own_groups" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithOwnGroups(opts ...func(*GroupQuery)) *MemberQuery {
	query := &GroupQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withOwnGroups = query
	return mq
}

// WithMessages tells the query-builder to eager-load the nodes that are connected to
// the "messages" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithMessages(opts ...func(*MessageQuery)) *MemberQuery {
	query := &MessageQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMessages = query
	return mq
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithGroups(opts ...func(*GroupQuery)) *MemberQuery {
	query := &GroupQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withGroups = query
	return mq
}

// WithGroupMembers tells the query-builder to eager-load the nodes that are connected to
// the "group_members" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MemberQuery) WithGroupMembers(opts ...func(*GroupMemberQuery)) *MemberQuery {
	query := &GroupMemberQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withGroupMembers = query
	return mq
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
//	client.Member.Query().
//		GroupBy(member.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MemberQuery) GroupBy(field string, fields ...string) *MemberGroupBy {
	grbuild := &MemberGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = member.Label
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
//	client.Member.Query().
//		Select(member.FieldCreatedAt).
//		Scan(ctx, &v)
func (mq *MemberQuery) Select(fields ...string) *MemberSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &MemberSelect{MemberQuery: mq}
	selbuild.label = member.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

func (mq *MemberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !member.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MemberQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Member, error) {
	var (
		nodes       = []*Member{}
		_spec       = mq.querySpec()
		loadedTypes = [4]bool{
			mq.withOwnGroups != nil,
			mq.withMessages != nil,
			mq.withGroups != nil,
			mq.withGroupMembers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Member).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Member{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withOwnGroups; query != nil {
		if err := mq.loadOwnGroups(ctx, query, nodes,
			func(n *Member) { n.Edges.OwnGroups = []*Group{} },
			func(n *Member, e *Group) { n.Edges.OwnGroups = append(n.Edges.OwnGroups, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withMessages; query != nil {
		if err := mq.loadMessages(ctx, query, nodes,
			func(n *Member) { n.Edges.Messages = []*Message{} },
			func(n *Member, e *Message) { n.Edges.Messages = append(n.Edges.Messages, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withGroups; query != nil {
		if err := mq.loadGroups(ctx, query, nodes,
			func(n *Member) { n.Edges.Groups = []*Group{} },
			func(n *Member, e *Group) { n.Edges.Groups = append(n.Edges.Groups, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withGroupMembers; query != nil {
		if err := mq.loadGroupMembers(ctx, query, nodes,
			func(n *Member) { n.Edges.GroupMembers = []*GroupMember{} },
			func(n *Member, e *GroupMember) { n.Edges.GroupMembers = append(n.Edges.GroupMembers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MemberQuery) loadOwnGroups(ctx context.Context, query *GroupQuery, nodes []*Member, init func(*Member), assign func(*Member, *Group)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Group(func(s *sql.Selector) {
		s.Where(sql.InValues(member.OwnGroupsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadMessages(ctx context.Context, query *MessageQuery, nodes []*Member, init func(*Member), assign func(*Member, *Message)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Message(func(s *sql.Selector) {
		s.Where(sql.InValues(member.MessagesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MemberQuery) loadGroups(ctx context.Context, query *GroupQuery, nodes []*Member, init func(*Member), assign func(*Member, *Group)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Member)
	nids := make(map[uint64]map[*Member]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(member.GroupsTable)
		s.Join(joinT).On(s.C(group.FieldID), joinT.C(member.GroupsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(member.GroupsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(member.GroupsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := uint64(values[0].(*sql.NullInt64).Int64)
			inValue := uint64(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Member]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "groups" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (mq *MemberQuery) loadGroupMembers(ctx context.Context, query *GroupMemberQuery, nodes []*Member, init func(*Member), assign func(*Member, *GroupMember)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Member)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.InValues(member.GroupMembersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MemberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mq *MemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for i := range fields {
			if fields[i] != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(member.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = member.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, m := range mq.modifiers {
		m(selector)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mq *MemberQuery) Modify(modifiers ...func(s *sql.Selector)) *MemberSelect {
	mq.modifiers = append(mq.modifiers, modifiers...)
	return mq.Select()
}

// MemberGroupBy is the group-by builder for Member entities.
type MemberGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MemberGroupBy) Aggregate(fns ...AggregateFunc) *MemberGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *MemberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *MemberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgb.fields {
		if !member.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MemberGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// MemberSelect is the builder for selecting fields of Member entities.
type MemberSelect struct {
	*MemberQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MemberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.MemberQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *MemberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ms *MemberSelect) Modify(modifiers ...func(s *sql.Selector)) *MemberSelect {
	ms.modifiers = append(ms.modifiers, modifiers...)
	return ms
}
