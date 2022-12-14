// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"time"

	"github.com/chatpuppy/puppychat/app/model"
	"github.com/chatpuppy/puppychat/internal/ent/group"
	"github.com/chatpuppy/puppychat/internal/ent/groupmember"
	"github.com/chatpuppy/puppychat/internal/ent/key"
	"github.com/chatpuppy/puppychat/internal/ent/member"
	"github.com/chatpuppy/puppychat/internal/ent/message"
	jsoniter "github.com/json-iterator/go"

	log "github.com/sirupsen/logrus"
)

type GroupSync struct {
	ID           string    `json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	Name         string    `json:"name,omitempty"`
	Category     string    `json:"category,omitempty"`
	OwnerID      string    `json:"owner_id,omitempty"`
	MembersMax   int       `json:"members_max,omitempty"`
	MembersCount int       `json:"members_count,omitempty"`
	Public       bool      `json:"public,omitempty"`
	Address      string    `json:"address,omitempty"`
	Intro        string    `json:"intro,omitempty"`
	Keys         string    `json:"keys,omitempty"`
}

func (m *GroupMutation) SyncData() (data *GroupSync) {
	ctx := context.Background()

	id, ok := m.ID()
	if !ok {
		log.Error("get Group sync data failed: id missing")
		return
	}

	if m.op.Is(OpDelete | OpDeleteOne) {
		data = &GroupSync{ID: id}
	} else {
		result, err := m.Client().Group.Get(ctx, id)
		if err != nil {
			log.Errorf("get Group sync data failed: %v", err)
			return
		}
		data = &GroupSync{
			ID:           id,
			CreatedAt:    result.CreatedAt,
			Name:         result.Name,
			Category:     result.Category,
			OwnerID:      result.OwnerID,
			MembersMax:   result.MembersMax,
			MembersCount: result.MembersCount,
			Public:       result.Public,
			Address:      result.Address,
			Intro:        result.Intro,
			Keys:         result.Keys,
		}
	}

	return data
}

func (m *GroupMutation) SetSyncData(data *GroupSync) {
	m.SetField(group.FieldCreatedAt, data.CreatedAt)
	m.SetField(group.FieldName, data.Name)
	m.SetField(group.FieldCategory, data.Category)
	m.SetField(group.FieldOwnerID, data.OwnerID)
	m.SetField(group.FieldMembersMax, data.MembersMax)
	m.SetField(group.FieldMembersCount, data.MembersCount)
	m.SetField(group.FieldPublic, data.Public)
	m.SetField(group.FieldAddress, data.Address)
	m.SetField(group.FieldIntro, data.Intro)
	m.SetField(group.FieldKeys, data.Keys)
}

func SaveGroupSyncData(ctx context.Context, b []byte, op Op, precall func(*GroupSync)) (err error) {
	data := new(GroupSync)
	err = jsoniter.Unmarshal(b, data)
	if err != nil {
		return
	}

	id := data.ID
	if id == "" {
		err = model.ErrSyncIDNotFound
		return
	}

	if op.Is(OpDelete | OpDeleteOne) {
		err = Database.Group.DeleteOneID(id).Exec(ctx)
	} else {
		creator := Database.Group.Create().SetID(id)
		if precall != nil {
			precall(data)
		}
		creator.Mutation().SetSyncData(data)
		columns := []string{
			"id",
			// group.FieldAddress,
		}
		err = creator.OnConflictColumns(columns...).UpdateNewValues().Exec(ctx)
	}

	return
}

type GroupMemberSync struct {
	ID           string                `json:"id,omitempty"`
	CreatedAt    time.Time             `json:"created_at,omitempty"`
	MemberID     string                `json:"member_id,omitempty"`
	GroupID      string                `json:"group_id,omitempty"`
	Permission   model.GroupMemberPerm `json:"permission,omitempty"`
	InviterID    *string               `json:"inviter_id,omitempty"`
	InviteCode   string                `json:"invite_code,omitempty"`
	InviteExpire time.Time             `json:"invite_expire,omitempty"`
	ReadID       *string               `json:"read_id,omitempty"`
	ReadTime     *time.Time            `json:"read_time,omitempty"`
}

func (m *GroupMemberMutation) SyncData() (data *GroupMemberSync) {
	ctx := context.Background()

	id, ok := m.ID()
	if !ok {
		log.Error("get GroupMember sync data failed: id missing")
		return
	}

	if m.op.Is(OpDelete | OpDeleteOne) {
		data = &GroupMemberSync{ID: id}
	} else {
		result, err := m.Client().GroupMember.Get(ctx, id)
		if err != nil {
			log.Errorf("get GroupMember sync data failed: %v", err)
			return
		}
		data = &GroupMemberSync{
			ID:           id,
			CreatedAt:    result.CreatedAt,
			MemberID:     result.MemberID,
			GroupID:      result.GroupID,
			Permission:   result.Permission,
			InviterID:    result.InviterID,
			InviteCode:   result.InviteCode,
			InviteExpire: result.InviteExpire,
			ReadID:       result.ReadID,
			ReadTime:     result.ReadTime,
		}
	}

	return data
}

func (m *GroupMemberMutation) SetSyncData(data *GroupMemberSync) {
	m.SetField(groupmember.FieldCreatedAt, data.CreatedAt)
	m.SetField(groupmember.FieldMemberID, data.MemberID)
	m.SetField(groupmember.FieldGroupID, data.GroupID)
	m.SetField(groupmember.FieldPermission, data.Permission)
	m.SetField(groupmember.FieldInviterID, data.InviterID)
	m.SetField(groupmember.FieldInviteCode, data.InviteCode)
	m.SetField(groupmember.FieldInviteExpire, data.InviteExpire)
	m.SetField(groupmember.FieldReadID, data.ReadID)
	m.SetField(groupmember.FieldReadTime, data.ReadTime)
}

func SaveGroupMemberSyncData(ctx context.Context, b []byte, op Op, precall func(*GroupMemberSync)) (err error) {
	data := new(GroupMemberSync)
	err = jsoniter.Unmarshal(b, data)
	if err != nil {
		return
	}

	id := data.ID
	if id == "" {
		err = model.ErrSyncIDNotFound
		return
	}

	if op.Is(OpDelete | OpDeleteOne) {
		err = Database.GroupMember.DeleteOneID(id).Exec(ctx)
	} else {
		creator := Database.GroupMember.Create().SetID(id)
		if precall != nil {
			precall(data)
		}
		creator.Mutation().SetSyncData(data)
		columns := []string{
			"id",
			// groupmember.FieldInviteCode,
		}
		err = creator.OnConflictColumns(columns...).UpdateNewValues().Exec(ctx)
	}

	return
}

type KeySync struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	MemberID  string    `json:"member_id,omitempty"`
	GroupID   string    `json:"group_id,omitempty"`
	Keys      string    `json:"keys,omitempty"`
}

func (m *KeyMutation) SyncData() (data *KeySync) {
	ctx := context.Background()

	id, ok := m.ID()
	if !ok {
		log.Error("get Key sync data failed: id missing")
		return
	}

	if m.op.Is(OpDelete | OpDeleteOne) {
		data = &KeySync{ID: id}
	} else {
		result, err := m.Client().Key.Get(ctx, id)
		if err != nil {
			log.Errorf("get Key sync data failed: %v", err)
			return
		}
		data = &KeySync{
			ID:        id,
			CreatedAt: result.CreatedAt,
			MemberID:  result.MemberID,
			GroupID:   result.GroupID,
			Keys:      result.Keys,
		}
	}

	return data
}

func (m *KeyMutation) SetSyncData(data *KeySync) {
	m.SetField(key.FieldCreatedAt, data.CreatedAt)
	m.SetField(key.FieldMemberID, data.MemberID)
	m.SetField(key.FieldGroupID, data.GroupID)
	m.SetField(key.FieldKeys, data.Keys)
}

func SaveKeySyncData(ctx context.Context, b []byte, op Op, precall func(*KeySync)) (err error) {
	data := new(KeySync)
	err = jsoniter.Unmarshal(b, data)
	if err != nil {
		return
	}

	id := data.ID
	if id == "" {
		err = model.ErrSyncIDNotFound
		return
	}

	if op.Is(OpDelete | OpDeleteOne) {
		err = Database.Key.DeleteOneID(id).Exec(ctx)
	} else {
		creator := Database.Key.Create().SetID(id)
		if precall != nil {
			precall(data)
		}
		creator.Mutation().SetSyncData(data)
		columns := []string{
			"id",
		}
		err = creator.OnConflictColumns(columns...).UpdateNewValues().Exec(ctx)
	}

	return
}

type MemberSync struct {
	ID           string    `json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	Address      string    `json:"address,omitempty"`
	Nickname     string    `json:"nickname,omitempty"`
	Avatar       string    `json:"avatar,omitempty"`
	Intro        string    `json:"intro,omitempty"`
	PublicKey    string    `json:"public_key,omitempty"`
	Nonce        string    `json:"nonce,omitempty"`
	ShowNickname bool      `json:"show_nickname,omitempty"`
}

