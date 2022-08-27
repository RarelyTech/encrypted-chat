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
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks     []Hook
	mutation  *GroupMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetCategory sets the "category" field.
func (gu *GroupUpdate) SetCategory(s string) *GroupUpdate {
	gu.mutation.SetCategory(s)
	return gu
}

// SetOwnerID sets the "owner_id" field.
func (gu *GroupUpdate) SetOwnerID(u uint64) *GroupUpdate {
	gu.mutation.SetOwnerID(u)
	return gu
}

// SetMembersMax sets the "members_max" field.
func (gu *GroupUpdate) SetMembersMax(i int) *GroupUpdate {
	gu.mutation.ResetMembersMax()
	gu.mutation.SetMembersMax(i)
	return gu
}

// AddMembersMax adds i to the "members_max" field.
func (gu *GroupUpdate) AddMembersMax(i int) *GroupUpdate {
	gu.mutation.AddMembersMax(i)
	return gu
}

// SetMembersCount sets the "members_count" field.
func (gu *GroupUpdate) SetMembersCount(i int) *GroupUpdate {
	gu.mutation.ResetMembersCount()
	gu.mutation.SetMembersCount(i)
	return gu
}

// SetNillableMembersCount sets the "members_count" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableMembersCount(i *int) *GroupUpdate {
	if i != nil {
		gu.SetMembersCount(*i)
	}
	return gu
}

// AddMembersCount adds i to the "members_count" field.
func (gu *GroupUpdate) AddMembersCount(i int) *GroupUpdate {
	gu.mutation.AddMembersCount(i)
	return gu
}

// SetPublic sets the "public" field.
func (gu *GroupUpdate) SetPublic(b bool) *GroupUpdate {
	gu.mutation.SetPublic(b)
	return gu
}

// SetIntro sets the "intro" field.
func (gu *GroupUpdate) SetIntro(s string) *GroupUpdate {
	gu.mutation.SetIntro(s)
	return gu
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableIntro(s *string) *GroupUpdate {
	if s != nil {
		gu.SetIntro(*s)
	}
	return gu
}

// ClearIntro clears the value of the "intro" field.
func (gu *GroupUpdate) ClearIntro() *GroupUpdate {
	gu.mutation.ClearIntro()
	return gu
}

// SetKeys sets the "keys" field.
func (gu *GroupUpdate) SetKeys(s string) *GroupUpdate {
	gu.mutation.SetKeys(s)
	return gu
}

// SetOwner sets the "owner" edge to the Member entity.
func (gu *GroupUpdate) SetOwner(m *Member) *GroupUpdate {
	return gu.SetOwnerID(m.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (gu *GroupUpdate) AddMessageIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.AddMessageIDs(ids...)
	return gu
}

// AddMessages adds the "messages" edges to the Message entity.
func (gu *GroupUpdate) AddMessages(m ...*Message) *GroupUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the Member entity by IDs.
func (gu *GroupUpdate) AddMemberIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.AddMemberIDs(ids...)
	return gu
}

// AddMembers adds the "members" edges to the Member entity.
func (gu *GroupUpdate) AddMembers(m ...*Member) *GroupUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMemberIDs(ids...)
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by IDs.
func (gu *GroupUpdate) AddGroupMemberIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.AddGroupMemberIDs(ids...)
	return gu
}

// AddGroupMembers adds the "group_members" edges to the GroupMember entity.
func (gu *GroupUpdate) AddGroupMembers(g ...*GroupMember) *GroupUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddGroupMemberIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearOwner clears the "owner" edge to the Member entity.
func (gu *GroupUpdate) ClearOwner() *GroupUpdate {
	gu.mutation.ClearOwner()
	return gu
}

// ClearMessages clears all "messages" edges to the Message entity.
func (gu *GroupUpdate) ClearMessages() *GroupUpdate {
	gu.mutation.ClearMessages()
	return gu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (gu *GroupUpdate) RemoveMessageIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.RemoveMessageIDs(ids...)
	return gu
}

// RemoveMessages removes "messages" edges to Message entities.
func (gu *GroupUpdate) RemoveMessages(m ...*Message) *GroupUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMessageIDs(ids...)
}

// ClearMembers clears all "members" edges to the Member entity.
func (gu *GroupUpdate) ClearMembers() *GroupUpdate {
	gu.mutation.ClearMembers()
	return gu
}

// RemoveMemberIDs removes the "members" edge to Member entities by IDs.
func (gu *GroupUpdate) RemoveMemberIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.RemoveMemberIDs(ids...)
	return gu
}

