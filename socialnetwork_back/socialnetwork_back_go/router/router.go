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
		浏览
	*/
	apiRouter.POST("/watch/action/", utils.Auth(), controller.WatchAction)
	/*
		点赞
	*/
	//赞操作
	apiRouter.POST("/favorite/action/", utils.Auth(), controller.LikeAction)
	//喜欢列表
	apiRouter.GET("/favorite/list/", utils.Auth(), controller.GetLikePostList)

	/*
		关注
	*/
	apiRouter.POST("/follow/action/", utils.Auth(), controller.FollowAction)
	//用户关注列表
	apiRouter.GET("/follow/followinglist/", utils.Auth(), controller.GetFollowingList)
	//用户粉丝列表
	apiRouter.GET("/follow/followerlist/", utils.Auth(), controller.GetFollowerList)
	//用户好友列表
	// 施工ing
	apiRouter.GET("/friend/list/", utils.Auth())

	/*
		评论
	*/
	//评论操作
	apiRouter.POST("/comment/sendcomment/", utils.Auth(), controller.SendComment)
	//帖子评论列表
	apiRouter.POST("/comment/getcommentlist/", utils.Auth(), controller.CommentList)
	//评论回复列表
	apiRouter.POST("/comment/getreplylist/", utils.Auth(), controller.ReplyList)

	/*
		vale的AI小屋
	*/
	//图片识别
	apiRouter.POST("/aihome/", utils.Auth(), controller.ImagePredict)
}
