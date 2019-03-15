package api

import (
	. "models"
	"samh_common_lib/base"
	"samh_common_lib/utils"
	"samh_common_lib/utils/log"
	"strconv"
)

func fillCommentUserInfo(commentArr []*Comment) (wholeCommentArr []*WholeComment, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	var (
		err   error
		exist bool
	)

	wholeCommentArr = make([]*WholeComment, 0)
	for _, comment := range commentArr {
		rq2 := MicroUsersRequest{}
		rsp2 := &MicroUsersResponse{}
		addr := Config.Internal_server["user"].Url + "/v1/users/" + strconv.Itoa(int(comment.Useridentifier))
		timeOut := Config.Web.Http_request_timeout
		headers := map[string]string{
			"x-access-key": "erciyuan",
			"content-type": "application/json",
		}
		retCode = utils.HttpGet(addr, rq2, rsp2, timeOut, headers, 3)
		if retCode != base.SamhResponseCode_Succ {
			log.Error(base.NowFuncError())
			return
		}
		if rsp2.Code != base.SamhResponseCode_Succ {
			log.Error(base.NowFuncError())
			retCode = rsp2.Code
			return
		}
		userData := rsp2.Data
		wholeComment := &WholeComment{
			UserInfoItem: &UserInfo{
				Uid:    comment.Useridentifier,
				Uname:  userData.Uname,
				Ulevel: userData.Ulevel,
				IsVip:  userData.IsVip,
			},
			CommentItem: comment,
		}

		wholeComment.CommentItem.CreatetimeOut = comment.Createtime.Unix()

		exist, err = CommerceDB.Select("*").Table("supportcomment").
			Where("Useridentifier=? and CommentId=? and Status=1 and Appid=?",
				comment.Useridentifier, comment.Id, Config.Base_info.App_id).
			Exist(&Supportcomment{})
		if err != nil {
			log.Error(err)
			retCode = base.SamhResponseCode_ServerError
			return
		}
		if exist {
			wholeComment.CommentItem.IsSupport = 1
		}

		wholeCommentArr = append(wholeCommentArr, wholeComment)
	}

	return
}
