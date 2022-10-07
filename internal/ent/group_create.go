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

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetCategory sets the "category" field.
func (gc *GroupCreate) SetCategory(s string) *GroupCreate {
	gc.mutation.SetCategory(s)
	return gc
}

// SetOwnerID sets the "owner_id" field.
func (gc *GroupCreate) SetOwnerID(s string) *GroupCreate {
	gc.mutation.SetOwnerID(s)
	return gc
}

// SetMembersMax sets the "members_max" field.
func (gc *GroupCreate) SetMembersMax(i int) *GroupCreate {
	gc.mutation.SetMembersMax(i)
	return gc
}

// SetMembersCount sets the "members_count" field.
func (gc *GroupCreate) SetMembersCount(i int) *GroupCreate {
	gc.mutation.SetMembersCount(i)
	return gc
}

// SetNillableMembersCount sets the "members_count" field if the given value is not nil.
func (gc *GroupCreate) SetNillableMembersCount(i *int) *GroupCreate {
	if i != nil {
		gc.SetMembersCount(*i)
	}
	return gc
}

// SetPublic sets the "public" field.
func (gc *GroupCreate) SetPublic(b bool) *GroupCreate {
	gc.mutation.SetPublic(b)
	return gc
}

// SetAddress sets the "address" field.
func (gc *GroupCreate) SetAddress(s string) *GroupCreate {
	gc.mutation.SetAddress(s)
	return gc
}

// SetIntro sets the "intro" field.
func (gc *GroupCreate) SetIntro(s string) *GroupCreate {
	gc.mutation.SetIntro(s)
	return gc
}

// SetNillableIntro sets the "intro" field if the given value is not nil.
func (gc *GroupCreate) SetNillableIntro(s *string) *GroupCreate {
	if s != nil {
		gc.SetIntro(*s)
	}
	return gc
}

// SetKeys sets the "keys" field.
func (gc *GroupCreate) SetKeys(s string) *GroupCreate {
	gc.mutation.SetKeys(s)
	return gc
}

// SetLastNode sets the "last_node" field.
func (gc *GroupCreate) SetLastNode(i int64) *GroupCreate {
	gc.mutation.SetLastNode(i)
	return gc
}

// SetID sets the "id" field.
func (gc *GroupCreate) SetID(s string) *GroupCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetOwner sets the "owner" edge to the Member entity.
func (gc *GroupCreate) SetOwner(m *Member) *GroupCreate {
	return gc.SetOwnerID(m.ID)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (gc *GroupCreate) AddMessageIDs(ids ...string) *GroupCreate {
	gc.mutation.AddMessageIDs(ids...)
	return gc
}

// AddMessages adds the "messages" edges to the Message entity.
func (gc *GroupCreate) AddMessages(m ...*Message) *GroupCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMessageIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the Member entity by IDs.
func (gc *GroupCreate) AddMemberIDs(ids ...string) *GroupCreate {
	gc.mutation.AddMemberIDs(ids...)
	return gc
}

// AddMembers adds the "members" edges to the Member entity.
func (gc *GroupCreate) AddMembers(m ...*Member) *GroupCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gc.AddMemberIDs(ids...)
}

// AddGroupMemberIDs adds the "group_members" edge to the GroupMember entity by IDs.
func (gc *GroupCreate) AddGroupMemberIDs(ids ...string) *GroupCreate {
	gc.mutation.AddGroupMemberIDs(ids...)
	return gc
}

