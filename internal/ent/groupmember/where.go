// Code generated by ent, DO NOT EDIT.

package groupmember

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/chatpuppy/puppychat/app/model"
	"github.com/chatpuppy/puppychat/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// Permission applies equality check predicate on the "permission" field. It's identical to PermissionEQ.
func Permission(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// InviterID applies equality check predicate on the "inviter_id" field. It's identical to InviterIDEQ.
func InviterID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviterID), v))
	})
}

// InviteCode applies equality check predicate on the "invite_code" field. It's identical to InviteCodeEQ.
func InviteCode(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteCode), v))
	})
}

// InviteExpire applies equality check predicate on the "invite_expire" field. It's identical to InviteExpireEQ.
func InviteExpire(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteExpire), v))
	})
}

// ReadID applies equality check predicate on the "read_id" field. It's identical to ReadIDEQ.
func ReadID(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadID), v))
	})
}

// ReadTime applies equality check predicate on the "read_time" field. It's identical to ReadTimeEQ.
func ReadTime(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadTime), v))
	})
}

// LastNode applies equality check predicate on the "last_node" field. It's identical to LastNodeEQ.
func LastNode(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastNode), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemberID), v))
	})
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemberID), v))
	})
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMemberID), v...))
	})
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMemberID), v...))
	})
}

// MemberIDGT applies the GT predicate on the "member_id" field.
func MemberIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemberID), v))
	})
}

// MemberIDGTE applies the GTE predicate on the "member_id" field.
func MemberIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemberID), v))
	})
}

// MemberIDLT applies the LT predicate on the "member_id" field.
func MemberIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemberID), v))
	})
}

// MemberIDLTE applies the LTE predicate on the "member_id" field.
func MemberIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemberID), v))
	})
}

// MemberIDContains applies the Contains predicate on the "member_id" field.
func MemberIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemberID), v))
	})
}

// MemberIDHasPrefix applies the HasPrefix predicate on the "member_id" field.
func MemberIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemberID), v))
	})
}

// MemberIDHasSuffix applies the HasSuffix predicate on the "member_id" field.
func MemberIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemberID), v))
	})
}

// MemberIDEqualFold applies the EqualFold predicate on the "member_id" field.
func MemberIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemberID), v))
	})
}

// MemberIDContainsFold applies the ContainsFold predicate on the "member_id" field.
func MemberIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemberID), v))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupID), v))
	})
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupID), v))
	})
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupID), v))
	})
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupID), v))
	})
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupID), v))
	})
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupID), v))
	})
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupID), v))
	})
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupID), v))
	})
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupID), v))
	})
}

// PermissionEQ applies the EQ predicate on the "permission" field.
func PermissionEQ(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// PermissionNEQ applies the NEQ predicate on the "permission" field.
func PermissionNEQ(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermission), v))
	})
}

// PermissionIn applies the In predicate on the "permission" field.
func PermissionIn(vs ...model.GroupMemberPerm) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPermission), v...))
	})
}

// PermissionNotIn applies the NotIn predicate on the "permission" field.
func PermissionNotIn(vs ...model.GroupMemberPerm) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPermission), v...))
	})
}

// PermissionGT applies the GT predicate on the "permission" field.
func PermissionGT(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermission), v))
	})
}

// PermissionGTE applies the GTE predicate on the "permission" field.
func PermissionGTE(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermission), v))
	})
}

// PermissionLT applies the LT predicate on the "permission" field.
func PermissionLT(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermission), v))
	})
}

// PermissionLTE applies the LTE predicate on the "permission" field.
func PermissionLTE(v model.GroupMemberPerm) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermission), v))
	})
}

// InviterIDEQ applies the EQ predicate on the "inviter_id" field.
func InviterIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviterID), v))
	})
}

// InviterIDNEQ applies the NEQ predicate on the "inviter_id" field.
func InviterIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviterID), v))
	})
}

// InviterIDIn applies the In predicate on the "inviter_id" field.
func InviterIDIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInviterID), v...))
	})
}

// InviterIDNotIn applies the NotIn predicate on the "inviter_id" field.
func InviterIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInviterID), v...))
	})
}

// InviterIDGT applies the GT predicate on the "inviter_id" field.
func InviterIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInviterID), v))
	})
}

// InviterIDGTE applies the GTE predicate on the "inviter_id" field.
func InviterIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInviterID), v))
	})
}

// InviterIDLT applies the LT predicate on the "inviter_id" field.
func InviterIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInviterID), v))
	})
}

