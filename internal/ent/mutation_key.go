// Code generated, DO NOT EDIT.

package ent

import (
    "context"
    "fmt"
    "sync"
    "errors"
    "time"
    "github.com/chatpuppy/puppychat/internal/ent/key"
    "github.com/chatpuppy/puppychat/internal/ent/predicate"

    "entgo.io/ent"
)


// KeyMutation represents an operation that mutates the Key nodes in the graph.
type KeyMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	created_at    *time.Time
	group_id      *uint64
	addgroup_id   *int64
	member_id     *uint64
	addmember_id  *int64
	key           *string
	enable        *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Key, error)
	predicates    []predicate.Key
}

var _ ent.Mutation = (*KeyMutation)(nil)

// keyOption allows management of the mutation configuration using functional options.
type keyOption func(*KeyMutation)

// newKeyMutation creates new mutation for the Key entity.
func newKeyMutation(c config, op Op, opts ...keyOption) *KeyMutation {
	m := &KeyMutation{
		config:        c,
		op:            op,
		typ:           TypeKey,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withKeyID sets the ID field of the mutation.
func withKeyID(id uint64) keyOption {
	return func(m *KeyMutation) {
		var (
			err   error
			once  sync.Once
			value *Key
		)
		m.oldValue = func(ctx context.Context) (*Key, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Key.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withKey sets the old Key of the mutation.
func withKey(node *Key) keyOption {
	return func(m *KeyMutation) {
		m.oldValue = func(context.Context) (*Key, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m KeyMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m KeyMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Key entities.
func (m *KeyMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *KeyMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *KeyMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Key.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *KeyMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *KeyMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *KeyMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetGroupID sets the "group_id" field.
func (m *KeyMutation) SetGroupID(u uint64) {
	m.group_id = &u
	m.addgroup_id = nil
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *KeyMutation) GroupID() (r uint64, exists bool) {
	v := m.group_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldGroupID(ctx context.Context) (v uint64, err error) {
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

// AddGroupID adds u to the "group_id" field.
func (m *KeyMutation) AddGroupID(u int64) {
	if m.addgroup_id != nil {
		*m.addgroup_id += u
	} else {
		m.addgroup_id = &u
	}
}

// AddedGroupID returns the value that was added to the "group_id" field in this mutation.
func (m *KeyMutation) AddedGroupID() (r int64, exists bool) {
	v := m.addgroup_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetGroupID resets all changes to the "group_id" field.
func (m *KeyMutation) ResetGroupID() {
	m.group_id = nil
	m.addgroup_id = nil
}

// SetMemberID sets the "member_id" field.
func (m *KeyMutation) SetMemberID(u uint64) {
	m.member_id = &u
	m.addmember_id = nil
}

// MemberID returns the value of the "member_id" field in the mutation.
func (m *KeyMutation) MemberID() (r uint64, exists bool) {
	v := m.member_id
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberID returns the old "member_id" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldMemberID(ctx context.Context) (v uint64, err error) {
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

// AddMemberID adds u to the "member_id" field.
func (m *KeyMutation) AddMemberID(u int64) {
	if m.addmember_id != nil {
		*m.addmember_id += u
	} else {
		m.addmember_id = &u
	}
}

// AddedMemberID returns the value that was added to the "member_id" field in this mutation.
func (m *KeyMutation) AddedMemberID() (r int64, exists bool) {
	v := m.addmember_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetMemberID resets all changes to the "member_id" field.
func (m *KeyMutation) ResetMemberID() {
	m.member_id = nil
	m.addmember_id = nil
}

// SetKey sets the "key" field.
func (m *KeyMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *KeyMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *KeyMutation) ResetKey() {
	m.key = nil
}

// SetEnable sets the "enable" field.
func (m *KeyMutation) SetEnable(b bool) {
	m.enable = &b
}

// Enable returns the value of the "enable" field in the mutation.
func (m *KeyMutation) Enable() (r bool, exists bool) {
	v := m.enable
	if v == nil {
		return
	}
	return *v, true
}

// OldEnable returns the old "enable" field's value of the Key entity.
// If the Key object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *KeyMutation) OldEnable(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEnable is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEnable requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEnable: %w", err)
	}
	return oldValue.Enable, nil
}

// ResetEnable resets all changes to the "enable" field.
func (m *KeyMutation) ResetEnable() {
	m.enable = nil
}

// Where appends a list predicates to the KeyMutation builder.
func (m *KeyMutation) Where(ps ...predicate.Key) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *KeyMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Key).
func (m *KeyMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *KeyMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, key.FieldCreatedAt)
	}
	if m.group_id != nil {
		fields = append(fields, key.FieldGroupID)
	}
	if m.member_id != nil {
		fields = append(fields, key.FieldMemberID)
	}
	if m.key != nil {
		fields = append(fields, key.FieldKey)
	}
	if m.enable != nil {
		fields = append(fields, key.FieldEnable)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *KeyMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case key.FieldCreatedAt:
		return m.CreatedAt()
	case key.FieldGroupID:
		return m.GroupID()
	case key.FieldMemberID:
		return m.MemberID()
	case key.FieldKey:
		return m.Key()
	case key.FieldEnable:
		return m.Enable()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *KeyMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case key.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case key.FieldGroupID:
		return m.OldGroupID(ctx)
	case key.FieldMemberID:
		return m.OldMemberID(ctx)
	case key.FieldKey:
		return m.OldKey(ctx)
	case key.FieldEnable:
		return m.OldEnable(ctx)
	}
	return nil, fmt.Errorf("unknown Key field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KeyMutation) SetField(name string, value ent.Value) error {
	switch name {
	case key.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case key.FieldGroupID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case key.FieldMemberID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberID(v)
		return nil
	case key.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case key.FieldEnable:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEnable(v)
		return nil
	}
	return fmt.Errorf("unknown Key field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *KeyMutation) AddedFields() []string {
	var fields []string
	if m.addgroup_id != nil {
		fields = append(fields, key.FieldGroupID)
	}
	if m.addmember_id != nil {
		fields = append(fields, key.FieldMemberID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *KeyMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case key.FieldGroupID:
		return m.AddedGroupID()
	case key.FieldMemberID:
		return m.AddedMemberID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *KeyMutation) AddField(name string, value ent.Value) error {
	switch name {
	case key.FieldGroupID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGroupID(v)
		return nil
	case key.FieldMemberID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMemberID(v)
		return nil
	}
	return fmt.Errorf("unknown Key numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *KeyMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *KeyMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *KeyMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Key nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *KeyMutation) ResetField(name string) error {
	switch name {
	case key.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case key.FieldGroupID:
		m.ResetGroupID()
		return nil
	case key.FieldMemberID:
		m.ResetMemberID()
		return nil
	case key.FieldKey:
		m.ResetKey()
		return nil
	case key.FieldEnable:
		m.ResetEnable()
		return nil
	}
	return fmt.Errorf("unknown Key field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *KeyMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *KeyMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *KeyMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *KeyMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *KeyMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *KeyMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *KeyMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Key unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *KeyMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Key edge %s", name)
}
