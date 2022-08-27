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
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
)

// GroupMemberCreate is the builder for creating a GroupMember entity.
type GroupMemberCreate struct {
	config
	mutation *GroupMemberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gmc *GroupMemberCreate) SetCreatedAt(t time.Time) *GroupMemberCreate {
	gmc.mutation.SetCreatedAt(t)
	return gmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gmc *GroupMemberCreate) SetNillableCreatedAt(t *time.Time) *GroupMemberCreate {
	if t != nil {
		gmc.SetCreatedAt(*t)
	}
	return gmc
}

// SetMemberID sets the "member_id" field.
func (gmc *GroupMemberCreate) SetMemberID(u uint64) *GroupMemberCreate {
	gmc.mutation.SetMemberID(u)
	return gmc
}

// SetGroupID sets the "group_id" field.
func (gmc *GroupMemberCreate) SetGroupID(u uint64) *GroupMemberCreate {
	gmc.mutation.SetGroupID(u)
	return gmc
}

// SetKeyID sets the "key_id" field.
func (gmc *GroupMemberCreate) SetKeyID(u uint64) *GroupMemberCreate {
	gmc.mutation.SetKeyID(u)
	return gmc
}

// SetPermission sets the "permission" field.
func (gmc *GroupMemberCreate) SetPermission(u uint8) *GroupMemberCreate {
	gmc.mutation.SetPermission(u)
	return gmc
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (gmc *GroupMemberCreate) SetNillablePermission(u *uint8) *GroupMemberCreate {
	if u != nil {
		gmc.SetPermission(*u)
	}
	return gmc
}

// SetSn sets the "sn" field.
func (gmc *GroupMemberCreate) SetSn(s string) *GroupMemberCreate {
	gmc.mutation.SetSn(s)
	return gmc
}

// SetID sets the "id" field.
func (gmc *GroupMemberCreate) SetID(u uint64) *GroupMemberCreate {
	gmc.mutation.SetID(u)
	return gmc
}

// SetMember sets the "member" edge to the Member entity.
func (gmc *GroupMemberCreate) SetMember(m *Member) *GroupMemberCreate {
	return gmc.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gmc *GroupMemberCreate) SetGroup(g *Group) *GroupMemberCreate {
	return gmc.SetGroupID(g.ID)
}

// SetKey sets the "key" edge to the Key entity.
func (gmc *GroupMemberCreate) SetKey(k *Key) *GroupMemberCreate {
	return gmc.SetKeyID(k.ID)
}

// Mutation returns the GroupMemberMutation object of the builder.
func (gmc *GroupMemberCreate) Mutation() *GroupMemberMutation {
	return gmc.mutation
}

// Save creates the GroupMember in the database.
func (gmc *GroupMemberCreate) Save(ctx context.Context) (*GroupMember, error) {
	var (
		err  error
		node *GroupMember
	)
	if err := gmc.defaults(); err != nil {
		return nil, err
	}
	if len(gmc.hooks) == 0 {
		if err = gmc.check(); err != nil {
			return nil, err
		}
		node, err = gmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmc.check(); err != nil {
				return nil, err
			}
			gmc.mutation = mutation
			if node, err = gmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gmc.hooks) - 1; i >= 0; i-- {
			if gmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gmc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMemberCreate) SaveX(ctx context.Context) *GroupMember {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMemberCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMemberCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmc *GroupMemberCreate) defaults() error {
	if _, ok := gmc.mutation.CreatedAt(); !ok {
		v := groupmember.DefaultCreatedAt
		gmc.mutation.SetCreatedAt(v)
	}
	if _, ok := gmc.mutation.Permission(); !ok {
		v := groupmember.DefaultPermission
		gmc.mutation.SetPermission(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMemberCreate) check() error {
	if _, ok := gmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GroupMember.created_at"`)}
	}
	if _, ok := gmc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "GroupMember.member_id"`)}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "GroupMember.group_id"`)}
	}
	if _, ok := gmc.mutation.KeyID(); !ok {
		return &ValidationError{Name: "key_id", err: errors.New(`ent: missing required field "GroupMember.key_id"`)}
	}
	if _, ok := gmc.mutation.Permission(); !ok {
		return &ValidationError{Name: "permission", err: errors.New(`ent: missing required field "GroupMember.permission"`)}
	}
	if _, ok := gmc.mutation.Sn(); !ok {
		return &ValidationError{Name: "sn", err: errors.New(`ent: missing required field "GroupMember.sn"`)}
	}
	if _, ok := gmc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member", err: errors.New(`ent: missing required edge "GroupMember.member"`)}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupMember.group"`)}
	}
	if _, ok := gmc.mutation.KeyID(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required edge "GroupMember.key"`)}
	}
	return nil
}

