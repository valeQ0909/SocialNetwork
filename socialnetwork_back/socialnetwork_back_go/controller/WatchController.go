package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/service"
	"strconv"
)

// WatchAction  /watch/action/ - 赞操作
func WatchAction(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)

	wsi := service.WatchServiceImpl{}
	err := wsi.WatchAction(postId, userId)
	if err != nil {
		log.Println("service层方法WatchAction失败", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "增加浏览记录失败",
		})
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "浏览帖子成功",
	})
}
