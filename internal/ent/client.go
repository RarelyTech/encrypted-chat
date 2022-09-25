// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/chatpuppy/puppychat/internal/ent/migrate"

	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/messageread"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// GroupMember is the client for interacting with the GroupMember builders.
	GroupMember *GroupMemberClient
	// Key is the client for interacting with the Key builders.
	Key *KeyClient
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// MessageRead is the client for interacting with the MessageRead builders.
	MessageRead *MessageReadClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Group = NewGroupClient(c.config)
	c.GroupMember = NewGroupMemberClient(c.config)
	c.Key = NewKeyClient(c.config)
	c.Member = NewMemberClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.MessageRead = NewMessageReadClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Group:       NewGroupClient(cfg),
		GroupMember: NewGroupMemberClient(cfg),
		Key:         NewKeyClient(cfg),
		Member:      NewMemberClient(cfg),
		Message:     NewMessageClient(cfg),
		MessageRead: NewMessageReadClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Group:       NewGroupClient(cfg),
		GroupMember: NewGroupMemberClient(cfg),
		Key:         NewKeyClient(cfg),
		Member:      NewMemberClient(cfg),
		Message:     NewMessageClient(cfg),
		MessageRead: NewMessageReadClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Group.Use(hooks...)
	c.GroupMember.Use(hooks...)
	c.Key.Use(hooks...)
	c.Member.Use(hooks...)
	c.Message.Use(hooks...)
	c.MessageRead.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id string) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id string) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id string) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id string) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Group.
