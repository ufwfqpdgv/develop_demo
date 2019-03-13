package api

import (
	. "models"
	process_api "process/api"
	"samh_common_lib/base"
	"samh_common_lib/utils/log"

	"github.com/gin-gonic/gin"
)

func PrivilegeInfoApi(c *gin.Context) {
	defer base.RecoverFunc(c)

	rq := &PrivilegeInfoRequest{}
	err := c.ShouldBind(rq)
	log.Infof(base.NowFunc()+"Request:%+v", *rq)
	if err == nil {
		rsp, retCode := process_api.PrivilegeInfoApi(rq)
		log.Infof(base.NowFunc()+"Response:%v,%+v", retCode, *rsp)
		base.SendResponse(c, retCode, rsp)
	} else {
		log.Warn(base.NowFuncError(), err)
		base.SendResponse(c, base.SamhResponseCode_Param_Less, nil)
	}
}

// 内部接口
func InternalUserPrivilegeInfoApi(c *gin.Context) {
	defer base.RecoverFunc(c)

	rq := &InternalUserPrivilegeInfoRequest{}
	err := c.ShouldBind(rq)
	log.Infof(base.NowFunc()+"Request:%+v", *rq)
	if err == nil {
		rsp, retCode := process_api.InternalUserPrivilegeInfoApi(rq)
		log.Infof(base.NowFunc()+"Response:%v,%+v", retCode, *rsp)
		base.SendResponse(c, retCode, rsp)
	} else {
		log.Warn(base.NowFuncError(), err)
		base.SendResponse(c, base.SamhResponseCode_Param_Less, nil)
	}
}
