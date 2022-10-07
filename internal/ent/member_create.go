// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MemberCreate) SetCreatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetAddress sets the "address" field.
func (mc *MemberCreate) SetAddress(s string) *MemberCreate {
	mc.mutation.SetAddress(s)
	return mc
}

// SetNickname sets the "nickname" field.
func (mc *MemberCreate) SetNickname(s string) *MemberCreate {
	mc.mutation.SetNickname(s)
	return mc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mc *MemberCreate) SetNillableNickname(s *string) *MemberCreate {
	if s != nil {
		mc.SetNickname(*s)
	}
	return mc
}

// SetAvatar sets the "avatar" field.
func (mc *MemberCreate) SetAvatar(s string) *MemberCreate {
	mc.mutation.SetAvatar(s)
	return mc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mc *MemberCreate) SetNillableAvatar(s *string) *MemberCreate {
	if s != nil {
		mc.SetAvatar(*s)
	}
	return mc
}

// SetIntro sets the "intro" field.
func (mc *MemberCreate) SetIntro(s string) *MemberCreate {
	mc.mutation.SetIntro(s)
	return mc
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (mc *MemberCreate) SetNillableIntro(s *string) *MemberCreate {
	if s != nil {
		mc.SetIntro(*s)
	}
	return mc
}

// SetPublicKey sets the "public_key" field.
func (mc *MemberCreate) SetPublicKey(s string) *MemberCreate {
	mc.mutation.SetPublicKey(s)
	return mc
}

// SetNillablePublicKey sets the "public_key" field if the given value is not nil.
func (mc *MemberCreate) SetNillablePublicKey(s *string) *MemberCreate {
	if s != nil {
		mc.SetPublicKey(*s)
	}
	return mc
}

// SetNonce sets the "nonce" field.
func (mc *MemberCreate) SetNonce(s string) *MemberCreate {
	mc.mutation.SetNonce(s)
	return mc
}

// SetShowNickname sets the "show_nickname" field.
func (mc *MemberCreate) SetShowNickname(b bool) *MemberCreate {
	mc.mutation.SetShowNickname(b)
	return mc
}

// SetNillableShowNickname sets the "show_nickname" field if the given value is not nil.
func (mc *MemberCreate) SetNillableShowNickname(b *bool) *MemberCreate {
	if b != nil {
		mc.SetShowNickname(*b)
	}
	return mc
}

// SetLastNode sets the "last_node" field.
func (mc *MemberCreate) SetLastNode(i int64) *MemberCreate {
	mc.mutation.SetLastNode(i)
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(s string) *MemberCreate {
	mc.mutation.SetID(s)
	return mc
}

// AddOwnGroupIDs adds the "own_groups" edge to the Group entity by IDs.
func (mc *MemberCreate) AddOwnGroupIDs(ids ...string) *MemberCreate {
	mc.mutation.AddOwnGroupIDs(ids...)
	return mc
}

// AddOwnGroups adds the "own_groups" edges to the Group entity.
func (mc *MemberCreate) AddOwnGroups(g ...*Group) *MemberCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mc.AddOwnGroupIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (mc *MemberCreate) AddMessageIDs(ids ...string) *MemberCreate {
	mc.mutation.AddMessageIDs(ids...)
	return mc
}

// AddMessages adds the "messages" edges to the Message entity.
func (mc *MemberCreate) AddMessages(m ...*Message) *MemberCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMessageIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (mc *MemberCreate) AddGroupIDs(ids ...string) *MemberCreate {
	mc.mutation.AddGroupIDs(ids...)
	return mc
}

// AddGroups adds the "groups" edges to the Group entity.
func (mc *MemberCreate) AddGroups(g ...*Group) *MemberCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mc.AddGroupIDs(ids...)
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by IDs.
func (mc *MemberCreate) AddGroupMemberIDs(ids ...string) *MemberCreate {
	mc.mutation.AddGroupMemberIDs(ids...)
	return mc
}

// AddGroupMembers adds the "group_members" edges to the GroupMember entity.
func (mc *MemberCreate) AddGroupMembers(g ...*GroupMember) *MemberCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mc.AddGroupMemberIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	var (
		err  error
		node *Member
	)
	if err := mc.defaults(); err != nil {
		return nil, err
	}
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MemberCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		if member.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized member.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := member.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.ShowNickname(); !ok {
		v := member.DefaultShowNickname
		mc.mutation.SetShowNickname(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Member.created_at"`)}
	}
	if _, ok := mc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Member.address"`)}
	}
	if _, ok := mc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`ent: missing required field "Member.nonce"`)}
	}
	if _, ok := mc.mutation.ShowNickname(); !ok {
		return &ValidationError{Name: "show_nickname", err: errors.New(`ent: missing required field "Member.show_nickname"`)}
	}
	if _, ok := mc.mutation.LastNode(); !ok {
		return &ValidationError{Name: "last_node", err: errors.New(`ent: missing required field "Member.last_node"`)}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := member.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Member.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Member.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: member.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: member.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := mc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := mc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := mc.mutation.Intro(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldIntro,
		})
		_node.Intro = value
	}
	if value, ok := mc.mutation.PublicKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldPublicKey,
		})
		_node.PublicKey = value
	}
	if value, ok := mc.mutation.Nonce(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNonce,
		})
		_node.Nonce = value
	}
	if value, ok := mc.mutation.ShowNickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: member.FieldShowNickname,
		})
		_node.ShowNickname = value
	}
	if value, ok := mc.mutation.LastNode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: member.FieldLastNode,
		})
		_node.LastNode = value
	}
	if nodes := mc.mutation.OwnGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
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
		createE := &GroupMemberCreate{config: mc.config, mutation: newGroupMemberMutation(mc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.GroupMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   member.GroupMembersTable,
			Columns: []string{member.GroupMembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: groupmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MemberCreate) OnConflict(opts ...sql.ConflictOption) *MemberUpsertOne {
	mc.conflict = opts
	return &MemberUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MemberCreate) OnConflictColumns(columns ...string) *MemberUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertOne{
		create: mc,
	}
}

type (
	// MemberUpsertOne is the builder for "upsert"-ing
	//  one Member node.
	MemberUpsertOne struct {
		create *MemberCreate
	}

	// MemberUpsert is the "OnConflict" setter.
	MemberUpsert struct {
		*sql.UpdateSet
	}
)

// SetAddress sets the "address" field.
func (u *MemberUpsert) SetAddress(v string) *MemberUpsert {
	u.Set(member.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *MemberUpsert) UpdateAddress() *MemberUpsert {
	u.SetExcluded(member.FieldAddress)
	return u
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsert) SetNickname(v string) *MemberUpsert {
	u.Set(member.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsert) UpdateNickname() *MemberUpsert {
	u.SetExcluded(member.FieldNickname)
	return u
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsert) ClearNickname() *MemberUpsert {
	u.SetNull(member.FieldNickname)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsert) SetAvatar(v string) *MemberUpsert {
	u.Set(member.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsert) UpdateAvatar() *MemberUpsert {
	u.SetExcluded(member.FieldAvatar)
	return u
}

// ClearAvatar clears the value of the "avatar" field.
func (u *MemberUpsert) ClearAvatar() *MemberUpsert {
	u.SetNull(member.FieldAvatar)
	return u
}

// SetIntro sets the "intro" field.
func (u *MemberUpsert) SetIntro(v string) *MemberUpsert {
	u.Set(member.FieldIntro, v)
	return u
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *MemberUpsert) UpdateIntro() *MemberUpsert {
	u.SetExcluded(member.FieldIntro)
	return u
}

// ClearIntro clears the value of the "intro" field.
func (u *MemberUpsert) ClearIntro() *MemberUpsert {
	u.SetNull(member.FieldIntro)
	return u
}

// SetPublicKey sets the "public_key" field.
func (u *MemberUpsert) SetPublicKey(v string) *MemberUpsert {
	u.Set(member.FieldPublicKey, v)
	return u
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *MemberUpsert) UpdatePublicKey() *MemberUpsert {
	u.SetExcluded(member.FieldPublicKey)
	return u
}

// ClearPublicKey clears the value of the "public_key" field.
func (u *MemberUpsert) ClearPublicKey() *MemberUpsert {
	u.SetNull(member.FieldPublicKey)
	return u
}

// SetNonce sets the "nonce" field.
func (u *MemberUpsert) SetNonce(v string) *MemberUpsert {
	u.Set(member.FieldNonce, v)
	return u
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *MemberUpsert) UpdateNonce() *MemberUpsert {
	u.SetExcluded(member.FieldNonce)
	return u
}

// SetShowNickname sets the "show_nickname" field.
func (u *MemberUpsert) SetShowNickname(v bool) *MemberUpsert {
	u.Set(member.FieldShowNickname, v)
	return u
}

// UpdateShowNickname sets the "show_nickname" field to the value that was provided on create.
func (u *MemberUpsert) UpdateShowNickname() *MemberUpsert {
	u.SetExcluded(member.FieldShowNickname)
	return u
}

// SetLastNode sets the "last_node" field.
func (u *MemberUpsert) SetLastNode(v int64) *MemberUpsert {
	u.Set(member.FieldLastNode, v)
	return u
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *MemberUpsert) UpdateLastNode() *MemberUpsert {
	u.SetExcluded(member.FieldLastNode)
	return u
}

// AddLastNode adds v to the "last_node" field.
func (u *MemberUpsert) AddLastNode(v int64) *MemberUpsert {
	u.Add(member.FieldLastNode, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertOne) UpdateNewValues() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(member.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(member.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MemberUpsertOne) Ignore() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertOne) DoNothing() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreate.OnConflict
// documentation for more info.
func (u *MemberUpsertOne) Update(set func(*MemberUpsert)) *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddress sets the "address" field.
func (u *MemberUpsertOne) SetAddress(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateAddress() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAddress()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertOne) SetNickname(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsertOne) ClearNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearNickname()
	})
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsertOne) SetAvatar(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateAvatar() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAvatar()
	})
}