func (c *GroupClient) QueryOwner(gr *Group) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, group.OwnerTable, group.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a Group.
func (c *GroupClient) QueryMessages(gr *Group) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.MessagesTable, group.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMembers queries the members edge of a Group.
func (c *GroupClient) QueryMembers(gr *Group) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.MembersTable, group.MembersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroupMembers queries the group_members edge of a Group.
func (c *GroupClient) QueryGroupMembers(gr *Group) *GroupMemberQuery {
	query := &GroupMemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(groupmember.Table, groupmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, group.GroupMembersTable, group.GroupMembersColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	hooks := c.hooks.Group
	return append(hooks[:len(hooks):len(hooks)], group.Hooks[:]...)
}

// GroupMemberClient is a client for the GroupMember schema.
type GroupMemberClient struct {
	config
}

// NewGroupMemberClient returns a client for the GroupMember from the given config.
func NewGroupMemberClient(c config) *GroupMemberClient {
	return &GroupMemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `groupmember.Hooks(f(g(h())))`.
func (c *GroupMemberClient) Use(hooks ...Hook) {
	c.hooks.GroupMember = append(c.hooks.GroupMember, hooks...)
}

// Create returns a builder for creating a GroupMember entity.
func (c *GroupMemberClient) Create() *GroupMemberCreate {
	mutation := newGroupMemberMutation(c.config, OpCreate)
	return &GroupMemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GroupMember entities.
func (c *GroupMemberClient) CreateBulk(builders ...*GroupMemberCreate) *GroupMemberCreateBulk {
	return &GroupMemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GroupMember.
func (c *GroupMemberClient) Update() *GroupMemberUpdate {
	mutation := newGroupMemberMutation(c.config, OpUpdate)
	return &GroupMemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupMemberClient) UpdateOne(gm *GroupMember) *GroupMemberUpdateOne {
	mutation := newGroupMemberMutation(c.config, OpUpdateOne, withGroupMember(gm))
	return &GroupMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupMemberClient) UpdateOneID(id string) *GroupMemberUpdateOne {
	mutation := newGroupMemberMutation(c.config, OpUpdateOne, withGroupMemberID(id))
	return &GroupMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GroupMember.
func (c *GroupMemberClient) Delete() *GroupMemberDelete {
	mutation := newGroupMemberMutation(c.config, OpDelete)
	return &GroupMemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupMemberClient) DeleteOne(gm *GroupMember) *GroupMemberDeleteOne {
	return c.DeleteOneID(gm.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupMemberClient) DeleteOneID(id string) *GroupMemberDeleteOne {
	builder := c.Delete().Where(groupmember.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupMemberDeleteOne{builder}
}

// Query returns a query builder for GroupMember.
func (c *GroupMemberClient) Query() *GroupMemberQuery {
	return &GroupMemberQuery{
		config: c.config,
	}
}

// Get returns a GroupMember entity by its id.
func (c *GroupMemberClient) Get(ctx context.Context, id string) (*GroupMember, error) {
	return c.Query().Where(groupmember.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupMemberClient) GetX(ctx context.Context, id string) *GroupMember {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMember queries the member edge of a GroupMember.
func (c *GroupMemberClient) QueryMember(gm *GroupMember) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmember.Table, groupmember.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, groupmember.MemberTable, groupmember.MemberColumn),
		)
		fromV = sqlgraph.Neighbors(gm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a GroupMember.
func (c *GroupMemberClient) QueryGroup(gm *GroupMember) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmember.Table, groupmember.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, groupmember.GroupTable, groupmember.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(gm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryInviter queries the inviter edge of a GroupMember.
func (c *GroupMemberClient) QueryInviter(gm *GroupMember) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gm.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(groupmember.Table, groupmember.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, groupmember.InviterTable, groupmember.InviterColumn),
		)
		fromV = sqlgraph.Neighbors(gm.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupMemberClient) Hooks() []Hook {
	hooks := c.hooks.GroupMember
	return append(hooks[:len(hooks):len(hooks)], groupmember.Hooks[:]...)
}

// KeyClient is a client for the Key schema.
type KeyClient struct {
	config
}

// NewKeyClient returns a client for the Key from the given config.
func NewKeyClient(c config) *KeyClient {
	return &KeyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `key.Hooks(f(g(h())))`.
func (c *KeyClient) Use(hooks ...Hook) {
	c.hooks.Key = append(c.hooks.Key, hooks...)
}

// Create returns a builder for creating a Key entity.
func (c *KeyClient) Create() *KeyCreate {
	mutation := newKeyMutation(c.config, OpCreate)
	return &KeyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Key entities.
func (c *KeyClient) CreateBulk(builders ...*KeyCreate) *KeyCreateBulk {
	return &KeyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Key.
func (c *KeyClient) Update() *KeyUpdate {
	mutation := newKeyMutation(c.config, OpUpdate)
	return &KeyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KeyClient) UpdateOne(k *Key) *KeyUpdateOne {
	mutation := newKeyMutation(c.config, OpUpdateOne, withKey(k))
	return &KeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KeyClient) UpdateOneID(id string) *KeyUpdateOne {
	mutation := newKeyMutation(c.config, OpUpdateOne, withKeyID(id))
	return &KeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Key.
func (c *KeyClient) Delete() *KeyDelete {
	mutation := newKeyMutation(c.config, OpDelete)
	return &KeyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *KeyClient) DeleteOne(k *Key) *KeyDeleteOne {
	return c.DeleteOneID(k.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *KeyClient) DeleteOneID(id string) *KeyDeleteOne {
	builder := c.Delete().Where(key.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KeyDeleteOne{builder}
}

// Query returns a query builder for Key.
func (c *KeyClient) Query() *KeyQuery {
	return &KeyQuery{
		config: c.config,
	}
}

// Get returns a Key entity by its id.
func (c *KeyClient) Get(ctx context.Context, id string) (*Key, error) {
	return c.Query().Where(key.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KeyClient) GetX(ctx context.Context, id string) *Key {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMember queries the member edge of a Key.
func (c *KeyClient) QueryMember(k *Key) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := k.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(key.Table, key.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, key.MemberTable, key.MemberColumn),
		)
		fromV = sqlgraph.Neighbors(k.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a Key.
func (c *KeyClient) QueryGroup(k *Key) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := k.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(key.Table, key.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, key.GroupTable, key.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(k.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *KeyClient) Hooks() []Hook {
	hooks := c.hooks.Key
	return append(hooks[:len(hooks):len(hooks)], key.Hooks[:]...)
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Create returns a builder for creating a Member entity.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id string) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MemberClient) DeleteOneID(id string) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{
		config: c.config,
	}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id string) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id string) *Member {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwnGroups queries the own_groups edge of a Member.
func (c *MemberClient) QueryOwnGroups(m *Member) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.OwnGroupsTable, member.OwnGroupsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a Member.
func (c *MemberClient) QueryMessages(m *Member) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MessagesTable, member.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a Member.
func (c *MemberClient) QueryGroups(m *Member) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, member.GroupsTable, member.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroupMembers queries the group_members edge of a Member.
func (c *MemberClient) QueryGroupMembers(m *Member) *GroupMemberQuery {
	query := &GroupMemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(groupmember.Table, groupmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, member.GroupMembersTable, member.GroupMembersColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	hooks := c.hooks.Member
	return append(hooks[:len(hooks):len(hooks)], member.Hooks[:]...)
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a builder for creating a Message entity.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id string) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MessageClient) DeleteOneID(id string) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id string) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id string) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMember queries the member edge of a Message.
func (c *MessageClient) QueryMember(m *Message) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.MemberTable, message.MemberColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a Message.
func (c *MessageClient) QueryGroup(m *Message) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.GroupTable, message.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Message.
func (c *MessageClient) QueryParent(m *Message) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.ParentTable, message.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Message.
func (c *MessageClient) QueryChildren(m *Message) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, message.ChildrenTable, message.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	hooks := c.hooks.Message
	return append(hooks[:len(hooks):len(hooks)], message.Hooks[:]...)
}

// MessageReadClient is a client for the MessageRead schema.
type MessageReadClient struct {
	config
}

// NewMessageReadClient returns a client for the MessageRead from the given config.
func NewMessageReadClient(c config) *MessageReadClient {
	return &MessageReadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `messageread.Hooks(f(g(h())))`.
func (c *MessageReadClient) Use(hooks ...Hook) {
	c.hooks.MessageRead = append(c.hooks.MessageRead, hooks...)
}

// Create returns a builder for creating a MessageRead entity.
func (c *MessageReadClient) Create() *MessageReadCreate {
	mutation := newMessageReadMutation(c.config, OpCreate)
	return &MessageReadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MessageRead entities.
func (c *MessageReadClient) CreateBulk(builders ...*MessageReadCreate) *MessageReadCreateBulk {
	return &MessageReadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MessageRead.
func (c *MessageReadClient) Update() *MessageReadUpdate {
	mutation := newMessageReadMutation(c.config, OpUpdate)
	return &MessageReadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageReadClient) UpdateOne(mr *MessageRead) *MessageReadUpdateOne {
	mutation := newMessageReadMutation(c.config, OpUpdateOne, withMessageRead(mr))
	return &MessageReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageReadClient) UpdateOneID(id string) *MessageReadUpdateOne {
	mutation := newMessageReadMutation(c.config, OpUpdateOne, withMessageReadID(id))
	return &MessageReadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MessageRead.
func (c *MessageReadClient) Delete() *MessageReadDelete {
	mutation := newMessageReadMutation(c.config, OpDelete)
	return &MessageReadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageReadClient) DeleteOne(mr *MessageRead) *MessageReadDeleteOne {
	return c.DeleteOneID(mr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MessageReadClient) DeleteOneID(id string) *MessageReadDeleteOne {
	builder := c.Delete().Where(messageread.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageReadDeleteOne{builder}
}

// Query returns a query builder for MessageRead.
func (c *MessageReadClient) Query() *MessageReadQuery {
	return &MessageReadQuery{
		config: c.config,
	}
}

// Get returns a MessageRead entity by its id.
func (c *MessageReadClient) Get(ctx context.Context, id string) (*MessageRead, error) {
	return c.Query().Where(messageread.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageReadClient) GetX(ctx context.Context, id string) *MessageRead {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMember queries the member edge of a MessageRead.
func (c *MessageReadClient) QueryMember(mr *MessageRead) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(messageread.Table, messageread.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, messageread.MemberTable, messageread.MemberColumn),
		)
		fromV = sqlgraph.Neighbors(mr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroup queries the group edge of a MessageRead.
func (c *MessageReadClient) QueryGroup(mr *MessageRead) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := mr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(messageread.Table, messageread.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, messageread.GroupTable, messageread.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(mr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageReadClient) Hooks() []Hook {
	hooks := c.hooks.MessageRead
	return append(hooks[:len(hooks):len(hooks)], messageread.Hooks[:]...)
}
