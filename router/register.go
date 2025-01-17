package router

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/middleware"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/MQEnergy/gin-framework/router/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	gin.SetMode(global.Cfg.Server.Mode)
	router := gin.New()
	// [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	router.SetTrustedProxies(config.AllowIpList)
	// header add X-Request-Id
	router.Use(requestid.New())
	router.Use(middleware.RequestIdAuth())
	// 404 not found
	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		response.NotFoundException(ctx, fmt.Sprintf("%s %s not found", method, path))
	})

	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		commonGroup = router.Group("/", publicMiddleware...)
		authGroup   = router.Group("/", append(publicMiddleware, middleware.LoginAuth(), middleware.CasbinAuth())...)
	)
	// 公用组
	routes.InitCommonGroup(commonGroup)
	// 后台组
	routes.InitBackendGroup(authGroup)
	// 前台组
	routes.InitFrontendGroup(authGroup)
	// 赋给全局
	global.Router = router
	return router
}