// RemoveMembers removes "members" edges to Member entities.
func (gu *GroupUpdate) RemoveMembers(m ...*Member) *GroupUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMemberIDs(ids...)
}

// ClearGroupMembers clears all "group_members" edges to the GroupMember entity.
func (gu *GroupUpdate) ClearGroupMembers() *GroupUpdate {
	gu.mutation.ClearGroupMembers()
	return gu
}

// RemoveGroupMemberIDs removes the "group_members" edge to GroupMember entities by IDs.
func (gu *GroupUpdate) RemoveGroupMemberIDs(ids ...uint64) *GroupUpdate {
	gu.mutation.RemoveGroupMemberIDs(ids...)
	return gu
}

// RemoveGroupMembers removes "group_members" edges to GroupMember entities.
func (gu *GroupUpdate) RemoveGroupMembers(g ...*GroupMember) *GroupUpdate {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveGroupMemberIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if _, ok := gu.mutation.OwnerID(); gu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gu *GroupUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupUpdate {
	gu.modifiers = append(gu.modifiers, modifiers...)
	return gu
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: group.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := gu.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldCategory,
		})
	}
	if value, ok := gu.mutation.MembersMax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
	}
	if value, ok := gu.mutation.AddedMembersMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
	}
	if value, ok := gu.mutation.MembersCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
	}
	if value, ok := gu.mutation.AddedMembersCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
	}
	if value, ok := gu.mutation.Public(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldPublic,
		})
	}
	if value, ok := gu.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldIntro,
		})
	}
	if gu.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: group.FieldIntro,
		})
	}
	if value, ok := gu.mutation.Keys(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldKeys,
		})
	}
	if gu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if nodes := gu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !gu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if nodes := gu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if gu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		createE := &GroupMemberCreate{config: gu.config, mutation: newGroupMemberMutation(gu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !gu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMemberCreate{config: gu.config, mutation: newGroupMemberMutation(gu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMemberCreate{config: gu.config, mutation: newGroupMemberMutation(gu.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.GroupMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedGroupMembersIDs(); len(nodes) > 0 && !gu.mutation.GroupMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.GroupMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = gu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GroupMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetCategory sets the "category" field.
func (guo *GroupUpdateOne) SetCategory(s string) *GroupUpdateOne {
	guo.mutation.SetCategory(s)
	return guo
}

// SetOwnerID sets the "owner_id" field.
func (guo *GroupUpdateOne) SetOwnerID(u uint64) *GroupUpdateOne {
	guo.mutation.SetOwnerID(u)
	return guo
}

// SetMembersMax sets the "members_max" field.
func (guo *GroupUpdateOne) SetMembersMax(i int) *GroupUpdateOne {
	guo.mutation.ResetMembersMax()
	guo.mutation.SetMembersMax(i)
	return guo
}

// AddMembersMax adds i to the "members_max" field.
func (guo *GroupUpdateOne) AddMembersMax(i int) *GroupUpdateOne {
	guo.mutation.AddMembersMax(i)
	return guo
}

// SetMembersCount sets the "members_count" field.
func (guo *GroupUpdateOne) SetMembersCount(i int) *GroupUpdateOne {
	guo.mutation.ResetMembersCount()
	guo.mutation.SetMembersCount(i)
	return guo
}

// SetNillableMembersCount sets the "members_count" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableMembersCount(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetMembersCount(*i)
	}
	return guo
}

// AddMembersCount adds i to the "members_count" field.
func (guo *GroupUpdateOne) AddMembersCount(i int) *GroupUpdateOne {
	guo.mutation.AddMembersCount(i)
	return guo
}

// SetPublic sets the "public" field.
func (guo *GroupUpdateOne) SetPublic(b bool) *GroupUpdateOne {
	guo.mutation.SetPublic(b)
	return guo
}

// SetIntro sets the "intro" field.
func (guo *GroupUpdateOne) SetIntro(s string) *GroupUpdateOne {
	guo.mutation.SetIntro(s)
	return guo
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableIntro(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetIntro(*s)
	}
	return guo
}

// ClearIntro clears the value of the "intro" field.
func (guo *GroupUpdateOne) ClearIntro() *GroupUpdateOne {
	guo.mutation.ClearIntro()
	return guo
}

// SetKeys sets the "keys" field.
func (guo *GroupUpdateOne) SetKeys(s string) *GroupUpdateOne {
	guo.mutation.SetKeys(s)
	return guo
}

// SetOwner sets the "owner" edge to the Member entity.
func (guo *GroupUpdateOne) SetOwner(m *Member) *GroupUpdateOne {
	return guo.SetOwnerID(m.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (guo *GroupUpdateOne) AddMessageIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.AddMessageIDs(ids...)
	return guo
}

// AddMessages adds the "messages" edges to the Message entity.
func (guo *GroupUpdateOne) AddMessages(m ...*Message) *GroupUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the Member entity by IDs.
func (guo *GroupUpdateOne) AddMemberIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.AddMemberIDs(ids...)
	return guo
}

// AddMembers adds the "members" edges to the Member entity.
func (guo *GroupUpdateOne) AddMembers(m ...*Member) *GroupUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMemberIDs(ids...)
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by IDs.
func (guo *GroupUpdateOne) AddGroupMemberIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.AddGroupMemberIDs(ids...)
	return guo
}

// AddGroupMembers adds the "group_members" edges to the GroupMember entity.
func (guo *GroupUpdateOne) AddGroupMembers(g ...*GroupMember) *GroupUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddGroupMemberIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearOwner clears the "owner" edge to the Member entity.
func (guo *GroupUpdateOne) ClearOwner() *GroupUpdateOne {
	guo.mutation.ClearOwner()
	return guo
}

// ClearMessages clears all "messages" edges to the Message entity.
func (guo *GroupUpdateOne) ClearMessages() *GroupUpdateOne {
	guo.mutation.ClearMessages()
	return guo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (guo *GroupUpdateOne) RemoveMessageIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.RemoveMessageIDs(ids...)
	return guo
}

// RemoveMessages removes "messages" edges to Message entities.
func (guo *GroupUpdateOne) RemoveMessages(m ...*Message) *GroupUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMessageIDs(ids...)
}

// ClearMembers clears all "members" edges to the Member entity.
func (guo *GroupUpdateOne) ClearMembers() *GroupUpdateOne {
	guo.mutation.ClearMembers()
	return guo
}

// RemoveMemberIDs removes the "members" edge to Member entities by IDs.
func (guo *GroupUpdateOne) RemoveMemberIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.RemoveMemberIDs(ids...)
	return guo
}

// RemoveMembers removes "members" edges to Member entities.
func (guo *GroupUpdateOne) RemoveMembers(m ...*Member) *GroupUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMemberIDs(ids...)
}