// ClearAvatar clears the value of the "avatar" field.
func (u *MemberUpsertOne) ClearAvatar() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearAvatar()
	})
}

// SetIntro sets the "intro" field.
func (u *MemberUpsertOne) SetIntro(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateIntro() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateIntro()
	})
}

// ClearIntro clears the value of the "intro" field.
func (u *MemberUpsertOne) ClearIntro() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearIntro()
	})
}

// SetPublicKey sets the "public_key" field.
func (u *MemberUpsertOne) SetPublicKey(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdatePublicKey() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePublicKey()
	})
}

// ClearPublicKey clears the value of the "public_key" field.
func (u *MemberUpsertOne) ClearPublicKey() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearPublicKey()
	})
}

// SetNonce sets the "nonce" field.
func (u *MemberUpsertOne) SetNonce(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateNonce() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNonce()
	})
}

// SetShowNickname sets the "show_nickname" field.
func (u *MemberUpsertOne) SetShowNickname(v bool) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetShowNickname(v)
	})
}

// UpdateShowNickname sets the "show_nickname" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateShowNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateShowNickname()
	})
}

// SetLastNode sets the "last_node" field.
func (u *MemberUpsertOne) SetLastNode(v int64) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetLastNode(v)
	})
}

// AddLastNode adds v to the "last_node" field.
func (u *MemberUpsertOne) AddLastNode(v int64) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.AddLastNode(v)
	})
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateLastNode() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateLastNode()
	})
}

