package router

import (
	"github.com/gin-gonic/gin"
	"socialnetwork_back_go/controller"
	"socialnetwork_back_go/utils"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./static")

	r.GET("/test/", controller.Test)
	apiRouter := r.Group("/socialnetwork")
	//用户注册
	apiRouter.POST("/user/register/", controller.Register)
	//用户登录
	apiRouter.POST("/user/login/", controller.Login)
	//用户信息
	apiRouter.GET("/user/info/", utils.Auth(), controller.UserInfo)
}
