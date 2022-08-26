// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupColumns holds the columns for the "group" table.
	GroupColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "members_max", Type: field.TypeInt},
		{Name: "members_count", Type: field.TypeInt, Default: 1},
		{Name: "public", Type: field.TypeBool},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "intro", Type: field.TypeString, Nullable: true},
		{Name: "keys", Type: field.TypeJSON},
		{Name: "member_id", Type: field.TypeUint64},
	}
	// GroupTable holds the schema information for the "group" table.
	GroupTable = &schema.Table{
		Name:       "group",
		Columns:    GroupColumns,
		PrimaryKey: []*schema.Column{GroupColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_member_own_groups",
				Columns:    []*schema.Column{GroupColumns[9]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "group_created_at",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[1]},
			},
			{
				Name:    "group_name",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[2]},
				Annotation: &entsql.IndexAnnotation{
					Types: map[string]string{
						"postgres": "GIN",
					},
				},
			},
			{
				Name:    "group_member_id",
				Unique:  false,
				Columns: []*schema.Column{GroupColumns[9]},
			},
		},
	}
	// KeyColumns holds the columns for the "key" table.
	KeyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "group_id", Type: field.TypeUint64},
		{Name: "member_id", Type: field.TypeUint64},
		{Name: "key", Type: field.TypeString},
		{Name: "enable", Type: field.TypeBool, Default: true},
	}
	// KeyTable holds the schema information for the "key" table.
	KeyTable = &schema.Table{
		Name:       "key",
		Columns:    KeyColumns,
		PrimaryKey: []*schema.Column{KeyColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "key_created_at",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[1]},
			},
			{
				Name:    "key_group_id_member_id",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[2], KeyColumns[3]},
			},
			{
				Name:    "key_enable",
				Unique:  false,
				Columns: []*schema.Column{KeyColumns[5]},
			},
		},
	}
	// MemberColumns holds the columns for the "member" table.
	MemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "address", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "intro", Type: field.TypeString, Nullable: true},
		{Name: "public_key", Type: field.TypeString, Nullable: true},
		{Name: "nonce", Type: field.TypeString},
		{Name: "show_nickname", Type: field.TypeBool, Default: true},
	}
	// MemberTable holds the schema information for the "member" table.
	MemberTable = &schema.Table{
		Name:       "member",
		Columns:    MemberColumns,
		PrimaryKey: []*schema.Column{MemberColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "member_created_at",
				Unique:  false,
				Columns: []*schema.Column{MemberColumns[1]},
			},
			{
				Name:    "member_nonce",
				Unique:  false,
				Columns: []*schema.Column{MemberColumns[7]},
			},
		},
	}
	// MessageColumns holds the columns for the "message" table.
	MessageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "key_id", Type: field.TypeUint64},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "group_id", Type: field.TypeUint64},
		{Name: "member_id", Type: field.TypeUint64},
	}
	// MessageTable holds the schema information for the "message" table.
	MessageTable = &schema.Table{
		Name:       "message",
		Columns:    MessageColumns,
		PrimaryKey: []*schema.Column{MessageColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "message_group_messages",
				Columns:    []*schema.Column{MessageColumns[4]},
				RefColumns: []*schema.Column{GroupColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "message_member_messages",
				Columns:    []*schema.Column{MessageColumns[5]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "message_created_at",
				Unique:  false,
				Columns: []*schema.Column{MessageColumns[1]},
			},
		},
	}
	// GroupMembersColumns holds the columns for the "group_members" table.
	GroupMembersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeUint64},
		{Name: "member_id", Type: field.TypeUint64},
	}
	// GroupMembersTable holds the schema information for the "group_members" table.
	GroupMembersTable = &schema.Table{
		Name:       "group_members",
		Columns:    GroupMembersColumns,
		PrimaryKey: []*schema.Column{GroupMembersColumns[0], GroupMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_members_group_id",
				Columns:    []*schema.Column{GroupMembersColumns[0]},
				RefColumns: []*schema.Column{GroupColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_members_member_id",
				Columns:    []*schema.Column{GroupMembersColumns[1]},
				RefColumns: []*schema.Column{MemberColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupTable,
		KeyTable,
		MemberTable,
		MessageTable,
		GroupMembersTable,
	}
)

func init() {
	GroupTable.ForeignKeys[0].RefTable = MemberTable
	GroupTable.Annotation = &entsql.Annotation{
		Table: "group",
	}
	KeyTable.Annotation = &entsql.Annotation{
		Table: "key",
	}
	MemberTable.Annotation = &entsql.Annotation{
		Table: "member",
	}
	MessageTable.ForeignKeys[0].RefTable = GroupTable
	MessageTable.ForeignKeys[1].RefTable = MemberTable
	MessageTable.Annotation = &entsql.Annotation{
		Table: "message",
	}
	GroupMembersTable.ForeignKeys[0].RefTable = GroupTable
	GroupMembersTable.ForeignKeys[1].RefTable = MemberTable
}
