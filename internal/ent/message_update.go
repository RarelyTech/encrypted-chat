// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/app/model"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks     []Hook
	mutation  *MessageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetGroupID sets the "group_id" field.
func (mu *MessageUpdate) SetGroupID(s string) *MessageUpdate {
	mu.mutation.SetGroupID(s)
	return mu
}

// SetMemberID sets the "member_id" field.
func (mu *MessageUpdate) SetMemberID(s string) *MessageUpdate {
	mu.mutation.SetMemberID(s)
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(b []byte) *MessageUpdate {
	mu.mutation.SetContent(b)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MessageUpdate) SetParentID(s string) *MessageUpdate {
	mu.mutation.SetParentID(s)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableParentID(s *string) *MessageUpdate {
	if s != nil {
		mu.SetParentID(*s)
	}
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MessageUpdate) ClearParentID() *MessageUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetOwner sets the "owner" field.
func (mu *MessageUpdate) SetOwner(m *model.Member) *MessageUpdate {
	mu.mutation.SetOwner(m)
	return mu
}

// SetLastNode sets the "last_node" field.
func (mu *MessageUpdate) SetLastNode(i int64) *MessageUpdate {
	mu.mutation.ResetLastNode()
	mu.mutation.SetLastNode(i)
	return mu
}

// AddLastNode adds i to the "last_node" field.
func (mu *MessageUpdate) AddLastNode(i int64) *MessageUpdate {
	mu.mutation.AddLastNode(i)
	return mu
}

// SetMember sets the "member" edge to the Member entity.
func (mu *MessageUpdate) SetMember(m *Member) *MessageUpdate {
	return mu.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (mu *MessageUpdate) SetGroup(g *Group) *MessageUpdate {
	return mu.SetGroupID(g.ID)
}

// SetParent sets the "parent" edge to the Message entity.
func (mu *MessageUpdate) SetParent(m *Message) *MessageUpdate {
	return mu.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Message entity by IDs.
func (mu *MessageUpdate) AddChildIDs(ids ...string) *MessageUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Message entity.
func (mu *MessageUpdate) AddChildren(m ...*Message) *MessageUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (mu *MessageUpdate) ClearMember() *MessageUpdate {
	mu.mutation.ClearMember()
	return mu
}

// ClearGroup clears the "group" edge to the Group entity.
func (mu *MessageUpdate) ClearGroup() *MessageUpdate {
	mu.mutation.ClearGroup()
	return mu
}

// ClearParent clears the "parent" edge to the Message entity.
func (mu *MessageUpdate) ClearParent() *MessageUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearChildren clears all "children" edges to the Message entity.
func (mu *MessageUpdate) ClearChildren() *MessageUpdate {
	mu.mutation.ClearChildren()
	return mu
}

// RemoveChildIDs removes the "children" edge to Message entities by IDs.
func (mu *MessageUpdate) RemoveChildIDs(ids ...string) *MessageUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Message entities.
func (mu *MessageUpdate) RemoveChildren(m ...*Message) *MessageUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if _, ok := mu.mutation.MemberID(); mu.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.member"`)
	}
	if _, ok := mu.mutation.GroupID(); mu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MessageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: message.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if value, ok := mu.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldOwner,
		})
	}
	if value, ok := mu.mutation.LastNode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLastNode,
		})
	}
	if value, ok := mu.mutation.AddedLastNode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLastNode,
		})
	}
	if mu.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.MemberTable,
			Columns: []string{message.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.MemberTable,
			Columns: []string{message.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = mu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MessageMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetGroupID sets the "group_id" field.
func (muo *MessageUpdateOne) SetGroupID(s string) *MessageUpdateOne {
	muo.mutation.SetGroupID(s)
	return muo
}

// SetMemberID sets the "member_id" field.
func (muo *MessageUpdateOne) SetMemberID(s string) *MessageUpdateOne {
	muo.mutation.SetMemberID(s)
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(b []byte) *MessageUpdateOne {
	muo.mutation.SetContent(b)
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MessageUpdateOne) SetParentID(s string) *MessageUpdateOne {
	muo.mutation.SetParentID(s)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableParentID(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetParentID(*s)
	}
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MessageUpdateOne) ClearParentID() *MessageUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetOwner sets the "owner" field.
func (muo *MessageUpdateOne) SetOwner(m *model.Member) *MessageUpdateOne {
	muo.mutation.SetOwner(m)
	return muo
}

// SetLastNode sets the "last_node" field.
func (muo *MessageUpdateOne) SetLastNode(i int64) *MessageUpdateOne {
	muo.mutation.ResetLastNode()
	muo.mutation.SetLastNode(i)
	return muo
}

// AddLastNode adds i to the "last_node" field.
func (muo *MessageUpdateOne) AddLastNode(i int64) *MessageUpdateOne {
	muo.mutation.AddLastNode(i)
	return muo
}

// SetMember sets the "member" edge to the Member entity.
func (muo *MessageUpdateOne) SetMember(m *Member) *MessageUpdateOne {
	return muo.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (muo *MessageUpdateOne) SetGroup(g *Group) *MessageUpdateOne {
	return muo.SetGroupID(g.ID)
}

// SetParent sets the "parent" edge to the Message entity.
func (muo *MessageUpdateOne) SetParent(m *Message) *MessageUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Message entity by IDs.
func (muo *MessageUpdateOne) AddChildIDs(ids ...string) *MessageUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Message entity.
func (muo *MessageUpdateOne) AddChildren(m ...*Message) *MessageUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (muo *MessageUpdateOne) ClearMember() *MessageUpdateOne {
	muo.mutation.ClearMember()
	return muo
}

// ClearGroup clears the "group" edge to the Group entity.
func (muo *MessageUpdateOne) ClearGroup() *MessageUpdateOne {
	muo.mutation.ClearGroup()
	return muo
}

// ClearParent clears the "parent" edge to the Message entity.
func (muo *MessageUpdateOne) ClearParent() *MessageUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearChildren clears all "children" edges to the Message entity.
func (muo *MessageUpdateOne) ClearChildren() *MessageUpdateOne {
	muo.mutation.ClearChildren()
	return muo
}

// RemoveChildIDs removes the "children" edge to Message entities by IDs.
func (muo *MessageUpdateOne) RemoveChildIDs(ids ...string) *MessageUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Message entities.
func (muo *MessageUpdateOne) RemoveChildren(m ...*Message) *MessageUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if _, ok := muo.mutation.MemberID(); muo.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.member"`)
	}
	if _, ok := muo.mutation.GroupID(); muo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Message.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MessageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MessageUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: message.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: message.FieldContent,
		})
	}
	if value, ok := muo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: message.FieldOwner,
		})
	}
	if value, ok := muo.mutation.LastNode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLastNode,
		})
	}
	if value, ok := muo.mutation.AddedLastNode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldLastNode,
		})
	}
	if muo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.MemberTable,
			Columns: []string{message.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.MemberTable,
			Columns: []string{message.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.ParentTable,
			Columns: []string{message.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   message.ChildrenTable,
			Columns: []string{message.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = muo.modifiers
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