func (gmc *GroupMemberCreate) sqlSave(ctx context.Context) (*GroupMember, error) {
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (gmc *GroupMemberCreate) createSpec() (*GroupMember, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMember{config: gmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupmember.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: groupmember.FieldID,
			},
		}
	)
	_spec.OnConflict = gmc.conflict
	if id, ok := gmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: groupmember.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gmc.mutation.Permission(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: groupmember.FieldPermission,
		})
		_node.Permission = value
	}
	if value, ok := gmc.mutation.Sn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupmember.FieldSn,
		})
		_node.Sn = value
	}
	if nodes := gmc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.MemberTable,
			Columns: []string{groupmember.MemberColumn},
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
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.GroupTable,
			Columns: []string{groupmember.GroupColumn},
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
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmember.KeyTable,
			Columns: []string{groupmember.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: key.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.KeyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupMember.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupMemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gmc *GroupMemberCreate) OnConflict(opts ...sql.ConflictOption) *GroupMemberUpsertOne {
	gmc.conflict = opts
	return &GroupMemberUpsertOne{
		create: gmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gmc *GroupMemberCreate) OnConflictColumns(columns ...string) *GroupMemberUpsertOne {
	gmc.conflict = append(gmc.conflict, sql.ConflictColumns(columns...))
	return &GroupMemberUpsertOne{
		create: gmc,
	}
}

type (
	// GroupMemberUpsertOne is the builder for "upsert"-ing
	//  one GroupMember node.
	GroupMemberUpsertOne struct {
		create *GroupMemberCreate
	}

	// GroupMemberUpsert is the "OnConflict" setter.
	GroupMemberUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GroupMemberUpsert) SetCreatedAt(v time.Time) *GroupMemberUpsert {
	u.Set(groupmember.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdateCreatedAt() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldCreatedAt)
	return u
}

// SetMemberID sets the "member_id" field.
func (u *GroupMemberUpsert) SetMemberID(v uint64) *GroupMemberUpsert {
	u.Set(groupmember.FieldMemberID, v)
	return u
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdateMemberID() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldMemberID)
	return u
}

// SetGroupID sets the "group_id" field.
func (u *GroupMemberUpsert) SetGroupID(v uint64) *GroupMemberUpsert {
	u.Set(groupmember.FieldGroupID, v)
	return u
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdateGroupID() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldGroupID)
	return u
}

// SetKeyID sets the "key_id" field.
func (u *GroupMemberUpsert) SetKeyID(v uint64) *GroupMemberUpsert {
	u.Set(groupmember.FieldKeyID, v)
	return u
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdateKeyID() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldKeyID)
	return u
}

// SetPermission sets the "permission" field.
func (u *GroupMemberUpsert) SetPermission(v uint8) *GroupMemberUpsert {
	u.Set(groupmember.FieldPermission, v)
	return u
}

// UpdatePermission sets the "permission" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdatePermission() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldPermission)
	return u
}

// AddPermission adds v to the "permission" field.
func (u *GroupMemberUpsert) AddPermission(v uint8) *GroupMemberUpsert {
	u.Add(groupmember.FieldPermission, v)
	return u
}

// SetSn sets the "sn" field.
func (u *GroupMemberUpsert) SetSn(v string) *GroupMemberUpsert {
	u.Set(groupmember.FieldSn, v)
	return u
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupMemberUpsert) UpdateSn() *GroupMemberUpsert {
	u.SetExcluded(groupmember.FieldSn)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(groupmember.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupMemberUpsertOne) UpdateNewValues() *GroupMemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(groupmember.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(groupmember.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupMemberUpsertOne) Ignore() *GroupMemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupMemberUpsertOne) DoNothing() *GroupMemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupMemberCreate.OnConflict
// documentation for more info.
func (u *GroupMemberUpsertOne) Update(set func(*GroupMemberUpsert)) *GroupMemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupMemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GroupMemberUpsertOne) SetCreatedAt(v time.Time) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdateCreatedAt() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetMemberID sets the "member_id" field.
func (u *GroupMemberUpsertOne) SetMemberID(v uint64) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdateMemberID() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateMemberID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *GroupMemberUpsertOne) SetGroupID(v uint64) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdateGroupID() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateGroupID()
	})
}

