// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks     []Hook
	mutation  *MemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetAddress sets the "address" field.
func (mu *MemberUpdate) SetAddress(s string) *MemberUpdate {
	mu.mutation.SetAddress(s)
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickname(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickname(*s)
	}
	return mu
}

// ClearNickname clears the value of the "nickname" field.
func (mu *MemberUpdate) ClearNickname() *MemberUpdate {
	mu.mutation.ClearNickname()
	return mu
}

// SetAvatar sets the "avatar" field.
func (mu *MemberUpdate) SetAvatar(s string) *MemberUpdate {
	mu.mutation.SetAvatar(s)
	return mu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableAvatar(s *string) *MemberUpdate {
	if s != nil {
		mu.SetAvatar(*s)
	}
	return mu
}

// ClearAvatar clears the value of the "avatar" field.
func (mu *MemberUpdate) ClearAvatar() *MemberUpdate {
	mu.mutation.ClearAvatar()
	return mu
}

// SetIntro sets the "intro" field.
func (mu *MemberUpdate) SetIntro(s string) *MemberUpdate {
	mu.mutation.SetIntro(s)
	return mu
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableIntro(s *string) *MemberUpdate {
	if s != nil {
		mu.SetIntro(*s)
	}
	return mu
}

// ClearIntro clears the value of the "intro" field.
func (mu *MemberUpdate) ClearIntro() *MemberUpdate {
	mu.mutation.ClearIntro()
	return mu
}

// SetNonce sets the "nonce" field.
func (mu *MemberUpdate) SetNonce(s string) *MemberUpdate {
	mu.mutation.SetNonce(s)
	return mu
}

// SetShowNickname sets the "show_nickname" field.
func (mu *MemberUpdate) SetShowNickname(b bool) *MemberUpdate {
	mu.mutation.SetShowNickname(b)
	return mu
}

// SetNillableShowNickname sets the "show_nickname" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableShowNickname(b *bool) *MemberUpdate {
	if b != nil {
		mu.SetShowNickname(*b)
	}
	return mu
}

// AddOwnGroupIDs adds the "own_groups" edge to the Group entity by IDs.
func (mu *MemberUpdate) AddOwnGroupIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.AddOwnGroupIDs(ids...)
	return mu
}

// AddOwnGroups adds the "own_groups" edges to the Group entity.
func (mu *MemberUpdate) AddOwnGroups(g ...*Group) *MemberUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.AddOwnGroupIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (mu *MemberUpdate) AddMessageIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.AddMessageIDs(ids...)
	return mu
}

// AddMessages adds the "messages" edges to the Message entity.
func (mu *MemberUpdate) AddMessages(m ...*Message) *MemberUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMessageIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (mu *MemberUpdate) AddGroupIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.AddGroupIDs(ids...)
	return mu
}

// AddGroups adds the "groups" edges to the Group entity.
func (mu *MemberUpdate) AddGroups(g ...*Group) *MemberUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.AddGroupIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearOwnGroups clears all "own_groups" edges to the Group entity.
func (mu *MemberUpdate) ClearOwnGroups() *MemberUpdate {
	mu.mutation.ClearOwnGroups()
	return mu
}

// RemoveOwnGroupIDs removes the "own_groups" edge to Group entities by IDs.
func (mu *MemberUpdate) RemoveOwnGroupIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.RemoveOwnGroupIDs(ids...)
	return mu
}

// RemoveOwnGroups removes "own_groups" edges to Group entities.
func (mu *MemberUpdate) RemoveOwnGroups(g ...*Group) *MemberUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.RemoveOwnGroupIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (mu *MemberUpdate) ClearMessages() *MemberUpdate {
	mu.mutation.ClearMessages()
	return mu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (mu *MemberUpdate) RemoveMessageIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.RemoveMessageIDs(ids...)
	return mu
}

// RemoveMessages removes "messages" edges to Message entities.
func (mu *MemberUpdate) RemoveMessages(m ...*Message) *MemberUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMessageIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (mu *MemberUpdate) ClearGroups() *MemberUpdate {
	mu.mutation.ClearGroups()
	return mu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (mu *MemberUpdate) RemoveGroupIDs(ids ...uint64) *MemberUpdate {
	mu.mutation.RemoveGroupIDs(ids...)
	return mu
}

// RemoveGroups removes "groups" edges to Group entities.
func (mu *MemberUpdate) RemoveGroups(g ...*Group) *MemberUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MemberUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
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
	if value, ok := mu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAddress,
		})
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickname,
		})
	}
	if mu.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldNickname,
		})
	}
	if value, ok := mu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAvatar,
		})
	}
	if mu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldAvatar,
		})
	}
	if value, ok := mu.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldIntro,
		})
	}
	if mu.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldIntro,
		})
	}
	if value, ok := mu.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNonce,
		})
	}
	if value, ok := mu.mutation.ShowNickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: member.FieldShowNickname,
		})
	}
	if mu.mutation.OwnGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedOwnGroupsIDs(); len(nodes) > 0 && !mu.mutation.OwnGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OwnGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !mu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !mu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
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
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetAddress sets the "address" field.
func (muo *MemberUpdateOne) SetAddress(s string) *MemberUpdateOne {
	muo.mutation.SetAddress(s)
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickname(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickname(*s)
	}
	return muo
}

