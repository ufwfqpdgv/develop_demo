package api

import (
	"fmt"

	. "models"
	"samh_common_lib/base"
	"samh_common_lib/utils/log"
	// "github.com/pkg/errors"
)

func CommentListsApi(rq *CommentListsRequest) (rsp *CommentListsResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &CommentListsResponse{}
	var (
		err error
	)
	if rq.FreshCommentNum > 10 || rq.WonderfulCommentNum > 10 {
		retCode = base.SamhResponseCode_Param_Invalid
		return
	}

	rsp.CommentSum, err = CommerceDB.Table("comment").
		Where(fmt.Sprintf("SSID=? and SSIDType=? and AppId in(%v) and RelateId!='-1' ", Config.Comic.Appids_str),
			rq.Ssid, rq.SsidType).
		Count(&Comment{})
	if err != nil {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}
	if rsp.CommentSum == 0 {
		return
	}

	freshCommentArr := make([]*Comment, 0)
	err = CommerceDB.Select("*").Table("comment").
		Where(fmt.Sprintf("SSID=? and SSIDType=? and AppId in(%v) and RelateId!='-1' and Fatherid=0 and Status=1 ", Config.Comic.Appids_str),
			rq.Ssid, rq.SsidType).
		OrderBy("Createtime desc").
		Limit(rq.FreshCommentNum).
		Find(&freshCommentArr)
	if err != nil {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}
	if len(freshCommentArr) == 0 {
		return
	}

	wonderfulCommentArr := make([]*Comment, 0)
	switch rq.ContentType {
	case 0, 3:
		err = CommerceDB.Select("*").Table("comment").
			Where(fmt.Sprintf("SSID=? and SSIDType=? and AppId in(%v) and Fatherid=0 and Status=1 and (Istop=1 or IsElite=1 or (Supportcount>5 and Contentlength>5))  ", Config.Comic.Appids_str),
				rq.Ssid, rq.SsidType).
			OrderBy("IsTop desc,IsElite desc,SupportCount desc,Createtime desc").
			Limit(rq.WonderfulCommentNum).
			Find(&wonderfulCommentArr)
	case 1:
		err = CommerceDB.Select("*").Table("comment").
			Where(fmt.Sprintf("SSID=? and SSIDType=? and AppId in(%v) and Fatherid=0 and Status=1 and (Istop=1 or IsElite=1)  ", Config.Comic.Appids_str),
				rq.Ssid, rq.SsidType).
			OrderBy("IsTop desc,IsElite desc,SupportCount desc,Createtime desc").
			Limit(rq.WonderfulCommentNum).
			Find(&wonderfulCommentArr)
	case 2:
		err = CommerceDB.Select("*").Table("comment").
			Where(fmt.Sprintf("SSID=? and SSIDType=? and AppId in(%v) and Fatherid=0 and Status=1 and Supportcount>5 and Contentlength>5 and Istop=0 and IsElite=0)  ", Config.Comic.Appids_str),
				rq.Ssid, rq.SsidType).
			OrderBy("IsTop desc,IsElite desc,SupportCount desc,Createtime desc").
			Limit(rq.WonderfulCommentNum).
			Find(&wonderfulCommentArr)
	default:
		retCode = base.SamhResponseCode_Param_Invalid
		return
	}
	if err != nil {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}

	rsp.FreshWholeCommentArr, retCode = fillCommentUserInfo(freshCommentArr)
	if retCode != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		return
	}
	rsp.WonderfulWholeCommentArr, retCode = fillCommentUserInfo(wonderfulCommentArr)
	if retCode != base.SamhResponseCode_Succ {
		log.Error(base.NowFuncError())
		return
	}

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
