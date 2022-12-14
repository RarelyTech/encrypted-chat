// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"reflect"
)

// EntitySetAttributes call set field value for the given entity
func EntitySetAttributes[T, O any](client *T, entity *O, data any) *T {
	nd := reflect.TypeOf(data)
	ptr := nd.Kind() == reflect.Pointer
	if !ptr {
		panic("data need pointer")
	}

	nd = nd.Elem()
	ndv := reflect.ValueOf(data).Elem()

	od := reflect.TypeOf(entity).Elem()

	uper := reflect.ValueOf(client)
	for i := 0; i < ndv.NumField(); i++ {
		nf := nd.Field(i)
		nfp := nf.Type.Kind() == reflect.Pointer
		if !nfp {
			continue
		}
		fn := nf.Name
		if fn == "ID" {
			continue
		}
		of, ok := od.FieldByName(fn)
		if !ok {
			continue
		}
		ofp := of.Type.Kind() == reflect.Pointer

		v := ndv.FieldByName(fn)
		if !v.IsZero() {
			method := uper.MethodByName(fmt.Sprintf("Set%s", fn))
			if !ofp {
				v = v.Elem()
			}
			if !method.IsZero() {
				method.Call([]reflect.Value{v})
			}
		}
	}
	return client
}

// ModifyOne returns an update with pointer struct builder for Group.
func (c *GroupClient) ModifyOne(old *Group, data any) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(old))
	up := &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
	return EntitySetAttributes[GroupUpdateOne, Group](up, old, data)
}

// ModifyOne returns an update with pointer struct builder for GroupMember.
func (c *GroupMemberClient) ModifyOne(old *GroupMember, data any) *GroupMemberUpdateOne {
	mutation := newGroupMemberMutation(c.config, OpUpdateOne, withGroupMember(old))
	up := &GroupMemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
	return EntitySetAttributes[GroupMemberUpdateOne, GroupMember](up, old, data)
}

// ModifyOne returns an update with pointer struct builder for Key.
func (c *KeyClient) ModifyOne(old *Key, data any) *KeyUpdateOne {
	mutation := newKeyMutation(c.config, OpUpdateOne, withKey(old))
	up := &KeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
	return EntitySetAttributes[KeyUpdateOne, Key](up, old, data)
}

// ModifyOne returns an update with pointer struct builder for Member.
func (c *MemberClient) ModifyOne(old *Member, data any) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(old))
	up := &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
	return EntitySetAttributes[MemberUpdateOne, Member](up, old, data)
}

// ModifyOne returns an update with pointer struct builder for Message.
func (c *MessageClient) ModifyOne(old *Message, data any) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(old))
	up := &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
	return EntitySetAttributes[MessageUpdateOne, Message](up, old, data)
}
