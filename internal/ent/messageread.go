// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/messageread"
)

// MessageRead is the model entity for the MessageRead schema.
type MessageRead struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// MemberID holds the value of the "member_id" field.
	MemberID string `json:"member_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID string `json:"group_id,omitempty"`
	// LastID holds the value of the "last_id" field.
	LastID string `json:"last_id,omitempty"`
	// LastTime holds the value of the "last_time" field.
	LastTime time.Time `json:"last_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageReadQuery when eager-loading is set.
	Edges MessageReadEdges `json:"edges"`
}

// MessageReadEdges holds the relations/edges for other nodes in the graph.
type MessageReadEdges struct {
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
func (e MessageReadEdges) MemberOrErr() (*Member, error) {
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
func (e MessageReadEdges) GroupOrErr() (*Group, error) {
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
func (*MessageRead) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case messageread.FieldID, messageread.FieldMemberID, messageread.FieldGroupID, messageread.FieldLastID:
			values[i] = new(sql.NullString)
		case messageread.FieldLastTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MessageRead", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageRead fields.
func (mr *MessageRead) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messageread.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mr.ID = value.String
			}
		case messageread.FieldMemberID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				mr.MemberID = value.String
			}
		case messageread.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				mr.GroupID = value.String
			}
		case messageread.FieldLastID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_id", values[i])
			} else if value.Valid {
				mr.LastID = value.String
			}
		case messageread.FieldLastTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_time", values[i])
			} else if value.Valid {
				mr.LastTime = value.Time
			}
		}
	}
	return nil
}

// QueryMember queries the "member" edge of the MessageRead entity.
func (mr *MessageRead) QueryMember() *MemberQuery {
	return (&MessageReadClient{config: mr.config}).QueryMember(mr)
}

// QueryGroup queries the "group" edge of the MessageRead entity.
func (mr *MessageRead) QueryGroup() *GroupQuery {
	return (&MessageReadClient{config: mr.config}).QueryGroup(mr)
}

// Update returns a builder for updating this MessageRead.
// Note that you need to call MessageRead.Unwrap() before calling this method if this MessageRead
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MessageRead) Update() *MessageReadUpdateOne {
	return (&MessageReadClient{config: mr.config}).UpdateOne(mr)
}

// Unwrap unwraps the MessageRead entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MessageRead) Unwrap() *MessageRead {
	_tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageRead is not a transactional entity")
	}
	mr.config.driver = _tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MessageRead) String() string {
	var builder strings.Builder
	builder.WriteString("MessageRead(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mr.ID))
	builder.WriteString("member_id=")
	builder.WriteString(mr.MemberID)
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(mr.GroupID)
	builder.WriteString(", ")
	builder.WriteString("last_id=")
	builder.WriteString(mr.LastID)
	builder.WriteString(", ")
	builder.WriteString("last_time=")
	builder.WriteString(mr.LastTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MessageReads is a parsable slice of MessageRead.
type MessageReads []*MessageRead

func (mr MessageReads) config(cfg config) {
	for _i := range mr {
		mr[_i].config = cfg
	}
}