// Exec executes the query.
func (u *MemberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MemberUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MemberUpsertOne.ID is not supported by MySQL driver. Use MemberUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MemberUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	builders []*MemberCreate
	conflict []sql.ConflictOption
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflict(opts ...sql.ConflictOption) *MemberUpsertBulk {
	mcb.conflict = opts
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflictColumns(columns ...string) *MemberUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// MemberUpsertBulk is the builder for "upsert"-ing
// a bulk of Member nodes.
type MemberUpsertBulk struct {
	create *MemberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertBulk) UpdateNewValues() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(member.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(member.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MemberUpsertBulk) Ignore() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertBulk) DoNothing() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreateBulk.OnConflict
// documentation for more info.
func (u *MemberUpsertBulk) Update(set func(*MemberUpsert)) *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetAddress sets the "address" field.
func (u *MemberUpsertBulk) SetAddress(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateAddress() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAddress()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertBulk) SetNickname(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsertBulk) ClearNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearNickname()
	})
}

// SetAvatar sets the "avatar" field.
func (u *MemberUpsertBulk) SetAvatar(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateAvatar() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateAvatar()
	})
}

// ClearAvatar clears the value of the "avatar" field.
func (u *MemberUpsertBulk) ClearAvatar() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearAvatar()
	})
}

// SetIntro sets the "intro" field.
func (u *MemberUpsertBulk) SetIntro(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateIntro() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateIntro()
	})
}

// ClearIntro clears the value of the "intro" field.
func (u *MemberUpsertBulk) ClearIntro() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearIntro()
	})
}

// SetPublicKey sets the "public_key" field.
func (u *MemberUpsertBulk) SetPublicKey(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetPublicKey(v)
	})
}

// UpdatePublicKey sets the "public_key" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdatePublicKey() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePublicKey()
	})
}

// ClearPublicKey clears the value of the "public_key" field.
func (u *MemberUpsertBulk) ClearPublicKey() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearPublicKey()
	})
}

// SetNonce sets the "nonce" field.
func (u *MemberUpsertBulk) SetNonce(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateNonce() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNonce()
	})
}

// SetShowNickname sets the "show_nickname" field.
func (u *MemberUpsertBulk) SetShowNickname(v bool) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetShowNickname(v)
	})
}

// UpdateShowNickname sets the "show_nickname" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateShowNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateShowNickname()
	})
}

// SetLastNode sets the "last_node" field.
func (u *MemberUpsertBulk) SetLastNode(v int64) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetLastNode(v)
	})
}

// AddLastNode adds v to the "last_node" field.
func (u *MemberUpsertBulk) AddLastNode(v int64) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.AddLastNode(v)
	})
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateLastNode() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateLastNode()
	})
}

// Exec executes the query.
func (u *MemberUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MemberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