// AddGroupMembers adds the "group_members" edges to the GroupMember entity.
func (gc *GroupCreate) AddGroupMembers(g ...*GroupMember) *GroupCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gc.AddGroupMemberIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	if err := gc.defaults(); err != nil {
		return nil, err
	}
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		if group.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized group.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := group.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.MembersCount(); !ok {
		v := group.DefaultMembersCount
		gc.mutation.SetMembersCount(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Group.created_at"`)}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if _, ok := gc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Group.category"`)}
	}
	if _, ok := gc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "Group.owner_id"`)}
	}
	if _, ok := gc.mutation.MembersMax(); !ok {
		return &ValidationError{Name: "members_max", err: errors.New(`ent: missing required field "Group.members_max"`)}
	}
	if _, ok := gc.mutation.MembersCount(); !ok {
		return &ValidationError{Name: "members_count", err: errors.New(`ent: missing required field "Group.members_count"`)}
	}
	if _, ok := gc.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "Group.public"`)}
	}
	if _, ok := gc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Group.address"`)}
	}
	if _, ok := gc.mutation.Keys(); !ok {
		return &ValidationError{Name: "keys", err: errors.New(`ent: missing required field "Group.keys"`)}
	}
	if _, ok := gc.mutation.LastNode(); !ok {
		return &ValidationError{Name: "last_node", err: errors.New(`ent: missing required field "Group.last_node"`)}
	}
	if v, ok := gc.mutation.ID(); ok {
		if err := group.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Group.id": %w`, err)}
		}
	}
	if _, ok := gc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Group.owner"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Group.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: group.FieldID,
			},
		}
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gc.mutation.Category(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldCategory,
		})
		_node.Category = value
	}
	if value, ok := gc.mutation.MembersMax(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersMax,
		})
		_node.MembersMax = value
	}
	if value, ok := gc.mutation.MembersCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMembersCount,
		})
		_node.MembersCount = value
	}
	if value, ok := gc.mutation.Public(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldPublic,
		})
		_node.Public = value
	}
	if value, ok := gc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := gc.mutation.Intro(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldIntro,
		})
		_node.Intro = value
	}
	if value, ok := gc.mutation.Keys(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldKeys,
		})
		_node.Keys = value
	}
	if value, ok := gc.mutation.LastNode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: group.FieldLastNode,
		})
		_node.LastNode = value
	}
	if nodes := gc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
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
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
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
	if nodes := gc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
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
		createE := &GroupMemberCreate{config: gc.config, mutation: newGroupMemberMutation(gc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.GroupMembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupMembersTable,
			Columns: []string{group.GroupMembersColumn},
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
//	client.Group.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gc *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	gc.conflict = opts
	return &GroupUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: gc,
	}
}

type (
	// GroupUpsertOne is the builder for "upsert"-ing
	//  one Group node.
	GroupUpsertOne struct {
		create *GroupCreate
	}

	// GroupUpsert is the "OnConflict" setter.
	GroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *GroupUpsert) SetName(v string) *GroupUpsert {
	u.Set(group.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsert) UpdateName() *GroupUpsert {
	u.SetExcluded(group.FieldName)
	return u
}

// SetCategory sets the "category" field.
func (u *GroupUpsert) SetCategory(v string) *GroupUpsert {
	u.Set(group.FieldCategory, v)
	return u
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *GroupUpsert) UpdateCategory() *GroupUpsert {
	u.SetExcluded(group.FieldCategory)
	return u
}

// SetOwnerID sets the "owner_id" field.
func (u *GroupUpsert) SetOwnerID(v string) *GroupUpsert {
	u.Set(group.FieldOwnerID, v)
	return u
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *GroupUpsert) UpdateOwnerID() *GroupUpsert {
	u.SetExcluded(group.FieldOwnerID)
	return u
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsert) SetMembersMax(v int) *GroupUpsert {
	u.Set(group.FieldMembersMax, v)
	return u
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMembersMax() *GroupUpsert {
	u.SetExcluded(group.FieldMembersMax)
	return u
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsert) AddMembersMax(v int) *GroupUpsert {
	u.Add(group.FieldMembersMax, v)
	return u
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsert) SetMembersCount(v int) *GroupUpsert {
	u.Set(group.FieldMembersCount, v)
	return u
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMembersCount() *GroupUpsert {
	u.SetExcluded(group.FieldMembersCount)
	return u
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsert) AddMembersCount(v int) *GroupUpsert {
	u.Add(group.FieldMembersCount, v)
	return u
}

// SetPublic sets the "public" field.
func (u *GroupUpsert) SetPublic(v bool) *GroupUpsert {
	u.Set(group.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsert) UpdatePublic() *GroupUpsert {
	u.SetExcluded(group.FieldPublic)
	return u
}

// SetIntro sets the "intro" field.
func (u *GroupUpsert) SetIntro(v string) *GroupUpsert {
	u.Set(group.FieldIntro, v)
	return u
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsert) UpdateIntro() *GroupUpsert {
	u.SetExcluded(group.FieldIntro)
	return u
}

// ClearIntro clears the value of the "intro" field.
func (u *GroupUpsert) ClearIntro() *GroupUpsert {
	u.SetNull(group.FieldIntro)
	return u
}

// SetKeys sets the "keys" field.
func (u *GroupUpsert) SetKeys(v string) *GroupUpsert {
	u.Set(group.FieldKeys, v)
	return u
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *GroupUpsert) UpdateKeys() *GroupUpsert {
	u.SetExcluded(group.FieldKeys)
	return u
}

// SetLastNode sets the "last_node" field.
func (u *GroupUpsert) SetLastNode(v int64) *GroupUpsert {
	u.Set(group.FieldLastNode, v)
	return u
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *GroupUpsert) UpdateLastNode() *GroupUpsert {
	u.SetExcluded(group.FieldLastNode)
	return u
}

// AddLastNode adds v to the "last_node" field.
func (u *GroupUpsert) AddLastNode(v int64) *GroupUpsert {
	u.Add(group.FieldLastNode, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(group.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(group.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(group.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Address(); exists {
			s.SetIgnore(group.FieldAddress)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupUpsertOne) Ignore() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertOne) DoNothing() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreate.OnConflict
// documentation for more info.
func (u *GroupUpsertOne) Update(set func(*GroupUpsert)) *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertOne) SetName(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateName() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetCategory sets the "category" field.
func (u *GroupUpsertOne) SetCategory(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateCategory() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCategory()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *GroupUpsertOne) SetOwnerID(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateOwnerID() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateOwnerID()
	})
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsertOne) SetMembersMax(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersMax(v)
	})
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsertOne) AddMembersMax(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersMax(v)
	})
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMembersMax() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersMax()
	})
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsertOne) SetMembersCount(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersCount(v)
	})
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsertOne) AddMembersCount(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersCount(v)
	})
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMembersCount() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersCount()
	})
}

