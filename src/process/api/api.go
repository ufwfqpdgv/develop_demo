package api

import (
	. "models"
	"samh_common_lib/base"
	"samh_common_lib/utils/log"
)

func PrivilegeInfoApi(rq *PrivilegeInfoRequest) (rsp *PrivilegeInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &PrivilegeInfoResponse{}
	// var (
	// err     error
	// exist   bool
	// nowTime int64 = time.Now().Unix()
	// )

	// rsp.VipPrivilegeArr = make([]*VipPrivilege, 0)
	// err := VipDB.Select("*").Table("vip_privilege").Where("status=?", base.SamhDataStatusCode_Normal).
	// Find(&rsp.VipPrivilegeArr)
	// if err != nil {
	// log.Error(err)
	// retCode = base.SamhResponseCode_ServerError
	// return
	// }
	// if len(rsp.VipPrivilegeArr) == 0 {
	// retCode = base.SamhResponseCode_Data_NotExist
	// return
	// }

	return
}

// 内部接口
func InternalUserPrivilegeInfoApi(rq *InternalUserPrivilegeInfoRequest) (rsp *InternalUserPrivilegeInfoResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &InternalUserPrivilegeInfoResponse{}
	// var (
	// err     error
	// exist   bool
	// nowTime int64 = time.Now().Unix()
	// )

	// userVip := &UserVip{}
	// exist, err = VipDB.Select("*").Table("user_vip").
	// Where("uid=? and status=?",
	// rq.Uid, base.SamhDataStatusCode_Normal).
	// Get(userVip)
	// if err != nil {
	// log.Error(err)
	// retCode = base.SamhResponseCode_ServerError
	// return
	// }
	// if !exist {
	// retCode = base.SamhResponseCode_Data_NotExist
	// return
	// }
	// rq2 := PayOrderRequest{
	// Uid:       rq.Uid,
	// DeviceId:  rq.DeviceId,
	// AppId:     Config.Base_info.App_id,
	// ProductId: int(vipCombo.ProductId),
	// PayType:   rq.PayType,
	// NotifyUrl: Config.Internal_server["ivip"].Url + "recharge_callback",
	// Supply:    " ",
	// }
	// rsp2 := &PayOrderResponse{}
	// retCode = utils.HttpPost(Config.Internal_server["pay"].Url+"prepay",
	// rq2, rsp2, Config.Web.Http_request_timeout)
	// if retCode != base.SamhResponseCode_Succ {
	// log.Error(base.NowFuncError())
	// return
	// }
	// if rsp2.Code != base.SamhResponseCode_Succ {
	// log.Error(base.NowFuncError())
	// retCode = rsp2.Code
	// return
	// }
	// rsp.Extra = rsp2.Data

	return
}
