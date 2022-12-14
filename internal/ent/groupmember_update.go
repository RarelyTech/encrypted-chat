// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/app/model"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// GroupMemberUpdate is the builder for updating GroupMember entities.
type GroupMemberUpdate struct {
	config
	hooks     []Hook
	mutation  *GroupMemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GroupMemberUpdate builder.
func (gmu *GroupMemberUpdate) Where(ps ...predicate.GroupMember) *GroupMemberUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetMemberID sets the "member_id" field.
func (gmu *GroupMemberUpdate) SetMemberID(s string) *GroupMemberUpdate {
	gmu.mutation.SetMemberID(s)
	return gmu
}

// SetGroupID sets the "group_id" field.
func (gmu *GroupMemberUpdate) SetGroupID(s string) *GroupMemberUpdate {
	gmu.mutation.SetGroupID(s)
	return gmu
}

// SetPermission sets the "permission" field.
func (gmu *GroupMemberUpdate) SetPermission(mmp model.GroupMemberPerm) *GroupMemberUpdate {
	gmu.mutation.SetPermission(mmp)
	return gmu
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillablePermission(mmp *model.GroupMemberPerm) *GroupMemberUpdate {
	if mmp != nil {
		gmu.SetPermission(*mmp)
	}
	return gmu
}

// SetInviterID sets the "inviter_id" field.
func (gmu *GroupMemberUpdate) SetInviterID(s string) *GroupMemberUpdate {
	gmu.mutation.SetInviterID(s)
	return gmu
}

// SetNillableInviterID sets the "inviter_id" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableInviterID(s *string) *GroupMemberUpdate {
	if s != nil {
		gmu.SetInviterID(*s)
	}
	return gmu
}

// ClearInviterID clears the value of the "inviter_id" field.
func (gmu *GroupMemberUpdate) ClearInviterID() *GroupMemberUpdate {
	gmu.mutation.ClearInviterID()
	return gmu
}

// SetInviteCode sets the "invite_code" field.
func (gmu *GroupMemberUpdate) SetInviteCode(s string) *GroupMemberUpdate {
	gmu.mutation.SetInviteCode(s)
	return gmu
}

// SetInviteExpire sets the "invite_expire" field.
func (gmu *GroupMemberUpdate) SetInviteExpire(t time.Time) *GroupMemberUpdate {
	gmu.mutation.SetInviteExpire(t)
	return gmu
}

// SetReadID sets the "read_id" field.
func (gmu *GroupMemberUpdate) SetReadID(s string) *GroupMemberUpdate {
	gmu.mutation.SetReadID(s)
	return gmu
}

// SetNillableReadID sets the "read_id" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableReadID(s *string) *GroupMemberUpdate {
	if s != nil {
		gmu.SetReadID(*s)
	}
	return gmu
}

// ClearReadID clears the value of the "read_id" field.
func (gmu *GroupMemberUpdate) ClearReadID() *GroupMemberUpdate {
	gmu.mutation.ClearReadID()
	return gmu
}

// SetReadTime sets the "read_time" field.
func (gmu *GroupMemberUpdate) SetReadTime(t time.Time) *GroupMemberUpdate {
	gmu.mutation.SetReadTime(t)
	return gmu
}

// SetNillableReadTime sets the "read_time" field if the given value is not nil.
func (gmu *GroupMemberUpdate) SetNillableReadTime(t *time.Time) *GroupMemberUpdate {
	if t != nil {
		gmu.SetReadTime(*t)
	}
	return gmu
}

// ClearReadTime clears the value of the "read_time" field.
func (gmu *GroupMemberUpdate) ClearReadTime() *GroupMemberUpdate {
	gmu.mutation.ClearReadTime()
	return gmu
}

// SetMember sets the "member" edge to the Member entity.
func (gmu *GroupMemberUpdate) SetMember(m *Member) *GroupMemberUpdate {
	return gmu.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gmu *GroupMemberUpdate) SetGroup(g *Group) *GroupMemberUpdate {
	return gmu.SetGroupID(g.ID)
}