func (m *MemberMutation) SyncData() (data *MemberSync) {
	ctx := context.Background()

	id, ok := m.ID()
	if !ok {
		log.Error("get Member sync data failed: id missing")
		return
	}

	if m.op.Is(OpDelete | OpDeleteOne) {
		data = &MemberSync{ID: id}
	} else {
		result, err := m.Client().Member.Get(ctx, id)
		if err != nil {
			log.Errorf("get Member sync data failed: %v", err)
			return
		}
		data = &MemberSync{
			ID:           id,
			CreatedAt:    result.CreatedAt,
			Address:      result.Address,
			Nickname:     result.Nickname,
			Avatar:       result.Avatar,
			Intro:        result.Intro,
			PublicKey:    result.PublicKey,
			Nonce:        result.Nonce,
			ShowNickname: result.ShowNickname,
		}
	}

	return data
}

func (m *MemberMutation) SetSyncData(data *MemberSync) {
	m.SetField(member.FieldCreatedAt, data.CreatedAt)
	m.SetField(member.FieldAddress, data.Address)
	m.SetField(member.FieldNickname, data.Nickname)
	m.SetField(member.FieldAvatar, data.Avatar)
	m.SetField(member.FieldIntro, data.Intro)
	m.SetField(member.FieldPublicKey, data.PublicKey)
	m.SetField(member.FieldNonce, data.Nonce)
	m.SetField(member.FieldShowNickname, data.ShowNickname)
}