// InviterIDLTE applies the LTE predicate on the "inviter_id" field.
func InviterIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInviterID), v))
	})
}

// InviterIDContains applies the Contains predicate on the "inviter_id" field.
func InviterIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInviterID), v))
	})
}

// InviterIDHasPrefix applies the HasPrefix predicate on the "inviter_id" field.
func InviterIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInviterID), v))
	})
}

// InviterIDHasSuffix applies the HasSuffix predicate on the "inviter_id" field.
func InviterIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInviterID), v))
	})
}

// InviterIDIsNil applies the IsNil predicate on the "inviter_id" field.
func InviterIDIsNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInviterID)))
	})
}

// InviterIDNotNil applies the NotNil predicate on the "inviter_id" field.
func InviterIDNotNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInviterID)))
	})
}

// InviterIDEqualFold applies the EqualFold predicate on the "inviter_id" field.
func InviterIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInviterID), v))
	})
}

// InviterIDContainsFold applies the ContainsFold predicate on the "inviter_id" field.
func InviterIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInviterID), v))
	})
}

// InviteCodeEQ applies the EQ predicate on the "invite_code" field.
func InviteCodeEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteCode), v))
	})
}

// InviteCodeNEQ applies the NEQ predicate on the "invite_code" field.
func InviteCodeNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviteCode), v))
	})
}

// InviteCodeIn applies the In predicate on the "invite_code" field.
func InviteCodeIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInviteCode), v...))
	})
}

// InviteCodeNotIn applies the NotIn predicate on the "invite_code" field.
func InviteCodeNotIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInviteCode), v...))
	})
}

// InviteCodeGT applies the GT predicate on the "invite_code" field.
func InviteCodeGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInviteCode), v))
	})
}

// InviteCodeGTE applies the GTE predicate on the "invite_code" field.
func InviteCodeGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInviteCode), v))
	})
}

// InviteCodeLT applies the LT predicate on the "invite_code" field.
func InviteCodeLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInviteCode), v))
	})
}

// InviteCodeLTE applies the LTE predicate on the "invite_code" field.
func InviteCodeLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInviteCode), v))
	})
}

// InviteCodeContains applies the Contains predicate on the "invite_code" field.
func InviteCodeContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInviteCode), v))
	})
}

// InviteCodeHasPrefix applies the HasPrefix predicate on the "invite_code" field.
func InviteCodeHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInviteCode), v))
	})
}

// InviteCodeHasSuffix applies the HasSuffix predicate on the "invite_code" field.
func InviteCodeHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInviteCode), v))
	})
}

// InviteCodeEqualFold applies the EqualFold predicate on the "invite_code" field.
func InviteCodeEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInviteCode), v))
	})
}

// InviteCodeContainsFold applies the ContainsFold predicate on the "invite_code" field.
func InviteCodeContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInviteCode), v))
	})
}

// InviteExpireEQ applies the EQ predicate on the "invite_expire" field.
func InviteExpireEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteExpire), v))
	})
}

// InviteExpireNEQ applies the NEQ predicate on the "invite_expire" field.
func InviteExpireNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviteExpire), v))
	})
}

// InviteExpireIn applies the In predicate on the "invite_expire" field.
func InviteExpireIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInviteExpire), v...))
	})
}

// InviteExpireNotIn applies the NotIn predicate on the "invite_expire" field.
func InviteExpireNotIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInviteExpire), v...))
	})
}

// InviteExpireGT applies the GT predicate on the "invite_expire" field.
func InviteExpireGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInviteExpire), v))
	})
}

// InviteExpireGTE applies the GTE predicate on the "invite_expire" field.
func InviteExpireGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInviteExpire), v))
	})
}

// InviteExpireLT applies the LT predicate on the "invite_expire" field.
func InviteExpireLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInviteExpire), v))
	})
}

// InviteExpireLTE applies the LTE predicate on the "invite_expire" field.
func InviteExpireLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInviteExpire), v))
	})
}

// ReadIDEQ applies the EQ predicate on the "read_id" field.
func ReadIDEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadID), v))
	})
}

// ReadIDNEQ applies the NEQ predicate on the "read_id" field.
func ReadIDNEQ(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadID), v))
	})
}

// ReadIDIn applies the In predicate on the "read_id" field.
func ReadIDIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReadID), v...))
	})
}

// ReadIDNotIn applies the NotIn predicate on the "read_id" field.
func ReadIDNotIn(vs ...string) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReadID), v...))
	})
}

