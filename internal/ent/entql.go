// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: group.FieldID,
			},
		},
		Type: "Group",
		Fields: map[string]*sqlgraph.FieldSpec{
			group.FieldCreatedAt:    {Type: field.TypeTime, Column: group.FieldCreatedAt},
			group.FieldName:         {Type: field.TypeString, Column: group.FieldName},
			group.FieldMemberID:     {Type: field.TypeUint64, Column: group.FieldMemberID},
			group.FieldMembersMax:   {Type: field.TypeInt, Column: group.FieldMembersMax},
			group.FieldMembersCount: {Type: field.TypeInt, Column: group.FieldMembersCount},
			group.FieldPublic:       {Type: field.TypeBool, Column: group.FieldPublic},
			group.FieldSn:           {Type: field.TypeString, Column: group.FieldSn},
			group.FieldIntro:        {Type: field.TypeString, Column: group.FieldIntro},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   key.Table,
			Columns: key.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: key.FieldID,
			},
		},
		Type: "Key",
		Fields: map[string]*sqlgraph.FieldSpec{
			key.FieldCreatedAt: {Type: field.TypeTime, Column: key.FieldCreatedAt},
			key.FieldGroupID:   {Type: field.TypeUint64, Column: key.FieldGroupID},
			key.FieldMemberID:  {Type: field.TypeUint64, Column: key.FieldMemberID},
			key.FieldKey:       {Type: field.TypeString, Column: key.FieldKey},
			key.FieldEnable:    {Type: field.TypeBool, Column: key.FieldEnable},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: member.FieldID,
			},
		},
		Type: "Member",
		Fields: map[string]*sqlgraph.FieldSpec{
			member.FieldCreatedAt:    {Type: field.TypeTime, Column: member.FieldCreatedAt},
			member.FieldAddress:      {Type: field.TypeString, Column: member.FieldAddress},
			member.FieldNickname:     {Type: field.TypeString, Column: member.FieldNickname},
			member.FieldAvatar:       {Type: field.TypeString, Column: member.FieldAvatar},
			member.FieldIntro:        {Type: field.TypeString, Column: member.FieldIntro},
			member.FieldNonce:        {Type: field.TypeString, Column: member.FieldNonce},
			member.FieldShowNickname: {Type: field.TypeBool, Column: member.FieldShowNickname},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   message.Table,
			Columns: message.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: message.FieldID,
			},
		},
		Type: "Message",
		Fields: map[string]*sqlgraph.FieldSpec{
			message.FieldCreatedAt: {Type: field.TypeTime, Column: message.FieldCreatedAt},
			message.FieldKeyID:     {Type: field.TypeUint64, Column: message.FieldKeyID},
			message.FieldGroupID:   {Type: field.TypeUint64, Column: message.FieldGroupID},
			message.FieldMemberID:  {Type: field.TypeUint64, Column: message.FieldMemberID},
			message.FieldContent:   {Type: field.TypeString, Column: message.FieldContent},
		},
	}
	graph.MustAddE(
		"owner",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
		},
		"Group",
		"Member",
	)
	graph.MustAddE(
		"messages",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MessagesTable,
			Columns: []string{group.MessagesColumn},
			Bidi:    false,
		},
		"Group",
		"Message",
	)
	graph.MustAddE(
		"members",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: group.MembersPrimaryKey,
			Bidi:    false,
		},
		"Group",
		"Member",
	)
	graph.MustAddE(
		"own_groups",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.OwnGroupsTable,
			Columns: []string{member.OwnGroupsColumn},
			Bidi:    false,
		},
		"Member",
		"Group",
	)
	graph.MustAddE(
		"messages",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MessagesTable,
			Columns: []string{member.MessagesColumn},
			Bidi:    false,
		},
		"Member",
		"Message",
	)
	graph.MustAddE(
		"groups",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   member.GroupsTable,
			Columns: member.GroupsPrimaryKey,
			Bidi:    false,
		},
		"Member",
		"Group",
	)
	graph.MustAddE(
		"owner",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.OwnerTable,
			Columns: []string{message.OwnerColumn},
			Bidi:    false,
		},
		"Message",
		"Member",
	)
	graph.MustAddE(
		"group",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.GroupTable,
			Columns: []string{message.GroupColumn},
			Bidi:    false,
		},
		"Message",
		"Group",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (gq *GroupQuery) addPredicate(pred func(s *sql.Selector)) {
	gq.predicates = append(gq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GroupQuery builder.
func (gq *GroupQuery) Filter() *GroupFilter {
	return &GroupFilter{config: gq.config, predicateAdder: gq}
}

// addPredicate implements the predicateAdder interface.
func (m *GroupMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GroupMutation builder.
func (m *GroupMutation) Filter() *GroupFilter {
	return &GroupFilter{config: m.config, predicateAdder: m}
}

// GroupFilter provides a generic filtering capability at runtime for GroupQuery.
type GroupFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GroupFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint64 predicate on the id field.
func (f *GroupFilter) WhereID(p entql.Uint64P) {
	f.Where(p.Field(group.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *GroupFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(group.FieldCreatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *GroupFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(group.FieldName))
}

// WhereMemberID applies the entql uint64 predicate on the member_id field.
func (f *GroupFilter) WhereMemberID(p entql.Uint64P) {
	f.Where(p.Field(group.FieldMemberID))
}

// WhereMembersMax applies the entql int predicate on the members_max field.
func (f *GroupFilter) WhereMembersMax(p entql.IntP) {
	f.Where(p.Field(group.FieldMembersMax))
}

// WhereMembersCount applies the entql int predicate on the members_count field.
func (f *GroupFilter) WhereMembersCount(p entql.IntP) {
	f.Where(p.Field(group.FieldMembersCount))
}

// WherePublic applies the entql bool predicate on the public field.
func (f *GroupFilter) WherePublic(p entql.BoolP) {
	f.Where(p.Field(group.FieldPublic))
}

// WhereSn applies the entql string predicate on the sn field.
func (f *GroupFilter) WhereSn(p entql.StringP) {
	f.Where(p.Field(group.FieldSn))
}

// WhereIntro applies the entql string predicate on the intro field.
func (f *GroupFilter) WhereIntro(p entql.StringP) {
	f.Where(p.Field(group.FieldIntro))
}

// WhereHasOwner applies a predicate to check if query has an edge owner.
func (f *GroupFilter) WhereHasOwner() {
	f.Where(entql.HasEdge("owner"))
}

// WhereHasOwnerWith applies a predicate to check if query has an edge owner with a given conditions (other predicates).
func (f *GroupFilter) WhereHasOwnerWith(preds ...predicate.Member) {
	f.Where(entql.HasEdgeWith("owner", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasMessages applies a predicate to check if query has an edge messages.
func (f *GroupFilter) WhereHasMessages() {
	f.Where(entql.HasEdge("messages"))
}

// WhereHasMessagesWith applies a predicate to check if query has an edge messages with a given conditions (other predicates).
func (f *GroupFilter) WhereHasMessagesWith(preds ...predicate.Message) {
	f.Where(entql.HasEdgeWith("messages", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasMembers applies a predicate to check if query has an edge members.
func (f *GroupFilter) WhereHasMembers() {
	f.Where(entql.HasEdge("members"))
}

// WhereHasMembersWith applies a predicate to check if query has an edge members with a given conditions (other predicates).
func (f *GroupFilter) WhereHasMembersWith(preds ...predicate.Member) {
	f.Where(entql.HasEdgeWith("members", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (kq *KeyQuery) addPredicate(pred func(s *sql.Selector)) {
	kq.predicates = append(kq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the KeyQuery builder.
func (kq *KeyQuery) Filter() *KeyFilter {
	return &KeyFilter{config: kq.config, predicateAdder: kq}
}

// addPredicate implements the predicateAdder interface.
func (m *KeyMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the KeyMutation builder.
func (m *KeyMutation) Filter() *KeyFilter {
	return &KeyFilter{config: m.config, predicateAdder: m}
}

// KeyFilter provides a generic filtering capability at runtime for KeyQuery.
type KeyFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *KeyFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint64 predicate on the id field.
func (f *KeyFilter) WhereID(p entql.Uint64P) {
	f.Where(p.Field(key.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *KeyFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(key.FieldCreatedAt))
}

// WhereGroupID applies the entql uint64 predicate on the group_id field.
func (f *KeyFilter) WhereGroupID(p entql.Uint64P) {
	f.Where(p.Field(key.FieldGroupID))
}

// WhereMemberID applies the entql uint64 predicate on the member_id field.
func (f *KeyFilter) WhereMemberID(p entql.Uint64P) {
	f.Where(p.Field(key.FieldMemberID))
}

// WhereKey applies the entql string predicate on the key field.
func (f *KeyFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(key.FieldKey))
}

// WhereEnable applies the entql bool predicate on the enable field.
func (f *KeyFilter) WhereEnable(p entql.BoolP) {
	f.Where(p.Field(key.FieldEnable))
}

// addPredicate implements the predicateAdder interface.
func (mq *MemberQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MemberQuery builder.
func (mq *MemberQuery) Filter() *MemberFilter {
	return &MemberFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MemberMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MemberMutation builder.
func (m *MemberMutation) Filter() *MemberFilter {
	return &MemberFilter{config: m.config, predicateAdder: m}
}

// MemberFilter provides a generic filtering capability at runtime for MemberQuery.
type MemberFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MemberFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint64 predicate on the id field.
func (f *MemberFilter) WhereID(p entql.Uint64P) {
	f.Where(p.Field(member.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MemberFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(member.FieldCreatedAt))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *MemberFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(member.FieldAddress))
}

// WhereNickname applies the entql string predicate on the nickname field.
func (f *MemberFilter) WhereNickname(p entql.StringP) {
	f.Where(p.Field(member.FieldNickname))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *MemberFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(member.FieldAvatar))
}

// WhereIntro applies the entql string predicate on the intro field.
func (f *MemberFilter) WhereIntro(p entql.StringP) {
	f.Where(p.Field(member.FieldIntro))
}

// WhereNonce applies the entql string predicate on the nonce field.
func (f *MemberFilter) WhereNonce(p entql.StringP) {
	f.Where(p.Field(member.FieldNonce))
}

// WhereShowNickname applies the entql bool predicate on the show_nickname field.
func (f *MemberFilter) WhereShowNickname(p entql.BoolP) {
	f.Where(p.Field(member.FieldShowNickname))
}

// WhereHasOwnGroups applies a predicate to check if query has an edge own_groups.
func (f *MemberFilter) WhereHasOwnGroups() {
	f.Where(entql.HasEdge("own_groups"))
}

// WhereHasOwnGroupsWith applies a predicate to check if query has an edge own_groups with a given conditions (other predicates).
func (f *MemberFilter) WhereHasOwnGroupsWith(preds ...predicate.Group) {
	f.Where(entql.HasEdgeWith("own_groups", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasMessages applies a predicate to check if query has an edge messages.
func (f *MemberFilter) WhereHasMessages() {
	f.Where(entql.HasEdge("messages"))
}

// WhereHasMessagesWith applies a predicate to check if query has an edge messages with a given conditions (other predicates).
func (f *MemberFilter) WhereHasMessagesWith(preds ...predicate.Message) {
	f.Where(entql.HasEdgeWith("messages", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasGroups applies a predicate to check if query has an edge groups.
func (f *MemberFilter) WhereHasGroups() {
	f.Where(entql.HasEdge("groups"))
}

// WhereHasGroupsWith applies a predicate to check if query has an edge groups with a given conditions (other predicates).
func (f *MemberFilter) WhereHasGroupsWith(preds ...predicate.Group) {
	f.Where(entql.HasEdgeWith("groups", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (mq *MessageQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MessageQuery builder.
func (mq *MessageQuery) Filter() *MessageFilter {
	return &MessageFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MessageMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MessageMutation builder.
func (m *MessageMutation) Filter() *MessageFilter {
	return &MessageFilter{config: m.config, predicateAdder: m}
}

// MessageFilter provides a generic filtering capability at runtime for MessageQuery.
type MessageFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MessageFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint64 predicate on the id field.
func (f *MessageFilter) WhereID(p entql.Uint64P) {
	f.Where(p.Field(message.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *MessageFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(message.FieldCreatedAt))
}

// WhereKeyID applies the entql uint64 predicate on the key_id field.
func (f *MessageFilter) WhereKeyID(p entql.Uint64P) {
	f.Where(p.Field(message.FieldKeyID))
}

// WhereGroupID applies the entql uint64 predicate on the group_id field.
func (f *MessageFilter) WhereGroupID(p entql.Uint64P) {
	f.Where(p.Field(message.FieldGroupID))
}

// WhereMemberID applies the entql uint64 predicate on the member_id field.
func (f *MessageFilter) WhereMemberID(p entql.Uint64P) {
	f.Where(p.Field(message.FieldMemberID))
}

// WhereContent applies the entql string predicate on the content field.
func (f *MessageFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(message.FieldContent))
}

// WhereHasOwner applies a predicate to check if query has an edge owner.
func (f *MessageFilter) WhereHasOwner() {
	f.Where(entql.HasEdge("owner"))
}

// WhereHasOwnerWith applies a predicate to check if query has an edge owner with a given conditions (other predicates).
func (f *MessageFilter) WhereHasOwnerWith(preds ...predicate.Member) {
	f.Where(entql.HasEdgeWith("owner", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasGroup applies a predicate to check if query has an edge group.
func (f *MessageFilter) WhereHasGroup() {
	f.Where(entql.HasEdge("group"))
}

// WhereHasGroupWith applies a predicate to check if query has an edge group with a given conditions (other predicates).
func (f *MessageFilter) WhereHasGroupWith(preds ...predicate.Group) {
	f.Where(entql.HasEdgeWith("group", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
