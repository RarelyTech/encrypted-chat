// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/chatpuppy/puppychat/app/model"
)

// Pagination returns pagination query builder for GroupQuery.
func (gq *GroupQuery) Pagination(req model.PaginationReq) *GroupQuery {
	gq.Offset(req.GetOffset()).Limit(req.GetLimit())
	return gq
}

// PaginationItems returns pagination query builder for GroupQuery.
func (gq *GroupQuery) PaginationItemsX(req model.PaginationReq) any {
	return gq.Pagination(req).AllX(context.Background())
}

// PaginationResult returns pagination for GroupQuery.
func (gq *GroupQuery) PaginationResult(req model.PaginationReq) model.Pagination {
	ids := gq.Clone().Select("id").GroupBy("id").IntsX(context.Background())
	total := len(ids)
	return model.Pagination{
		Current: req.GetCurrent(),
		Pages:   req.GetPages(total),
		Total:   total,
	}
}

// Pagination returns pagination query builder for KeyQuery.
func (kq *KeyQuery) Pagination(req model.PaginationReq) *KeyQuery {
	kq.Offset(req.GetOffset()).Limit(req.GetLimit())
	return kq
}

// PaginationItems returns pagination query builder for KeyQuery.
func (kq *KeyQuery) PaginationItemsX(req model.PaginationReq) any {
	return kq.Pagination(req).AllX(context.Background())
}

// PaginationResult returns pagination for KeyQuery.
func (kq *KeyQuery) PaginationResult(req model.PaginationReq) model.Pagination {
	ids := kq.Clone().Select("id").GroupBy("id").IntsX(context.Background())
	total := len(ids)
	return model.Pagination{
		Current: req.GetCurrent(),
		Pages:   req.GetPages(total),
		Total:   total,
	}
}

// Pagination returns pagination query builder for MemberQuery.
func (mq *MemberQuery) Pagination(req model.PaginationReq) *MemberQuery {
	mq.Offset(req.GetOffset()).Limit(req.GetLimit())
	return mq
}

// PaginationItems returns pagination query builder for MemberQuery.
func (mq *MemberQuery) PaginationItemsX(req model.PaginationReq) any {
	return mq.Pagination(req).AllX(context.Background())
}

// PaginationResult returns pagination for MemberQuery.
func (mq *MemberQuery) PaginationResult(req model.PaginationReq) model.Pagination {
	ids := mq.Clone().Select("id").GroupBy("id").IntsX(context.Background())
	total := len(ids)
	return model.Pagination{
		Current: req.GetCurrent(),
		Pages:   req.GetPages(total),
		Total:   total,
	}
}

// Pagination returns pagination query builder for MessageQuery.
func (mq *MessageQuery) Pagination(req model.PaginationReq) *MessageQuery {
	mq.Offset(req.GetOffset()).Limit(req.GetLimit())
	return mq
}

// PaginationItems returns pagination query builder for MessageQuery.
func (mq *MessageQuery) PaginationItemsX(req model.PaginationReq) any {
	return mq.Pagination(req).AllX(context.Background())
}

// PaginationResult returns pagination for MessageQuery.
func (mq *MessageQuery) PaginationResult(req model.PaginationReq) model.Pagination {
	ids := mq.Clone().Select("id").GroupBy("id").IntsX(context.Background())
	total := len(ids)
	return model.Pagination{
		Current: req.GetCurrent(),
		Pages:   req.GetPages(total),
		Total:   total,
	}
}
