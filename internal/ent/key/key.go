// Code generated by ent, DO NOT EDIT.

package key

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the key type in the database.
	Label = "key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldEnable holds the string denoting the enable field in the database.
	FieldEnable = "enable"
	// Table holds the table name of the key in the database.
	Table = "key"
)

// Columns holds all SQL columns for key fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldGroupID,
	FieldMemberID,
	FieldKey,
	FieldEnable,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/chatpuppy/puppychat/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultEnable holds the default value on creation for the "enable" field.
	DefaultEnable bool
)