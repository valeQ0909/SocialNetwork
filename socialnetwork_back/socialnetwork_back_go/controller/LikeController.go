package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/service"
	"strconv"
)

// LikeActionResponse 点赞或取消点赞操作的返回体
type LikeActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type LikeListResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	PostList   []service.FmtPost `json:"video_list"`
}

// LikeAction  /favorite/action/ - 赞操作
func LikeAction(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)
	actionTypestr := c.PostForm("action_type")
	actionType, _ := strconv.ParseInt(actionTypestr, 10, 32)

	fmt.Println("postId: ", postId)
	fmt.Println("actiontype: ", actionType)
	lsi := service.LikeServiceImpl{}
	err := lsi.FavouriteAction(userId, postId, int32(actionType))
	if err != nil {
		log.Println("service层方法LikeAction失败", err)
		c.JSON(http.StatusOK, LikeActionResponse{
			StatusCode: 1,
			StatusMsg:  "点赞或取消赞失败",
		})
	}

	c.JSON(http.StatusOK, LikeActionResponse{
		StatusCode: 0,
		StatusMsg:  "点赞或取消赞成功",
	})
}

// GetLikePostList /favorite/list/ - 喜欢列表
// 登录用户的所有点赞帖子。
func GetLikePostList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	curId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	if curId == 0 {
		c.JSON(http.StatusOK, LikeListResponse{
			StatusCode: 1,
			StatusMsg:  "用户未登录",
		})
	}

	lsi := service.LikeServiceImpl{}
	postList, err := lsi.GetLikeFmtPostList(userId, curId)
	if err != nil {
		log.Println("service层方法GetLikeVideoList失败", err)
		c.JSON(http.StatusOK, LikeListResponse{
			StatusCode: 1,
			StatusMsg:  "获取喜欢视频列表失败",
		})
	}
	log.Println("service层方法GetLikeVideoList成功")
	c.JSON(http.StatusOK, LikeListResponse{
		StatusCode: 0,
		StatusMsg:  "获取喜欢视频列表成功",
		PostList:   postList,
	})

}