// SetKeyID sets the "key_id" field.
func (u *GroupMemberUpsertOne) SetKeyID(v uint64) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetKeyID(v)
	})
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdateKeyID() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateKeyID()
	})
}

// SetPermission sets the "permission" field.
func (u *GroupMemberUpsertOne) SetPermission(v uint8) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetPermission(v)
	})
}

// AddPermission adds v to the "permission" field.
func (u *GroupMemberUpsertOne) AddPermission(v uint8) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.AddPermission(v)
	})
}

// UpdatePermission sets the "permission" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdatePermission() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdatePermission()
	})
}

// SetSn sets the "sn" field.
func (u *GroupMemberUpsertOne) SetSn(v string) *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupMemberUpsertOne) UpdateSn() *GroupMemberUpsertOne {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateSn()
	})
}

// Exec executes the query.
func (u *GroupMemberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupMemberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupMemberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupMemberUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupMemberUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupMemberCreateBulk is the builder for creating many GroupMember entities in bulk.
type GroupMemberCreateBulk struct {
	config
	builders []*GroupMemberCreate
	conflict []sql.ConflictOption
}

// Save creates the GroupMember entities in the database.
func (gmcb *GroupMemberCreateBulk) Save(ctx context.Context) ([]*GroupMember, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMember, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMemberMutation)
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
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMemberCreateBulk) SaveX(ctx context.Context) []*GroupMember {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMemberCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMemberCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupMember.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupMemberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gmcb *GroupMemberCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupMemberUpsertBulk {
	gmcb.conflict = opts
	return &GroupMemberUpsertBulk{
		create: gmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gmcb *GroupMemberCreateBulk) OnConflictColumns(columns ...string) *GroupMemberUpsertBulk {
	gmcb.conflict = append(gmcb.conflict, sql.ConflictColumns(columns...))
	return &GroupMemberUpsertBulk{
		create: gmcb,
	}
}

// GroupMemberUpsertBulk is the builder for "upsert"-ing
// a bulk of GroupMember nodes.
type GroupMemberUpsertBulk struct {
	create *GroupMemberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(groupmember.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupMemberUpsertBulk) UpdateNewValues() *GroupMemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(groupmember.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(groupmember.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupMember.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupMemberUpsertBulk) Ignore() *GroupMemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupMemberUpsertBulk) DoNothing() *GroupMemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupMemberCreateBulk.OnConflict
// documentation for more info.
func (u *GroupMemberUpsertBulk) Update(set func(*GroupMemberUpsert)) *GroupMemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupMemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GroupMemberUpsertBulk) SetCreatedAt(v time.Time) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdateCreatedAt() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetMemberID sets the "member_id" field.
func (u *GroupMemberUpsertBulk) SetMemberID(v uint64) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdateMemberID() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateMemberID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *GroupMemberUpsertBulk) SetGroupID(v uint64) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdateGroupID() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateGroupID()
	})
}

// SetKeyID sets the "key_id" field.
func (u *GroupMemberUpsertBulk) SetKeyID(v uint64) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetKeyID(v)
	})
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdateKeyID() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateKeyID()
	})
}

// SetPermission sets the "permission" field.
func (u *GroupMemberUpsertBulk) SetPermission(v uint8) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetPermission(v)
	})
}

// AddPermission adds v to the "permission" field.
func (u *GroupMemberUpsertBulk) AddPermission(v uint8) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.AddPermission(v)
	})
}

// UpdatePermission sets the "permission" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdatePermission() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdatePermission()
	})
}

// SetSn sets the "sn" field.
func (u *GroupMemberUpsertBulk) SetSn(v string) *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *GroupMemberUpsertBulk) UpdateSn() *GroupMemberUpsertBulk {
	return u.Update(func(s *GroupMemberUpsert) {
		s.UpdateSn()
	})
}

// Exec executes the query.
func (u *GroupMemberUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupMemberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupMemberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupMemberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
