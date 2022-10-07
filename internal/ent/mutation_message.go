// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/message"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// MessageMutation represents an operation that mutates the Message nodes in the graph.
type MessageMutation struct {
	config
	op              Op
	typ             string
	id              *string
	created_at      *time.Time
	content         *[]byte
	owner           **model.Member
	last_node       *int64
	addlast_node    *int64
	clearedFields   map[string]struct{}
	member          *string
	clearedmember   bool
	group           *string
	clearedgroup    bool
	parent          *string
	clearedparent   bool
	children        map[string]struct{}
	removedchildren map[string]struct{}
	clearedchildren bool
	done            bool
	oldValue        func(context.Context) (*Message, error)
	predicates      []predicate.Message
}

var _ ent.Mutation = (*MessageMutation)(nil)

// messageOption allows management of the mutation configuration using functional options.
type messageOption func(*MessageMutation)

// newMessageMutation creates new mutation for the Message entity.
func newMessageMutation(c config, op Op, opts ...messageOption) *MessageMutation {
	m := &MessageMutation{
		config:        c,
		op:            op,
		typ:           TypeMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageID sets the ID field of the mutation.
func withMessageID(id string) messageOption {
	return func(m *MessageMutation) {
		var (
			err   error
			once  sync.Once
			value *Message
		)
		m.oldValue = func(ctx context.Context) (*Message, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Message.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessage sets the old Message of the mutation.
func withMessage(node *Message) messageOption {
	return func(m *MessageMutation) {
		m.oldValue = func(context.Context) (*Message, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Message entities.
func (m *MessageMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessageMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessageMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Message.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *MessageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MessageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *MessageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetGroupID sets the "group_id" field.
func (m *MessageMutation) SetGroupID(s string) {
	m.group = &s
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *MessageMutation) GroupID() (r string, exists bool) {
	v := m.group
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldGroupID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGroupID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGroupID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupID: %w", err)
	}
	return oldValue.GroupID, nil
}

// ResetGroupID resets all changes to the "group_id" field.
func (m *MessageMutation) ResetGroupID() {
	m.group = nil
}

// SetMemberID sets the "member_id" field.
func (m *MessageMutation) SetMemberID(s string) {
	m.member = &s
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *MessageMutation) MemberID() (r string, exists bool) {
	v := m.member
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldMemberID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMemberID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMemberID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMemberID: %w", err)
	}
	return oldValue.MemberID, nil
}

// ResetMemberID resets all changes to the "member_id" field.
func (m *MessageMutation) ResetMemberID() {
	m.member = nil
}

// SetContent sets the "content" field.
func (m *MessageMutation) SetContent(b []byte) {
	m.content = &b
}

// Content returns the value of the "content" field in the mutation.
func (m *MessageMutation) Content() (r []byte, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldContent(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *MessageMutation) ResetContent() {
	m.content = nil
}

// SetParentID sets the "parent_id" field.
func (m *MessageMutation) SetParentID(s string) {
	m.parent = &s
}

// ParentID returns the value of the "parent_id" field in the mutation.
func (m *MessageMutation) ParentID() (r string, exists bool) {
	v := m.parent
	if v == nil {
		return
	}
	return *v, true
}

// OldParentID returns the old "parent_id" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldParentID(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParentID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParentID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentID: %w", err)
	}
	return oldValue.ParentID, nil
}

// ClearParentID clears the value of the "parent_id" field.
func (m *MessageMutation) ClearParentID() {
	m.parent = nil
	m.clearedFields[message.FieldParentID] = struct{}{}
}

// ParentIDCleared returns if the "parent_id" field was cleared in this mutation.
func (m *MessageMutation) ParentIDCleared() bool {
	_, ok := m.clearedFields[message.FieldParentID]
	return ok
}

// ResetParentID resets all changes to the "parent_id" field.
func (m *MessageMutation) ResetParentID() {
	m.parent = nil
	delete(m.clearedFields, message.FieldParentID)
}

// SetOwner sets the "owner" field.
func (m *MessageMutation) SetOwner(value *model.Member) {
	m.owner = &value
}

// Owner returns the value of the "owner" field in the mutation.
func (m *MessageMutation) Owner() (r *model.Member, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldOwner returns the old "owner" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldOwner(ctx context.Context) (v *model.Member, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwner is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwner requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwner: %w", err)
	}
	return oldValue.Owner, nil
}

// ResetOwner resets all changes to the "owner" field.
func (m *MessageMutation) ResetOwner() {
	m.owner = nil
}

// SetLastNode sets the "last_node" field.
func (m *MessageMutation) SetLastNode(i int64) {
	m.last_node = &i
	m.addlast_node = nil
}

// LastNode returns the value of the "last_node" field in the mutation.
func (m *MessageMutation) LastNode() (r int64, exists bool) {
	v := m.last_node
	if v == nil {
		return
	}
	return *v, true
}

// OldLastNode returns the old "last_node" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldLastNode(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastNode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastNode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastNode: %w", err)
	}
	return oldValue.LastNode, nil
}

// AddLastNode adds i to the "last_node" field.
func (m *MessageMutation) AddLastNode(i int64) {
	if m.addlast_node != nil {
		*m.addlast_node += i
	} else {
		m.addlast_node = &i
	}
}

// AddedLastNode returns the value that was added to the "last_node" field in this mutation.
func (m *MessageMutation) AddedLastNode() (r int64, exists bool) {
	v := m.addlast_node
	if v == nil {
		return
	}
	return *v, true
}

// ResetLastNode resets all changes to the "last_node" field.
func (m *MessageMutation) ResetLastNode() {
	m.last_node = nil
	m.addlast_node = nil
}

// ClearMember clears the "member" edge to the Member entity.
func (m *MessageMutation) ClearMember() {
	m.clearedmember = true
}

// MemberCleared reports if the "member" edge to the Member entity was cleared.
func (m *MessageMutation) MemberCleared() bool {
	return m.clearedmember
}

// MemberIDs returns the "member" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// MemberID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) MemberIDs() (ids []string) {
	if id := m.member; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetMember resets all changes to the "member" edge.
func (m *MessageMutation) ResetMember() {
	m.member = nil
	m.clearedmember = false
}

// ClearGroup clears the "group" edge to the Group entity.
func (m *MessageMutation) ClearGroup() {
	m.clearedgroup = true
}

// GroupCleared reports if the "group" edge to the Group entity was cleared.
func (m *MessageMutation) GroupCleared() bool {
	return m.clearedgroup
}

// GroupIDs returns the "group" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// GroupID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) GroupIDs() (ids []string) {
	if id := m.group; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetGroup resets all changes to the "group" edge.
func (m *MessageMutation) ResetGroup() {
	m.group = nil
	m.clearedgroup = false
}

// ClearParent clears the "parent" edge to the Message entity.
func (m *MessageMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the Message entity was cleared.
func (m *MessageMutation) ParentCleared() bool {
	return m.ParentIDCleared() || m.clearedparent
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *MessageMutation) ParentIDs() (ids []string) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *MessageMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the "children" edge to the Message entity by ids.
func (m *MessageMutation) AddChildIDs(ids ...string) {
	if m.children == nil {
		m.children = make(map[string]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the Message entity.
func (m *MessageMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Message entity was cleared.
func (m *MessageMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the Message entity by IDs.
func (m *MessageMutation) RemoveChildIDs(ids ...string) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the Message entity.
func (m *MessageMutation) RemovedChildrenIDs() (ids []string) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *MessageMutation) ChildrenIDs() (ids []string) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *MessageMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// Where appends a list predicates to the MessageMutation builder.
func (m *MessageMutation) Where(ps ...predicate.Message) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MessageMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Message).
func (m *MessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.created_at != nil {
		fields = append(fields, message.FieldCreatedAt)
	}
	if m.group != nil {
		fields = append(fields, message.FieldGroupID)
	}
	if m.member != nil {
		fields = append(fields, message.FieldMemberID)
	}
	if m.content != nil {
		fields = append(fields, message.FieldContent)
	}
	if m.parent != nil {
		fields = append(fields, message.FieldParentID)
	}
	if m.owner != nil {
		fields = append(fields, message.FieldOwner)
	}
	if m.last_node != nil {
		fields = append(fields, message.FieldLastNode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case message.FieldCreatedAt:
		return m.CreatedAt()
	case message.FieldGroupID:
		return m.GroupID()
	case message.FieldMemberID:
		return m.MemberID()
	case message.FieldContent:
		return m.Content()
	case message.FieldParentID:
		return m.ParentID()
	case message.FieldOwner:
		return m.Owner()
	case message.FieldLastNode:
		return m.LastNode()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case message.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case message.FieldGroupID:
		return m.OldGroupID(ctx)
	case message.FieldMemberID:
		return m.OldMemberID(ctx)
	case message.FieldContent:
		return m.OldContent(ctx)
	case message.FieldParentID:
		return m.OldParentID(ctx)
	case message.FieldOwner:
		return m.OldOwner(ctx)
	case message.FieldLastNode:
		return m.OldLastNode(ctx)
	}
	return nil, fmt.Errorf("unknown Message field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case message.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case message.FieldGroupID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case message.FieldMemberID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case message.FieldContent:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case message.FieldParentID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentID(v)
		return nil
	case message.FieldOwner:
		v, ok := value.(*model.Member)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwner(v)
		return nil
	case message.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageMutation) AddedFields() []string {
	var fields []string
	if m.addlast_node != nil {
		fields = append(fields, message.FieldLastNode)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case message.FieldLastNode:
		return m.AddedLastNode()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case message.FieldLastNode:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLastNode(v)
		return nil
	}
	return fmt.Errorf("unknown Message numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(message.FieldParentID) {
		fields = append(fields, message.FieldParentID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageMutation) ClearField(name string) error {
	switch name {
	case message.FieldParentID:
		m.ClearParentID()
		return nil
	}
	return fmt.Errorf("unknown Message nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageMutation) ResetField(name string) error {
	switch name {
	case message.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case message.FieldGroupID:
		m.ResetGroupID()
		return nil
	case message.FieldMemberID:
		m.ResetMemberID()
		return nil
	case message.FieldContent:
		m.ResetContent()
		return nil
	case message.FieldParentID:
		m.ResetParentID()
		return nil
	case message.FieldOwner:
		m.ResetOwner()
		return nil
	case message.FieldLastNode:
		m.ResetLastNode()
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 4)
	if m.member != nil {
		edges = append(edges, message.EdgeMember)
	}
	if m.group != nil {
		edges = append(edges, message.EdgeGroup)
	}
	if m.parent != nil {
		edges = append(edges, message.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, message.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case message.EdgeMember:
		if id := m.member; id != nil {
			return []ent.Value{*id}
		}
	case message.EdgeGroup:
		if id := m.group; id != nil {
			return []ent.Value{*id}
		}
	case message.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case message.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 4)
	if m.removedchildren != nil {
		edges = append(edges, message.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case message.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 4)
	if m.clearedmember {
		edges = append(edges, message.EdgeMember)
	}
	if m.clearedgroup {
		edges = append(edges, message.EdgeGroup)
	}
	if m.clearedparent {
		edges = append(edges, message.EdgeParent)
	}
	if m.clearedchildren {
		edges = append(edges, message.EdgeChildren)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageMutation) EdgeCleared(name string) bool {
	switch name {
	case message.EdgeMember:
		return m.clearedmember
	case message.EdgeGroup:
		return m.clearedgroup
	case message.EdgeParent:
		return m.clearedparent
	case message.EdgeChildren:
		return m.clearedchildren
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageMutation) ClearEdge(name string) error {
	switch name {
	case message.EdgeMember:
		m.ClearMember()
		return nil
	case message.EdgeGroup:
		m.ClearGroup()
		return nil
	case message.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Message unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageMutation) ResetEdge(name string) error {
	switch name {
	case message.EdgeMember:
		m.ResetMember()
		return nil
	case message.EdgeGroup:
		m.ResetGroup()
		return nil
	case message.EdgeParent:
		m.ResetParent()
		return nil
	case message.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown Message edge %s", name)
}