// ReadIDGT applies the GT predicate on the "read_id" field.
func ReadIDGT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReadID), v))
	})
}

// ReadIDGTE applies the GTE predicate on the "read_id" field.
func ReadIDGTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReadID), v))
	})
}

// ReadIDLT applies the LT predicate on the "read_id" field.
func ReadIDLT(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReadID), v))
	})
}

// ReadIDLTE applies the LTE predicate on the "read_id" field.
func ReadIDLTE(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReadID), v))
	})
}

// ReadIDContains applies the Contains predicate on the "read_id" field.
func ReadIDContains(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReadID), v))
	})
}

// ReadIDHasPrefix applies the HasPrefix predicate on the "read_id" field.
func ReadIDHasPrefix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReadID), v))
	})
}

// ReadIDHasSuffix applies the HasSuffix predicate on the "read_id" field.
func ReadIDHasSuffix(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReadID), v))
	})
}

// ReadIDIsNil applies the IsNil predicate on the "read_id" field.
func ReadIDIsNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReadID)))
	})
}

// ReadIDNotNil applies the NotNil predicate on the "read_id" field.
func ReadIDNotNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReadID)))
	})
}

// ReadIDEqualFold applies the EqualFold predicate on the "read_id" field.
func ReadIDEqualFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReadID), v))
	})
}

// ReadIDContainsFold applies the ContainsFold predicate on the "read_id" field.
func ReadIDContainsFold(v string) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReadID), v))
	})
}

// ReadTimeEQ applies the EQ predicate on the "read_time" field.
func ReadTimeEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadTime), v))
	})
}

// ReadTimeNEQ applies the NEQ predicate on the "read_time" field.
func ReadTimeNEQ(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadTime), v))
	})
}

// ReadTimeIn applies the In predicate on the "read_time" field.
func ReadTimeIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReadTime), v...))
	})
}

// ReadTimeNotIn applies the NotIn predicate on the "read_time" field.
func ReadTimeNotIn(vs ...time.Time) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReadTime), v...))
	})
}

// ReadTimeGT applies the GT predicate on the "read_time" field.
func ReadTimeGT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReadTime), v))
	})
}

// ReadTimeGTE applies the GTE predicate on the "read_time" field.
func ReadTimeGTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReadTime), v))
	})
}

// ReadTimeLT applies the LT predicate on the "read_time" field.
func ReadTimeLT(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReadTime), v))
	})
}

// ReadTimeLTE applies the LTE predicate on the "read_time" field.
func ReadTimeLTE(v time.Time) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReadTime), v))
	})
}

// ReadTimeIsNil applies the IsNil predicate on the "read_time" field.
func ReadTimeIsNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReadTime)))
	})
}

// ReadTimeNotNil applies the NotNil predicate on the "read_time" field.
func ReadTimeNotNil() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReadTime)))
	})
}

// LastNodeEQ applies the EQ predicate on the "last_node" field.
func LastNodeEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastNode), v))
	})
}

// LastNodeNEQ applies the NEQ predicate on the "last_node" field.
func LastNodeNEQ(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastNode), v))
	})
}

// LastNodeIn applies the In predicate on the "last_node" field.
func LastNodeIn(vs ...int64) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastNode), v...))
	})
}

// LastNodeNotIn applies the NotIn predicate on the "last_node" field.
func LastNodeNotIn(vs ...int64) predicate.GroupMember {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastNode), v...))
	})
}

// LastNodeGT applies the GT predicate on the "last_node" field.
func LastNodeGT(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastNode), v))
	})
}

// LastNodeGTE applies the GTE predicate on the "last_node" field.
func LastNodeGTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastNode), v))
	})
}

// LastNodeLT applies the LT predicate on the "last_node" field.
func LastNodeLT(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastNode), v))
	})
}

// LastNodeLTE applies the LTE predicate on the "last_node" field.
func LastNodeLTE(v int64) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastNode), v))
	})
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInviter applies the HasEdge predicate on the "inviter" edge.
func HasInviter() predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InviterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InviterTable, InviterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInviterWith applies the HasEdge predicate on the "inviter" edge with a given conditions (other predicates).
func HasInviterWith(preds ...predicate.Member) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InviterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InviterTable, InviterColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupMember) predicate.GroupMember {
	return predicate.GroupMember(func(s *sql.Selector) {
		p(s.Not())
	})
}
