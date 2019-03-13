package models

import (
	"samh_common_lib/base"
)

type PrivilegeInfoRequest struct {
	base.SamhBaseRequest
}

type PrivilegeInfoResponse struct {
	// VipPrivilegeArr []*VipPrivilege `json:"vip_privilege_arr"`
}

// 内部的
type InternalUserPrivilegeInfoRequest struct {
	base.SamhBaseRequest
}

type InternalUserPrivilegeInfoResponse struct {
	// *VipPrivilege
}

/* 请求其他服务的 */
type OperationActivityRequest struct {
	Uid        int64  `form:"uid" json:"uid"`
	DeviceId   string `form:"udid" json:"udid"`
	ActivityId int64  `form:"activity_id" json:"activity_id"`
}

type OperationActivityResponse struct {
	Code base.SamhResponseCode `json:"status"`
	Msg  string                `json:"msg"`
	Data DataTemp1             `json:"data"`
}
