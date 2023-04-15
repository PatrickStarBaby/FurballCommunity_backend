package routers

import (
	"FurballCommunity_backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// 配置路由
	router := gin.Default()
	//设置默认路由当访问一个错误网站时返回
	router.NoRoute(controller.NotFound)
	// 提供静态文件服务
	// router.Static("/", "./public")
	router.Use(gin.Logger()) // 设置 gin 的日志级别为 Debug
	// v1
	v1 := router.Group("/v1")
	{

		user := v1.Group("/user")
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
		user.PUT("/updateUsername/:id", controller.UpdateUserName)
		user.PUT("/updatePassword/:id", controller.UpdatePassword)
		user.DELETE("/deleteUser/:id", controller.DeleteUser)
		user.GET("/getUserList", controller.GetUserList)
	}

	return router
}