// SetInviter sets the "inviter" edge to the Member entity.
func (gmu *GroupMemberUpdate) SetInviter(m *Member) *GroupMemberUpdate {
	return gmu.SetInviterID(m.ID)
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmu *GroupMemberUpdate) Mutation() *GroupMemberMutation {
	return gmu.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (gmu *GroupMemberUpdate) ClearMember() *GroupMemberUpdate {
	gmu.mutation.ClearMember()
	return gmu
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmu *GroupMemberUpdate) ClearGroup() *GroupMemberUpdate {
	gmu.mutation.ClearGroup()
	return gmu
}

// ClearInviter clears the "inviter" edge to the Member entity.
func (gmu *GroupMemberUpdate) ClearInviter() *GroupMemberUpdate {
	gmu.mutation.ClearInviter()
	return gmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GroupMemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gmu.hooks) == 0 {
		if err = gmu.check(); err != nil {
			return 0, err
		}
		affected, err = gmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmu.check(); err != nil {
				return 0, err
			}
			gmu.mutation = mutation
			affected, err = gmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gmu.hooks) - 1; i >= 0; i-- {
			if gmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GroupMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GroupMemberUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GroupMemberUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmu *GroupMemberUpdate) check() error {
	if _, ok := gmu.mutation.MemberID(); gmu.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.member"`)
	}
	if _, ok := gmu.mutation.GroupID(); gmu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmu *GroupMemberUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupMemberUpdate {
	gmu.modifiers = append(gmu.modifiers, modifiers...)
	return gmu
}

func (gmu *GroupMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupmember.Table,
			Columns: groupmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmember.FieldID,
			},
		},
	}
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.Permission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: groupmember.FieldPermission,
		})
	}
	if value, ok := gmu.mutation.InviteCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupmember.FieldInviteCode,
		})
	}
	if value, ok := gmu.mutation.InviteExpire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupmember.FieldInviteExpire,
		})
	}
	if value, ok := gmu.mutation.ReadID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupmember.FieldReadID,
		})
	}
	if gmu.mutation.ReadIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: groupmember.FieldReadID,
		})
	}
	if value, ok := gmu.mutation.ReadTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupmember.FieldReadTime,
		})
	}
	if gmu.mutation.ReadTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: groupmember.FieldReadTime,
		})
	}
	if gmu.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.MemberTable,
			Columns: []string{groupmember.MemberColumn},
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
	if nodes := gmu.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.MemberTable,
			Columns: []string{groupmember.MemberColumn},
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
	if gmu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.GroupTable,
			Columns: []string{groupmember.GroupColumn},
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
	if nodes := gmu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.GroupTable,
			Columns: []string{groupmember.GroupColumn},
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
	if gmu.mutation.InviterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.InviterTable,
			Columns: []string{groupmember.InviterColumn},
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
	if nodes := gmu.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.InviterTable,
			Columns: []string{groupmember.InviterColumn},
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
	_spec.Modifiers = gmu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GroupMemberUpdateOne is the builder for updating a single GroupMember entity.
type GroupMemberUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GroupMemberMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetMemberID sets the "member_id" field.
func (gmuo *GroupMemberUpdateOne) SetMemberID(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetMemberID(s)
	return gmuo
}

// SetGroupID sets the "group_id" field.
func (gmuo *GroupMemberUpdateOne) SetGroupID(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetGroupID(s)
	return gmuo
}

// SetPermission sets the "permission" field.
func (gmuo *GroupMemberUpdateOne) SetPermission(mmp model.GroupMemberPerm) *GroupMemberUpdateOne {
	gmuo.mutation.SetPermission(mmp)
	return gmuo
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillablePermission(mmp *model.GroupMemberPerm) *GroupMemberUpdateOne {
	if mmp != nil {
		gmuo.SetPermission(*mmp)
	}
	return gmuo
}

// SetInviterID sets the "inviter_id" field.
func (gmuo *GroupMemberUpdateOne) SetInviterID(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetInviterID(s)
	return gmuo
}

// SetNillableInviterID sets the "inviter_id" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableInviterID(s *string) *GroupMemberUpdateOne {
	if s != nil {
		gmuo.SetInviterID(*s)
	}
	return gmuo
}

// ClearInviterID clears the value of the "inviter_id" field.
func (gmuo *GroupMemberUpdateOne) ClearInviterID() *GroupMemberUpdateOne {
	gmuo.mutation.ClearInviterID()
	return gmuo
}

// SetInviteCode sets the "invite_code" field.
func (gmuo *GroupMemberUpdateOne) SetInviteCode(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetInviteCode(s)
	return gmuo
}

// SetInviteExpire sets the "invite_expire" field.
func (gmuo *GroupMemberUpdateOne) SetInviteExpire(t time.Time) *GroupMemberUpdateOne {
	gmuo.mutation.SetInviteExpire(t)
	return gmuo
}

// SetReadID sets the "read_id" field.
func (gmuo *GroupMemberUpdateOne) SetReadID(s string) *GroupMemberUpdateOne {
	gmuo.mutation.SetReadID(s)
	return gmuo
}

// SetNillableReadID sets the "read_id" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableReadID(s *string) *GroupMemberUpdateOne {
	if s != nil {
		gmuo.SetReadID(*s)
	}
	return gmuo
}

// ClearReadID clears the value of the "read_id" field.
func (gmuo *GroupMemberUpdateOne) ClearReadID() *GroupMemberUpdateOne {
	gmuo.mutation.ClearReadID()
	return gmuo
}

