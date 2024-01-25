package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/service"
	"strconv"
)

// CommentListResponse
// 评论列表返回参数
type CommentListResponse struct {
	Response
	CommentList []service.FmtComment `json:"comment_list,omitempty"`
}

// SendCommentResponse
// 发表评论返回参数
type SendCommentResponse struct {
	Response
}

// SendComment /comment/sendcomment/ - 评论操作
// 登录用户对视频进行评论。
func SendComment(c *gin.Context) {
	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)
	parentIdstr := c.PostForm("parent_comment_id")
	parentId, _ := strconv.ParseInt(parentIdstr, 10, 64)
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)
	comment_text := c.PostForm("comment_text")

	csi := new(service.CommentServiceImpl)

	err := csi.SendComment(userId, postId, parentId, comment_text)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "发表评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, SendCommentResponse{
		Response: Response{StatusCode: 0, StatusMsg: "发表评论成功"},
	})
}

// CommentList /comment/getcommentlist/ - 视频评论列表
// 查看帖子的所有评论，按发布时间倒序。
func CommentList(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)
	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)
	log.Println("postId: ", postId)
	csi := new(service.CommentServiceImpl)

	list, err := csi.GetCommentList(postId, userId)
	log.Println("comment list: ", list)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "GetCommentsList error",
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "GetCommentsList success"},
		CommentList: list,
	})
}

// ReplyList /comment/getreplylist/ - 视频评论列表
// 查看评论的所有回复，按发布时间倒序。
func ReplyList(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)
	postIdstr := c.PostForm("post_id")
	postId, _ := strconv.ParseInt(postIdstr, 10, 64)
	parentCommentIdstr := c.PostForm("parent_comment_id")
	parentCommentId, _ := strconv.ParseInt(parentCommentIdstr, 10, 64)
	csi := new(service.CommentServiceImpl)

	list, err := csi.GetReplyList(postId, parentCommentId, userId)
	log.Println("reply list: ", list)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "GetReplyList error",
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "GetGetReplyList success"},
		CommentList: list,
	})
}
