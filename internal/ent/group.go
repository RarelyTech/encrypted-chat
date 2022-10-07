// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/member"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// created by
	OwnerID string `json:"owner_id,omitempty"`
	// MembersMax holds the value of the "members_max" field.
	MembersMax int `json:"members_max,omitempty"`
	// members count of group
	MembersCount int `json:"members_count,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Intro holds the value of the "intro" field.
	Intro string `json:"intro,omitempty"`
	// group's ethereum keys
	Keys string `json:"keys,omitempty"`
	// LastNode holds the value of the "last_node" field.
	LastNode int64 `json:"last_node,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges GroupEdges `json:"edges"`
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Member `json:"owner,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// Members holds the value of the members edge.
	Members []*Member `json:"members,omitempty"`
	// GroupMembers holds the value of the group_members edge.
	GroupMembers []*GroupMember `json:"group_members,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) OwnerOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[1] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) MembersOrErr() ([]*Member, error) {
	if e.loadedTypes[2] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// GroupMembersOrErr returns the GroupMembers value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GroupMembersOrErr() ([]*GroupMember, error) {
	if e.loadedTypes[3] {
		return e.GroupMembers, nil
	}
	return nil, &NotLoadedError{edge: "group_members"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldPublic:
			values[i] = new(sql.NullBool)
		case group.FieldMembersMax, group.FieldMembersCount, group.FieldLastNode:
			values[i] = new(sql.NullInt64)
		case group.FieldID, group.FieldName, group.FieldCategory, group.FieldOwnerID, group.FieldAddress, group.FieldIntro, group.FieldKeys:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Group", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gr.ID = value.String
			}
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				gr.Category = value.String
			}
		case group.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				gr.OwnerID = value.String
			}
		case group.FieldMembersMax:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field members_max", values[i])
			} else if value.Valid {
				gr.MembersMax = int(value.Int64)
			}
		case group.FieldMembersCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field members_count", values[i])
			} else if value.Valid {
				gr.MembersCount = int(value.Int64)
			}
		case group.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				gr.Public = value.Bool
			}
		case group.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				gr.Address = value.String
			}
		case group.FieldIntro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intro", values[i])
			} else if value.Valid {
				gr.Intro = value.String
			}
		case group.FieldKeys:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keys", values[i])
			} else if value.Valid {
				gr.Keys = value.String
			}
		case group.FieldLastNode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_node", values[i])
			} else if value.Valid {
				gr.LastNode = value.Int64
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Group entity.
func (gr *Group) QueryOwner() *MemberQuery {
	return (&GroupClient{config: gr.config}).QueryOwner(gr)
}

// QueryMessages queries the "messages" edge of the Group entity.
func (gr *Group) QueryMessages() *MessageQuery {
	return (&GroupClient{config: gr.config}).QueryMessages(gr)
}

// QueryMembers queries the "members" edge of the Group entity.
func (gr *Group) QueryMembers() *MemberQuery {
	return (&GroupClient{config: gr.config}).QueryMembers(gr)
}

// QueryGroupMembers queries the "group_members" edge of the Group entity.
func (gr *Group) QueryGroupMembers() *GroupMemberQuery {
	return (&GroupClient{config: gr.config}).QueryGroupMembers(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(gr.Category)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(gr.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("members_max=")
	builder.WriteString(fmt.Sprintf("%v", gr.MembersMax))
	builder.WriteString(", ")
	builder.WriteString("members_count=")
	builder.WriteString(fmt.Sprintf("%v", gr.MembersCount))
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", gr.Public))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(gr.Address)
	builder.WriteString(", ")
	builder.WriteString("intro=")
	builder.WriteString(gr.Intro)
	builder.WriteString(", ")
	builder.WriteString("keys=")
	builder.WriteString(gr.Keys)
	builder.WriteString(", ")
	builder.WriteString("last_node=")
	builder.WriteString(fmt.Sprintf("%v", gr.LastNode))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
