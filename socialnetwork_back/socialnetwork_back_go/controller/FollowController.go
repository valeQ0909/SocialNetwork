package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/service"
	"strconv"
)

// FollowingResp 获取关注列表需要返回的结构。
type FollowingResp struct {
	Response
	UserList []service.FmtUser `json:"user_list,omitempty"`
}

// FollowerResp 获取粉丝列表需要返回的结构。
type FollowerResp struct {
	Response
	UserList []service.FmtUser `json:"user_list,omitempty"`
}

type FriendResp struct {
	Response
	FriendList []service.FmtFriend `json:"user_list"`
}

// FollowAction /follow/action/ - 关系操作
// 登录用户对其他用户进行关注或取消关注。
func FollowAction(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	followerId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	targetUserIdstr := c.PostForm("user_id")
	targetUserId, _ := strconv.ParseInt(targetUserIdstr, 10, 64)
	actionTypestr := c.PostForm("action_type")
	actionType, _ := strconv.ParseInt(actionTypestr, 10, 32)

	fsi := service.FollowServiceImpl{}
	err := fsi.FollowAction(targetUserId, followerId, int32(actionType))

	if err != nil {
		log.Println("关注或取消关注操作发生错误：", err)
		c.JSON(http.StatusOK, LikeActionResponse{
			StatusCode: 1,
			StatusMsg:  "关注或取消关注失败",
		})
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "关注或取消关注成功",
	})
}

// GetFollowingList /relation/follow/followinglist/ - 用户关注列表
// 登录用户关注的所有用户列表。
func GetFollowingList(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	fsi := service.FollowServiceImpl{}

	users, err := fsi.GetFollowingsList(userId)
	// 获取关注列表时出错。
	if err != nil {
		c.JSON(http.StatusOK, FollowingResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取关注列表时出错。",
			},
			UserList: nil,
		})
		return
	}
	// 成功获取到关注列表。
	log.Println("获取关注列表成功。")
	c.JSON(http.StatusOK, FollowingResp{
		UserList: users,
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
	})
}

// GetFollowerList /relation/follow/followerlist/ - 用户粉丝列表
// 所有关注登录用户的粉丝列表。
func GetFollowerList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	fsi := service.FollowServiceImpl{}
	followers, err := fsi.GetFollowersList(userId)
	if err != nil { // 获取粉丝列表时发生错误
		c.JSON(http.StatusOK, FollowerResp{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "获取粉丝列表时出错。",
			},
			UserList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, FollowerResp{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "OK",
		},
		UserList: followers,
	})
}

// GetFriendList /relation/friend/list/ - 用户好友列表
// 所有和用户互关的粉丝列表
//func GetFriendList(c *gin.Context) {
//	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
//	log.Println("controller 查询用户", userId, "的粉丝列表")
//
//	fsi := service.FollowServiceImpl{}
//
//	friendList, err := fsi.GetFriendList(userId)
//
//	if err != nil { // 获取粉丝列表时发生错误
//		log.Println("获取好友列表时发生错误：", err)
//		c.JSON(http.StatusOK, FriendResp{
//			Response: Response{
//				StatusCode: -1,
//				StatusMsg:  "获取粉丝列表时出错。",
//			},
//			FriendList: nil,
//		})
//		return
//	}
//
//	log.Println("打印好友列表：", friendList)
//	c.JSON(http.StatusOK, FriendResp{
//		Response: Response{
//			StatusCode: 0,
//			StatusMsg:  "OK",
//		},
//		FriendList: friendList,
//	})
//}