// SetReadTime sets the "read_time" field.
func (gmuo *GroupMemberUpdateOne) SetReadTime(t time.Time) *GroupMemberUpdateOne {
	gmuo.mutation.SetReadTime(t)
	return gmuo
}

// SetNillableReadTime sets the "read_time" field if the given value is not nil.
func (gmuo *GroupMemberUpdateOne) SetNillableReadTime(t *time.Time) *GroupMemberUpdateOne {
	if t != nil {
		gmuo.SetReadTime(*t)
	}
	return gmuo
}

// ClearReadTime clears the value of the "read_time" field.
func (gmuo *GroupMemberUpdateOne) ClearReadTime() *GroupMemberUpdateOne {
	gmuo.mutation.ClearReadTime()
	return gmuo
}

// SetMember sets the "member" edge to the Member entity.
func (gmuo *GroupMemberUpdateOne) SetMember(m *Member) *GroupMemberUpdateOne {
	return gmuo.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gmuo *GroupMemberUpdateOne) SetGroup(g *Group) *GroupMemberUpdateOne {
	return gmuo.SetGroupID(g.ID)
}

// SetInviter sets the "inviter" edge to the Member entity.
func (gmuo *GroupMemberUpdateOne) SetInviter(m *Member) *GroupMemberUpdateOne {
	return gmuo.SetInviterID(m.ID)
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmuo *GroupMemberUpdateOne) Mutation() *GroupMemberMutation {
	return gmuo.mutation
}

// ClearMember clears the "member" edge to the Member entity.
func (gmuo *GroupMemberUpdateOne) ClearMember() *GroupMemberUpdateOne {
	gmuo.mutation.ClearMember()
	return gmuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (gmuo *GroupMemberUpdateOne) ClearGroup() *GroupMemberUpdateOne {
	gmuo.mutation.ClearGroup()
	return gmuo
}

// ClearInviter clears the "inviter" edge to the Member entity.
func (gmuo *GroupMemberUpdateOne) ClearInviter() *GroupMemberUpdateOne {
	gmuo.mutation.ClearInviter()
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GroupMemberUpdateOne) Select(field string, fields ...string) *GroupMemberUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GroupMember entity.
func (gmuo *GroupMemberUpdateOne) Save(ctx context.Context) (*GroupMember, error) {
	var (
		err  error
		node *GroupMember
	)
	if len(gmuo.hooks) == 0 {
		if err = gmuo.check(); err != nil {
			return nil, err
		}
		node, err = gmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmuo.check(); err != nil {
				return nil, err
			}
			gmuo.mutation = mutation
			node, err = gmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gmuo.hooks) - 1; i >= 0; i-- {
			if gmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GroupMember)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GroupMemberMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GroupMemberUpdateOne) SaveX(ctx context.Context) *GroupMember {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GroupMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GroupMemberUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmuo *GroupMemberUpdateOne) check() error {
	if _, ok := gmuo.mutation.MemberID(); gmuo.mutation.MemberCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.member"`)
	}
	if _, ok := gmuo.mutation.GroupID(); gmuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupMember.group"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmuo *GroupMemberUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupMemberUpdateOne {
	gmuo.modifiers = append(gmuo.modifiers, modifiers...)
	return gmuo
}

func (gmuo *GroupMemberUpdateOne) sqlSave(ctx context.Context) (_node *GroupMember, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupmember.Table,
			Columns: groupmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmember.FieldID,
			},
		},
	}
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmember.FieldID)
		for _, f := range fields {
			if !groupmember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != groupmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.Permission(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: groupmember.FieldPermission,
		})
	}
	if value, ok := gmuo.mutation.InviteCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupmember.FieldInviteCode,
		})
	}
	if value, ok := gmuo.mutation.InviteExpire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupmember.FieldInviteExpire,
		})
	}
	if value, ok := gmuo.mutation.ReadID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupmember.FieldReadID,
		})
	}
	if gmuo.mutation.ReadIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: groupmember.FieldReadID,
		})
	}
	if value, ok := gmuo.mutation.ReadTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupmember.FieldReadTime,
		})
	}
	if gmuo.mutation.ReadTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: groupmember.FieldReadTime,
		})
	}
	if gmuo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.MemberTable,
			Columns: []string{groupmember.MemberColumn},
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
	if nodes := gmuo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.MemberTable,
			Columns: []string{groupmember.MemberColumn},
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
	if gmuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.GroupTable,
			Columns: []string{groupmember.GroupColumn},
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
	if nodes := gmuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.GroupTable,
			Columns: []string{groupmember.GroupColumn},
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
	if gmuo.mutation.InviterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.InviterTable,
			Columns: []string{groupmember.InviterColumn},
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
	if nodes := gmuo.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.InviterTable,
			Columns: []string{groupmember.InviterColumn},
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
	_spec.Modifiers = gmuo.modifiers
	_node = &GroupMember{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
