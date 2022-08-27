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
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessageCreate) SetCreatedAt(t time.Time) *MessageCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreatedAt(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetKeyID sets the "key_id" field.
func (mc *MessageCreate) SetKeyID(u uint64) *MessageCreate {
	mc.mutation.SetKeyID(u)
	return mc
}

// SetGroupID sets the "group_id" field.
func (mc *MessageCreate) SetGroupID(u uint64) *MessageCreate {
	mc.mutation.SetGroupID(u)
	return mc
}

// SetMemberID sets the "member_id" field.
func (mc *MessageCreate) SetMemberID(u uint64) *MessageCreate {
	mc.mutation.SetMemberID(u)
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(u uint64) *MessageCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetKey sets the "key" edge to the Key entity.
func (mc *MessageCreate) SetKey(k *Key) *MessageCreate {
	return mc.SetKeyID(k.ID)
}

// SetOwnerID sets the "owner" edge to the Member entity by ID.
func (mc *MessageCreate) SetOwnerID(id uint64) *MessageCreate {
	mc.mutation.SetOwnerID(id)
	return mc
}

// SetOwner sets the "owner" edge to the Member entity.
func (mc *MessageCreate) SetOwner(m *Member) *MessageCreate {
	return mc.SetOwnerID(m.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (mc *MessageCreate) SetGroup(g *Group) *MessageCreate {
	return mc.SetGroupID(g.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
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
			mutation, ok := m.(*MessageMutation)
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
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := message.DefaultCreatedAt
		mc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Message.created_at"`)}
	}
	if _, ok := mc.mutation.KeyID(); !ok {
		return &ValidationError{Name: "key_id", err: errors.New(`ent: missing required field "Message.key_id"`)}
	}
	if _, ok := mc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "Message.group_id"`)}
	}
	if _, ok := mc.mutation.MemberID(); !ok {
		return &ValidationError{Name: "member_id", err: errors.New(`ent: missing required field "Message.member_id"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Message.content"`)}
	}
	if _, ok := mc.mutation.KeyID(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required edge "Message.key"`)}
	}
	if _, ok := mc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Message.owner"`)}
	}
	if _, ok := mc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Message.group"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: message.FieldID,
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
			Column: message.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldContent,
		})
		_node.Content = value
	}
	if nodes := mc.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   message.KeyTable,
			Columns: []string{message.KeyColumn},
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
	if nodes := mc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.OwnerTable,
			Columns: []string{message.OwnerColumn},
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
	if nodes := mc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsert) SetCreatedAt(v time.Time) *MessageUpsert {
	u.Set(message.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateCreatedAt() *MessageUpsert {
	u.SetExcluded(message.FieldCreatedAt)
	return u
}

// SetKeyID sets the "key_id" field.
func (u *MessageUpsert) SetKeyID(v uint64) *MessageUpsert {
	u.Set(message.FieldKeyID, v)
	return u
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateKeyID() *MessageUpsert {
	u.SetExcluded(message.FieldKeyID)
	return u
}

// SetGroupID sets the "group_id" field.
func (u *MessageUpsert) SetGroupID(v uint64) *MessageUpsert {
	u.Set(message.FieldGroupID, v)
	return u
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateGroupID() *MessageUpsert {
	u.SetExcluded(message.FieldGroupID)
	return u
}

// SetMemberID sets the "member_id" field.
func (u *MessageUpsert) SetMemberID(v uint64) *MessageUpsert {
	u.Set(message.FieldMemberID, v)
	return u
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMemberID() *MessageUpsert {
	u.SetExcluded(message.FieldMemberID)
	return u
}

// SetContent sets the "content" field.
func (u *MessageUpsert) SetContent(v string) *MessageUpsert {
	u.Set(message.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsert) UpdateContent() *MessageUpsert {
	u.SetExcluded(message.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(message.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertOne) SetCreatedAt(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateCreatedAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetKeyID sets the "key_id" field.
func (u *MessageUpsertOne) SetKeyID(v uint64) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetKeyID(v)
	})
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateKeyID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateKeyID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *MessageUpsertOne) SetGroupID(v uint64) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateGroupID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateGroupID()
	})
}

// SetMemberID sets the "member_id" field.
func (u *MessageUpsertOne) SetMemberID(v uint64) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMemberID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMemberID()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertOne) SetContent(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateContent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
				return
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(message.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *MessageUpsertBulk) SetCreatedAt(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateCreatedAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetKeyID sets the "key_id" field.
func (u *MessageUpsertBulk) SetKeyID(v uint64) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetKeyID(v)
	})
}

// UpdateKeyID sets the "key_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateKeyID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateKeyID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *MessageUpsertBulk) SetGroupID(v uint64) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateGroupID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateGroupID()
	})
}

// SetMemberID sets the "member_id" field.
func (u *MessageUpsertBulk) SetMemberID(v uint64) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMemberID(v)
	})
}

// UpdateMemberID sets the "member_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMemberID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMemberID()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertBulk) SetContent(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateContent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