func SaveMemberSyncData(ctx context.Context, b []byte, op Op, precall func(*MemberSync)) (err error) {
	data := new(MemberSync)
	err = jsoniter.Unmarshal(b, data)
	if err != nil {
		return
	}

	id := data.ID
	if id == "" {
		err = model.ErrSyncIDNotFound
		return
	}

	if op.Is(OpDelete | OpDeleteOne) {
		err = Database.Member.DeleteOneID(id).Exec(ctx)
	} else {
		creator := Database.Member.Create().SetID(id)
		if precall != nil {
			precall(data)
		}
		creator.Mutation().SetSyncData(data)
		columns := []string{
			"id",
			// member.FieldAddress,
		}
		err = creator.OnConflictColumns(columns...).UpdateNewValues().Exec(ctx)
	}

	return
}

type MessageSync struct {
	ID        string        `json:"id,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
	GroupID   string        `json:"group_id,omitempty"`
	MemberID  string        `json:"member_id,omitempty"`
	Content   []byte        `json:"content,omitempty"`
	ParentID  *string       `json:"parent_id,omitempty"`
	Owner     *model.Member `json:"owner,omitempty"`
}

func (m *MessageMutation) SyncData() (data *MessageSync) {
	ctx := context.Background()

	id, ok := m.ID()
	if !ok {
		log.Error("get Message sync data failed: id missing")
		return
	}

	if m.op.Is(OpDelete | OpDeleteOne) {
		data = &MessageSync{ID: id}
	} else {
		result, err := m.Client().Message.Get(ctx, id)
		if err != nil {
			log.Errorf("get Message sync data failed: %v", err)
			return
		}
		data = &MessageSync{
			ID:        id,
			CreatedAt: result.CreatedAt,
			GroupID:   result.GroupID,
			MemberID:  result.MemberID,
			Content:   result.Content,
			ParentID:  result.ParentID,
			Owner:     result.Owner,
		}
	}

	return data
}

func (m *MessageMutation) SetSyncData(data *MessageSync) {
	m.SetField(message.FieldCreatedAt, data.CreatedAt)
	m.SetField(message.FieldGroupID, data.GroupID)
	m.SetField(message.FieldMemberID, data.MemberID)
	m.SetField(message.FieldContent, data.Content)
	m.SetField(message.FieldParentID, data.ParentID)
	m.SetField(message.FieldOwner, data.Owner)
}

func SaveMessageSyncData(ctx context.Context, b []byte, op Op, precall func(*MessageSync)) (err error) {
	data := new(MessageSync)
	err = jsoniter.Unmarshal(b, data)
	if err != nil {
		return
	}

	id := data.ID
	if id == "" {
		err = model.ErrSyncIDNotFound
		return
	}

	if op.Is(OpDelete | OpDeleteOne) {
		err = Database.Message.DeleteOneID(id).Exec(ctx)
	} else {
		creator := Database.Message.Create().SetID(id)
		if precall != nil {
			precall(data)
		}
		creator.Mutation().SetSyncData(data)
		columns := []string{
			"id",
		}
		err = creator.OnConflictColumns(columns...).UpdateNewValues().Exec(ctx)
	}

	return
}
