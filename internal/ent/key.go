// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
)

// Key is the model entity for the Key schema.
type Key struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID string `json:"member_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID string `json:"group_id,omitempty"`
	// Keys holds the value of the "keys" field.
	Keys string `json:"keys,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KeyQuery when eager-loading is set.
	Edges KeyEdges `json:"edges"`
}

// KeyEdges holds the relations/edges for other nodes in the graph.
type KeyEdges struct {
	// Member holds the value of the member edge.
	Member *Member `json:"member,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KeyEdges) MemberOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.Member == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "member"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e KeyEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Key) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case key.FieldID, key.FieldMemberID, key.FieldGroupID, key.FieldKeys:
			values[i] = new(sql.NullString)
		case key.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Key", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Key fields.
func (k *Key) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case key.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				k.ID = value.String
			}
		case key.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				k.CreatedAt = value.Time
			}
		case key.FieldMemberID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				k.MemberID = value.String
			}
		case key.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				k.GroupID = value.String
			}
		case key.FieldKeys:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keys", values[i])
			} else if value.Valid {
				k.Keys = value.String
			}
		}
	}
	return nil
}

// QueryMember queries the "member" edge of the Key entity.
func (k *Key) QueryMember() *MemberQuery {
	return (&KeyClient{config: k.config}).QueryMember(k)
}

// QueryGroup queries the "group" edge of the Key entity.
func (k *Key) QueryGroup() *GroupQuery {
	return (&KeyClient{config: k.config}).QueryGroup(k)
}

// Update returns a builder for updating this Key.
// Note that you need to call Key.Unwrap() before calling this method if this Key
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Key) Update() *KeyUpdateOne {
	return (&KeyClient{config: k.config}).UpdateOne(k)
}

// Unwrap unwraps the Key entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Key) Unwrap() *Key {
	_tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Key is not a transactional entity")
	}
	k.config.driver = _tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Key) String() string {
	var builder strings.Builder
	builder.WriteString("Key(")
	builder.WriteString(fmt.Sprintf("id=%v, ", k.ID))
	builder.WriteString("created_at=")
	builder.WriteString(k.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(k.MemberID)
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(k.GroupID)
	builder.WriteString(", ")
	builder.WriteString("keys=")
	builder.WriteString(k.Keys)
	builder.WriteByte(')')
	return builder.String()
}

// Keys is a parsable slice of Key.
type Keys []*Key

func (k Keys) config(cfg config) {
	for _i := range k {
		k[_i].config = cfg
	}
}
