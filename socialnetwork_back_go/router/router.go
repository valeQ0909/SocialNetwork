package router

import (
	"github.com/gin-gonic/gin"
	"socialnetwork_back_go/controller"
	"socialnetwork_back_go/utils"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./static")
	r.Static("/media", "./media")

	r.GET("/test/", controller.Test)

	apiRouter := r.Group("/socialnetwork")

	/*
		用户
	*/
	//用户注册
	apiRouter.POST("/user/register/", controller.Register)
	//用户登录
	apiRouter.POST("/user/login/", controller.Login)
	//用户信息
	apiRouter.POST("/user/info/", utils.Auth(), controller.UserInfo)
	//根据token返回用户信息
	apiRouter.GET("/user/token/", utils.Auth(), controller.Token)
	//
	apiRouter.POST("/user/updateavatar/", utils.Auth(), controller.UpdateAvatar)
	//
	apiRouter.POST("/user/updatepersonalinfo/", utils.Auth(), controller.UpdatePersonalInfo)
	/*
		帖子
	*/
	// 获取帖子feed流
	apiRouter.POST("/post/feed/", utils.Auth(), controller.Feed)
	// 发布帖子
	apiRouter.POST("/post/sendpost/", utils.Auth(), controller.SendPost)
	// 帖子列表
	apiRouter.POST("/post/postlist/", utils.Auth(), controller.PostList)
	// 获取帖子详情
	apiRouter.POST("/post/postdetail/", utils.Auth(), controller.PostDetail)

	/*
		评论
	*/
	//评论操作
	apiRouter.POST("/comment/sendcomment/", utils.Auth(), controller.SendComment)
	//帖子评论列表
	apiRouter.POST("/comment/getcommentlist/", utils.Auth(), controller.CommentList)
	//评论回复列表
	apiRouter.POST("/comment/getreplylist/", utils.Auth(), controller.ReplyList)
}
