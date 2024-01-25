package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/service"
	"strconv"
	"time"
)

type PostResponse struct {
	Response
	Post service.FmtPost `json:"post"`
}
type PostsListResponse struct {
	Response
	PostList []service.FmtPost `json:"post_list"`
}

// Feed 获取首页的帖子
func Feed(c *gin.Context) {
	tokenId, _ := c.Get("tokenId") //返回的这个any类型我不是很懂，留着以后再学
	tokenIdstr := tokenId.(string) // 这个any必须转换未string，我也不是很懂，留着以后再学
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)
	lastTime := time.Now()

	psi := service.PostServiceImpl{}
	feed, err := psi.Feed(lastTime, userId)
	if err != nil {
		log.Printf("方法postService.Feed(lastTime, userId) 失败：%v", err)
		c.JSON(http.StatusOK, PostsListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频流失败"},
		})
		return
	}

	c.JSON(http.StatusOK, PostsListResponse{
		Response: Response{StatusCode: 0},
		PostList: feed,
	})
}

// SendPost 发送帖子
func SendPost(c *gin.Context) {
	tokenId, _ := c.Get("tokenId") //返回的这个any类型我不是很懂，留着以后再学
	tokenIdstr := tokenId.(string) // 这个any必须转换未string，我也不是很懂，留着以后再学
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	category := c.PostForm("category")
	postHtml := c.PostForm("post_html")
	postText := c.PostForm("post_text")

	log.Println("postHtml: ", postHtml)
	log.Println("postText: ", postText)
	psi := service.PostServiceImpl{}
	err := psi.SendPost(userId, category, postHtml, postText)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "后端data入库失败",
		})
		return
	}
	log.Println("成功发布帖子")
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "发布帖子成功",
	})

}

// PostList 获取帖子列表  /post/postlist/
func PostList(c *gin.Context) {
	tokenId, _ := c.Get("tokenId") //返回的这个any类型我不是很懂，留着以后再学
	tokenIdstr := tokenId.(string) // 这个any必须转换未string，我也不是很懂，留着以后再学
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	username := c.PostForm("username")

	psi := service.PostServiceImpl{}
	usi := service.UserServiceImpl{}
	user, _ := usi.GetUserByUsername(username)
	// 获取用户所发布视频的列表
	postList, err := psi.GetPostList(user.Id, userId)
	// 获取用户所发布视频的列表 -- 失败
	if err != nil {
		c.JSON(http.StatusOK, PostsListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取用户所发布视频的列表失败"},
		})
		return
	}
	// 获取用户所发布视频的列表 -- 成功
	c.JSON(http.StatusOK, PostsListResponse{
		Response: Response{StatusCode: 0, StatusMsg: "获取用户所发布视频的列表成功"},
		PostList: postList,
	})
}

// PostDetail 帖子详情
func PostDetail(c *gin.Context) {
	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)

	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	psi := service.PostServiceImpl{}
	post, err := psi.GetPostByPostId(postId, userId)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "查询失败",
		})
		return
	}
	log.Println("成功查询到帖子")
	c.JSON(http.StatusOK, PostResponse{
		Response: Response{StatusCode: 0, StatusMsg: "查询帖子成功"},
		Post:     post,
	})
}
