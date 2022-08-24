// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/key"
)

// Key is the model entity for the Key schema.
type Key struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID uint64 `json:"group_id,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID uint64 `json:"member_id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Enable holds the value of the "enable" field.
	Enable bool `json:"enable,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Key) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case key.FieldEnable:
			values[i] = new(sql.NullBool)
		case key.FieldID, key.FieldGroupID, key.FieldMemberID:
			values[i] = new(sql.NullInt64)
		case key.FieldKey:
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
func (k *Key) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case key.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			k.ID = uint64(value.Int64)
		case key.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				k.CreatedAt = value.Time
			}
		case key.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				k.GroupID = uint64(value.Int64)
			}
		case key.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				k.MemberID = uint64(value.Int64)
			}
		case key.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				k.Key = value.String
			}
		case key.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable", values[i])
			} else if value.Valid {
				k.Enable = value.Bool
			}
		}
	}
	return nil
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
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", k.GroupID))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", k.MemberID))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(k.Key)
	builder.WriteString(", ")
	builder.WriteString("enable=")
	builder.WriteString(fmt.Sprintf("%v", k.Enable))
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