// ClearNickname clears the value of the "nickname" field.
func (muo *MemberUpdateOne) ClearNickname() *MemberUpdateOne {
	muo.mutation.ClearNickname()
	return muo
}

// SetAvatar sets the "avatar" field.
func (muo *MemberUpdateOne) SetAvatar(s string) *MemberUpdateOne {
	muo.mutation.SetAvatar(s)
	return muo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableAvatar(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetAvatar(*s)
	}
	return muo
}

// ClearAvatar clears the value of the "avatar" field.
func (muo *MemberUpdateOne) ClearAvatar() *MemberUpdateOne {
	muo.mutation.ClearAvatar()
	return muo
}

// SetIntro sets the "intro" field.
func (muo *MemberUpdateOne) SetIntro(s string) *MemberUpdateOne {
	muo.mutation.SetIntro(s)
	return muo
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableIntro(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetIntro(*s)
	}
	return muo
}

// ClearIntro clears the value of the "intro" field.
func (muo *MemberUpdateOne) ClearIntro() *MemberUpdateOne {
	muo.mutation.ClearIntro()
	return muo
}

// SetNonce sets the "nonce" field.
func (muo *MemberUpdateOne) SetNonce(s string) *MemberUpdateOne {
	muo.mutation.SetNonce(s)
	return muo
}

// SetShowNickname sets the "show_nickname" field.
func (muo *MemberUpdateOne) SetShowNickname(b bool) *MemberUpdateOne {
	muo.mutation.SetShowNickname(b)
	return muo
}

// SetNillableShowNickname sets the "show_nickname" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableShowNickname(b *bool) *MemberUpdateOne {
	if b != nil {
		muo.SetShowNickname(*b)
	}
	return muo
}

// AddOwnGroupIDs adds the "own_groups" edge to the Group entity by IDs.
func (muo *MemberUpdateOne) AddOwnGroupIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.AddOwnGroupIDs(ids...)
	return muo
}

// AddOwnGroups adds the "own_groups" edges to the Group entity.
func (muo *MemberUpdateOne) AddOwnGroups(g ...*Group) *MemberUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.AddOwnGroupIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (muo *MemberUpdateOne) AddMessageIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.AddMessageIDs(ids...)
	return muo
}

// AddMessages adds the "messages" edges to the Message entity.
func (muo *MemberUpdateOne) AddMessages(m ...*Message) *MemberUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMessageIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (muo *MemberUpdateOne) AddGroupIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.AddGroupIDs(ids...)
	return muo
}

// AddGroups adds the "groups" edges to the Group entity.
func (muo *MemberUpdateOne) AddGroups(g ...*Group) *MemberUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.AddGroupIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearOwnGroups clears all "own_groups" edges to the Group entity.
func (muo *MemberUpdateOne) ClearOwnGroups() *MemberUpdateOne {
	muo.mutation.ClearOwnGroups()
	return muo
}

// RemoveOwnGroupIDs removes the "own_groups" edge to Group entities by IDs.
func (muo *MemberUpdateOne) RemoveOwnGroupIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.RemoveOwnGroupIDs(ids...)
	return muo
}

// RemoveOwnGroups removes "own_groups" edges to Group entities.
func (muo *MemberUpdateOne) RemoveOwnGroups(g ...*Group) *MemberUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.RemoveOwnGroupIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (muo *MemberUpdateOne) ClearMessages() *MemberUpdateOne {
	muo.mutation.ClearMessages()
	return muo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (muo *MemberUpdateOne) RemoveMessageIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.RemoveMessageIDs(ids...)
	return muo
}

// RemoveMessages removes "messages" edges to Message entities.
func (muo *MemberUpdateOne) RemoveMessages(m ...*Message) *MemberUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMessageIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (muo *MemberUpdateOne) ClearGroups() *MemberUpdateOne {
	muo.mutation.ClearGroups()
	return muo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (muo *MemberUpdateOne) RemoveGroupIDs(ids ...uint64) *MemberUpdateOne {
	muo.mutation.RemoveGroupIDs(ids...)
	return muo
}

// RemoveGroups removes "groups" edges to Group entities.
func (muo *MemberUpdateOne) RemoveGroups(g ...*Group) *MemberUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return muo.RemoveGroupIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	var (
		err  error
		node *Member
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
		nv, ok := v.(*Member)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MemberMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MemberUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MemberUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
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
	if value, ok := muo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAddress,
		})
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickname,
		})
	}
	if muo.mutation.NicknameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldNickname,
		})
	}
	if value, ok := muo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAvatar,
		})
	}
	if muo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldAvatar,
		})
	}
	if value, ok := muo.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldIntro,
		})
	}
	if muo.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldIntro,
		})
	}
	if value, ok := muo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNonce,
		})
	}
	if value, ok := muo.mutation.ShowNickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: member.FieldShowNickname,
		})
	}
	if muo.mutation.OwnGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedOwnGroupsIDs(); len(nodes) > 0 && !muo.mutation.OwnGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OwnGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !muo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !muo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = muo.modifiers
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
