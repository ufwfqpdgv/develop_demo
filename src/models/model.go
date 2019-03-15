package models

import (
	"samh_common_lib/base"
)

type CommentListsRequest struct {
	base.SamhBaseRequest
	ContentType         int   `form:"content_type" json:"content_type" binding:"-"`
	FreshCommentNum     int   `form:"fresh_comment_num" json:"fresh_comment_num" binding:"required"`
	WonderfulCommentNum int   `form:"wonderful_comment_num" json:"wonderful_comment_num" binding:"required"`
	Ssid                int64 `form:"ssid" json:"ssid" binding:"required"`
	SsidType            int   `form:"ssid_type" json:"ssid_type" binding:"-"`
}

type CommentListsResponse struct {
	FreshWholeCommentArr     []*WholeComment `json:"fresh_whole_comment_arr,omitempty"`
	WonderfulWholeCommentArr []*WholeComment `json:"wonderful_whole_comment_arr,omitempty"`
	CommentSum               int64           `json:"comment_sum,omitempty"`
}

// 内部的
type InternalUserPrivilegeInfoRequest struct {
	base.SamhBaseRequest
}

type InternalUserPrivilegeInfoResponse struct {
	// *VipPrivilege
}

/* 请求其他服务的 */
type MicroUsersRequest struct {
}

type MicroUsersResponse struct {
	Code base.SamhResponseCode `json:"status"`
	Msg  string                `json:"message"`
	Data UserBase              `json:"data"`
}
