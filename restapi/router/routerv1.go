package router

import (
	"Go-WMS/controller/v1/goods"
	"Go-WMS/controller/v1/other"
	"Go-WMS/controller/v1/user"
	"Go-WMS/middleware"
	"github.com/gin-gonic/gin"
)

// Routerv1 V1版本路由
// 参数：
//		r：Gin的路由分组的指针
// 返回值：
//		无
func Routerv1(r *gin.RouterGroup) {
	// 基础功能的路由
	baseRouter := r.Group("/base")
	{
		// 需要Token的接口
		baseRouterToken := baseRouter.Group("")
		baseRouterToken.Use(middleware.JWTAuth())
		// 需要Token和权限的接口
		{
			baseRouterTokenAdmin := baseRouterToken.Group("")
			baseRouterTokenAdmin.Use(middleware.IsAdminAuth())
			baseRouterTokenAdmin.GET("/health", other.ConGetSystemInfo) // 系统资源使用情况，CPU，内存，硬盘
		}
	}

	// 用户模块路由
	userRouter := r.Group("/user")
	{
		// 无需Token的接口
		userRouter.POST("/login", user.ConLogin) // 登录
		// 需要Token的接口
		userRouterToken := userRouter.Group("")
		userRouterToken.Use(middleware.JWTAuth())
		userRouterToken.DELETE("/logout", user.ConLogout)      // 登出
		userRouterToken.GET("/info/:id", user.ConGetUser)      // 查看用户信息
		userRouterToken.PUT("/info/:id", user.ConUpdateUser)   // 修改用户信息
		userRouterToken.PUT("/pwd/:id", user.ConUpdateUserPwd) // 修改用户密码
		// 需要Token和权限的接口
		{
			userRouterTokenAdmin := userRouterToken.Group("")
			userRouterTokenAdmin.Use(middleware.IsAdminAuth())
			userRouterTokenAdmin.POST("/info/:user_name", user.ConCreateUser) // 创建用户
			userRouterTokenAdmin.GET("/list", user.ConGetUserList)            // 用户列表
			userRouterTokenAdmin.DELETE("/info/:id", user.ConDeleteUser)      // 删除用户
		}
	}

	// 商品模块
	goodsRouter := r.Group("/goods")
	{
		// 需要Token的接口
		goodsRouterToken := goodsRouter.Group("")
		goodsRouterToken.Use(middleware.JWTAuth())
		goodsRouterToken.POST("/category", goods.ConCreateGoodsCategory)            // 创建商品分类
		goodsRouterToken.GET("/category/list/:type", goods.ConGetGoodsCategoryList) // 商品分类列表
	}
}