// SetPublic sets the "public" field.
func (u *GroupUpsertOne) SetPublic(v bool) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdatePublic() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePublic()
	})
}

// SetIntro sets the "intro" field.
func (u *GroupUpsertOne) SetIntro(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateIntro() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateIntro()
	})
}

// ClearIntro clears the value of the "intro" field.
func (u *GroupUpsertOne) ClearIntro() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearIntro()
	})
}

// SetKeys sets the "keys" field.
func (u *GroupUpsertOne) SetKeys(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetKeys(v)
	})
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateKeys() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateKeys()
	})
}

// SetLastNode sets the "last_node" field.
func (u *GroupUpsertOne) SetLastNode(v int64) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetLastNode(v)
	})
}

// AddLastNode adds v to the "last_node" field.
func (u *GroupUpsertOne) AddLastNode(v int64) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddLastNode(v)
	})
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateLastNode() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateLastNode()
	})
}

// Exec executes the query.
func (u *GroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GroupUpsertOne.ID is not supported by MySQL driver. Use GroupUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	gcb.conflict = opts
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// GroupUpsertBulk is the builder for "upsert"-ing
// a bulk of Group nodes.
type GroupUpsertBulk struct {
	create *GroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(group.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupUpsertBulk) UpdateNewValues() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(group.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(group.FieldCreatedAt)
			}
			if _, exists := b.mutation.Address(); exists {
				s.SetIgnore(group.FieldAddress)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupUpsertBulk) Ignore() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertBulk) DoNothing() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreateBulk.OnConflict
// documentation for more info.
func (u *GroupUpsertBulk) Update(set func(*GroupUpsert)) *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertBulk) SetName(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateName() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetCategory sets the "category" field.
func (u *GroupUpsertBulk) SetCategory(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateCategory() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCategory()
	})
}

// SetOwnerID sets the "owner_id" field.
func (u *GroupUpsertBulk) SetOwnerID(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetOwnerID(v)
	})
}

// UpdateOwnerID sets the "owner_id" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateOwnerID() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateOwnerID()
	})
}

// SetMembersMax sets the "members_max" field.
func (u *GroupUpsertBulk) SetMembersMax(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersMax(v)
	})
}

// AddMembersMax adds v to the "members_max" field.
func (u *GroupUpsertBulk) AddMembersMax(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersMax(v)
	})
}

// UpdateMembersMax sets the "members_max" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMembersMax() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersMax()
	})
}

// SetMembersCount sets the "members_count" field.
func (u *GroupUpsertBulk) SetMembersCount(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMembersCount(v)
	})
}

// AddMembersCount adds v to the "members_count" field.
func (u *GroupUpsertBulk) AddMembersCount(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMembersCount(v)
	})
}

// UpdateMembersCount sets the "members_count" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMembersCount() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMembersCount()
	})
}

// SetPublic sets the "public" field.
func (u *GroupUpsertBulk) SetPublic(v bool) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdatePublic() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdatePublic()
	})
}

// SetIntro sets the "intro" field.
func (u *GroupUpsertBulk) SetIntro(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetIntro(v)
	})
}

// UpdateIntro sets the "intro" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateIntro() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateIntro()
	})
}

// ClearIntro clears the value of the "intro" field.
func (u *GroupUpsertBulk) ClearIntro() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearIntro()
	})
}

// SetKeys sets the "keys" field.
func (u *GroupUpsertBulk) SetKeys(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetKeys(v)
	})
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateKeys() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateKeys()
	})
}

// SetLastNode sets the "last_node" field.
func (u *GroupUpsertBulk) SetLastNode(v int64) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetLastNode(v)
	})
}

// AddLastNode adds v to the "last_node" field.
func (u *GroupUpsertBulk) AddLastNode(v int64) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddLastNode(v)
	})
}

// UpdateLastNode sets the "last_node" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateLastNode() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateLastNode()
	})
}

// Exec executes the query.
func (u *GroupUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
