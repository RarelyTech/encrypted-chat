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
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/messageread"
)

// MessageReadCreate is the builder for creating a MessageRead entity.
type MessageReadCreate struct {
	config
	mutation *MessageReadMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetMemberID sets the "member_id" field.
func (mrc *MessageReadCreate) SetMemberID(s string) *MessageReadCreate {
	mrc.mutation.SetMemberID(s)
	return mrc
}

// SetGroupID sets the "group_id" field.
func (mrc *MessageReadCreate) SetGroupID(s string) *MessageReadCreate {
	mrc.mutation.SetGroupID(s)
	return mrc
}

// SetLastID sets the "last_id" field.
func (mrc *MessageReadCreate) SetLastID(s string) *MessageReadCreate {
	mrc.mutation.SetLastID(s)
	return mrc
}

// SetLastTime sets the "last_time" field.
func (mrc *MessageReadCreate) SetLastTime(t time.Time) *MessageReadCreate {
	mrc.mutation.SetLastTime(t)
	return mrc
}

// SetID sets the "id" field.
func (mrc *MessageReadCreate) SetID(s string) *MessageReadCreate {
	mrc.mutation.SetID(s)
	return mrc
}

// SetMember sets the "member" edge to the Member entity.
func (mrc *MessageReadCreate) SetMember(m *Member) *MessageReadCreate {
	return mrc.SetMemberID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (mrc *MessageReadCreate) SetGroup(g *Group) *MessageReadCreate {
	return mrc.SetGroupID(g.ID)
}

// Mutation returns the MessageReadMutation object of the builder.
func (mrc *MessageReadCreate) Mutation() *MessageReadMutation {
	return mrc.mutation
}

// Save creates the MessageRead in the database.
func (mrc *MessageReadCreate) Save(ctx context.Context) (*MessageRead, error) {
	var (
		err  error
		node *MessageRead
	)
	if len(mrc.hooks) == 0 {
		if err = mrc.check(); err != nil {
			return nil, err
		}
		node, err = mrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mrc.check(); err != nil {
				return nil, err
			}
			mrc.mutation = mutation
			if node, err = mrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mrc.hooks) - 1; i >= 0; i-- {
			if mrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mrc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageRead)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageReadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MessageReadCreate) SaveX(ctx context.Context) *MessageRead {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MessageReadCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MessageReadCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MessageReadCreate) check() error {
	if _, ok := mrc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "MessageRead.member_id"`)}
	}
	if _, ok := mrc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "MessageRead.group_id"`)}
	}
	if _, ok := mrc.mutation.LastID(); !ok {
		return &ValidationError{Name: "last_id", err: errors.New(`ent: missing required field "MessageRead.last_id"`)}
	}
	if _, ok := mrc.mutation.LastTime(); !ok {
		return &ValidationError{Name: "last_time", err: errors.New(`ent: missing required field "MessageRead.last_time"`)}
	}
	if v, ok := mrc.mutation.ID(); ok {
		if err := messageread.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "MessageRead.id": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member", err: errors.New(`ent: missing required edge "MessageRead.member"`)}
	}
	if _, ok := mrc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "MessageRead.group"`)}
	}
	return nil
}

func (mrc *MessageReadCreate) sqlSave(ctx context.Context) (*MessageRead, error) {
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected MessageRead.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (mrc *MessageReadCreate) createSpec() (*MessageRead, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageRead{config: mrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messageread.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: messageread.FieldID,
			},
		}
	)
	_spec.OnConflict = mrc.conflict
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mrc.mutation.LastID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageread.FieldLastID,
		})
		_node.LastID = value
	}
	if value, ok := mrc.mutation.LastTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messageread.FieldLastTime,
		})
		_node.LastTime = value
	}
	if nodes := mrc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.MemberTable,
			Columns: []string{messageread.MemberColumn},
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
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mrc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   messageread.GroupTable,
			Columns: []string{messageread.GroupColumn},
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
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MessageRead.Create().
//		SetMemberID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageReadUpsert) {
//			SetMemberID(v+v).
//		}).
//		Exec(ctx)
func (mrc *MessageReadCreate) OnConflict(opts ...sql.ConflictOption) *MessageReadUpsertOne {
	mrc.conflict = opts
	return &MessageReadUpsertOne{
		create: mrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mrc *MessageReadCreate) OnConflictColumns(columns ...string) *MessageReadUpsertOne {
	mrc.conflict = append(mrc.conflict, sql.ConflictColumns(columns...))
	return &MessageReadUpsertOne{
		create: mrc,
	}
}

type (
	// MessageReadUpsertOne is the builder for "upsert"-ing
	//  one MessageRead node.
	MessageReadUpsertOne struct {
		create *MessageReadCreate
	}

	// MessageReadUpsert is the "OnConflict" setter.
	MessageReadUpsert struct {
		*sql.UpdateSet
	}
)

// SetMemberID sets the "member_id" field.
func (u *MessageReadUpsert) SetMemberID(v string) *MessageReadUpsert {
	u.Set(messageread.FieldMemberID, v)
	return u
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageReadUpsert) UpdateMemberID() *MessageReadUpsert {
	u.SetExcluded(messageread.FieldMemberID)
	return u
}

// SetGroupID sets the "group_id" field.
func (u *MessageReadUpsert) SetGroupID(v string) *MessageReadUpsert {
	u.Set(messageread.FieldGroupID, v)
	return u
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageReadUpsert) UpdateGroupID() *MessageReadUpsert {
	u.SetExcluded(messageread.FieldGroupID)
	return u
}

// SetLastID sets the "last_id" field.
func (u *MessageReadUpsert) SetLastID(v string) *MessageReadUpsert {
	u.Set(messageread.FieldLastID, v)
	return u
}

// UpdateLastID sets the "last_id" field to the value that was provided on create.
func (u *MessageReadUpsert) UpdateLastID() *MessageReadUpsert {
	u.SetExcluded(messageread.FieldLastID)
	return u
}

// SetLastTime sets the "last_time" field.
func (u *MessageReadUpsert) SetLastTime(v time.Time) *MessageReadUpsert {
	u.Set(messageread.FieldLastTime, v)
	return u
}

// UpdateLastTime sets the "last_time" field to the value that was provided on create.
func (u *MessageReadUpsert) UpdateLastTime() *MessageReadUpsert {
	u.SetExcluded(messageread.FieldLastTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(messageread.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageReadUpsertOne) UpdateNewValues() *MessageReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(messageread.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageReadUpsertOne) Ignore() *MessageReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageReadUpsertOne) DoNothing() *MessageReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageReadCreate.OnConflict
// documentation for more info.
func (u *MessageReadUpsertOne) Update(set func(*MessageReadUpsert)) *MessageReadUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageReadUpsert{UpdateSet: update})
	}))
	return u
}

// SetMemberID sets the "member_id" field.
func (u *MessageReadUpsertOne) SetMemberID(v string) *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageReadUpsertOne) UpdateMemberID() *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateMemberID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *MessageReadUpsertOne) SetGroupID(v string) *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageReadUpsertOne) UpdateGroupID() *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateGroupID()
	})
}

// SetLastID sets the "last_id" field.
func (u *MessageReadUpsertOne) SetLastID(v string) *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetLastID(v)
	})
}

// UpdateLastID sets the "last_id" field to the value that was provided on create.
func (u *MessageReadUpsertOne) UpdateLastID() *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateLastID()
	})
}

// SetLastTime sets the "last_time" field.
func (u *MessageReadUpsertOne) SetLastTime(v time.Time) *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetLastTime(v)
	})
}

// UpdateLastTime sets the "last_time" field to the value that was provided on create.
func (u *MessageReadUpsertOne) UpdateLastTime() *MessageReadUpsertOne {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateLastTime()
	})
}

// Exec executes the query.
func (u *MessageReadUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageReadCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageReadUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageReadUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MessageReadUpsertOne.ID is not supported by MySQL driver. Use MessageReadUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageReadUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageReadCreateBulk is the builder for creating many MessageRead entities in bulk.
type MessageReadCreateBulk struct {
	config
	builders []*MessageReadCreate
	conflict []sql.ConflictOption
}

// Save creates the MessageRead entities in the database.
func (mrcb *MessageReadCreateBulk) Save(ctx context.Context) ([]*MessageRead, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MessageRead, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageReadMutation)
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
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MessageReadCreateBulk) SaveX(ctx context.Context) []*MessageRead {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MessageReadCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MessageReadCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MessageRead.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageReadUpsert) {
//			SetMemberID(v+v).
//		}).
//		Exec(ctx)
func (mrcb *MessageReadCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageReadUpsertBulk {
	mrcb.conflict = opts
	return &MessageReadUpsertBulk{
		create: mrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mrcb *MessageReadCreateBulk) OnConflictColumns(columns ...string) *MessageReadUpsertBulk {
	mrcb.conflict = append(mrcb.conflict, sql.ConflictColumns(columns...))
	return &MessageReadUpsertBulk{
		create: mrcb,
	}
}

// MessageReadUpsertBulk is the builder for "upsert"-ing
// a bulk of MessageRead nodes.
type MessageReadUpsertBulk struct {
	create *MessageReadCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(messageread.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageReadUpsertBulk) UpdateNewValues() *MessageReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(messageread.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MessageRead.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageReadUpsertBulk) Ignore() *MessageReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageReadUpsertBulk) DoNothing() *MessageReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageReadCreateBulk.OnConflict
// documentation for more info.
func (u *MessageReadUpsertBulk) Update(set func(*MessageReadUpsert)) *MessageReadUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageReadUpsert{UpdateSet: update})
	}))
	return u
}

// SetMemberID sets the "member_id" field.
func (u *MessageReadUpsertBulk) SetMemberID(v string) *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageReadUpsertBulk) UpdateMemberID() *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateMemberID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *MessageReadUpsertBulk) SetGroupID(v string) *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageReadUpsertBulk) UpdateGroupID() *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateGroupID()
	})
}

// SetLastID sets the "last_id" field.
func (u *MessageReadUpsertBulk) SetLastID(v string) *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetLastID(v)
	})
}

// UpdateLastID sets the "last_id" field to the value that was provided on create.
func (u *MessageReadUpsertBulk) UpdateLastID() *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateLastID()
	})
}

// SetLastTime sets the "last_time" field.
func (u *MessageReadUpsertBulk) SetLastTime(v time.Time) *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.SetLastTime(v)
	})
}

// UpdateLastTime sets the "last_time" field to the value that was provided on create.
func (u *MessageReadUpsertBulk) UpdateLastTime() *MessageReadUpsertBulk {
	return u.Update(func(s *MessageReadUpsert) {
		s.UpdateLastTime()
	})
}

// Exec executes the query.
func (u *MessageReadUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageReadCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageReadCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageReadUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}