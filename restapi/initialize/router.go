package initialize

import (
	"Go-IMS/global"
	"Go-IMS/middleware"
	"Go-IMS/router"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
// 参数：
//		*gin.Engine：Gin引擎的指针
// 返回值：
//		无
func InitRouter() *gin.Engine {
	Router := gin.Default()
	// 注册pprof路由
	pprof.Register(Router)
	// 跨域中间件
	Router.Use(middleware.Cors())
	// 注册zap相关中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 注册i18n国际化中间件
	Router.Use(middleware.I18n())
	// 限制IP访问频率中间件
	Router.Use(middleware.Frequency())
	// 限制服务器并发请求数量，默认关闭限流
	if global.Settings.RaLiSw {
		Router.Use(middleware.RateLimit())
	}
	// 路由分组
	APIGroup := Router.Group("/api")
	{
		// v1版本路由
		APIv1 := APIGroup.Group("/v1")
		{
			router.Routerv1(APIv1)
		}
	}
	return Router
}