// ClearGroupMembers clears all "group_members" edges to the GroupMember entity.
func (guo *GroupUpdateOne) ClearGroupMembers() *GroupUpdateOne {
	guo.mutation.ClearGroupMembers()
	return guo
}

// RemoveGroupMemberIDs removes the "group_members" edge to GroupMember entities by IDs.
func (guo *GroupUpdateOne) RemoveGroupMemberIDs(ids ...uint64) *GroupUpdateOne {
	guo.mutation.RemoveGroupMemberIDs(ids...)
	return guo
}

// RemoveGroupMembers removes "group_members" edges to GroupMember entities.
func (guo *GroupUpdateOne) RemoveGroupMembers(g ...*GroupMember) *GroupUpdateOne {
	ids := make([]uint64, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveGroupMemberIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Group)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GroupMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if _, ok := guo.mutation.OwnerID(); guo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.owner"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (guo *GroupUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupUpdateOne {
	guo.modifiers = append(guo.modifiers, modifiers...)
	return guo
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: group.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
	}
	if value, ok := guo.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldCategory,
		})
	}
	if value, ok := guo.mutation.MembersMax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
	}
	if value, ok := guo.mutation.AddedMembersMax(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
	}
	if value, ok := guo.mutation.MembersCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
	}
	if value, ok := guo.mutation.AddedMembersCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
	}
	if value, ok := guo.mutation.Public(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldPublic,
		})
	}
	if value, ok := guo.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldIntro,
		})
	}
	if guo.mutation.IntroCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: group.FieldIntro,
		})
	}
	if value, ok := guo.mutation.Keys(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldKeys,
		})
	}
	if guo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if nodes := guo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !guo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if nodes := guo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if guo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		createE := &GroupMemberCreate{config: guo.config, mutation: newGroupMemberMutation(guo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !guo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMemberCreate{config: guo.config, mutation: newGroupMemberMutation(guo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &GroupMemberCreate{config: guo.config, mutation: newGroupMemberMutation(guo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.GroupMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedGroupMembersIDs(); len(nodes) > 0 && !guo.mutation.GroupMembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.GroupMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: groupmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = guo.modifiers
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
