package router

import (
	"server/api"
	"server/custom_middleware"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(custom_middleware.ServerHeader)
	// router.Use(custom_middleware.Log())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	/* 对外接口 */
	router.GET("api/v1/vip/privilege_info", api.PrivilegeInfoApi)
	/* 内部接口 */
	router.POST("internal/api/v1/vip/user_privilege_info", api.InternalUserPrivilegeInfoApi)

	return router
}
