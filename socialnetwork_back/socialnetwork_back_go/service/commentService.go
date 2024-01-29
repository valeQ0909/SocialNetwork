package service

import "socialnetwork_back_go/models"

// FmtComment CommentService 接口定义
// 发表评论-使用的结构体-service层引用models层↑的Comment。
type FmtComment struct {
	models.TableComment
	Commenter      FmtUser `json:"commenter"`
	ReplyCount     int64   `json:"reply_count"`
	FmtPublishTime string  `json:"fmt_publish_time"`
}

type CommentService interface {
	// CountByPostId 根据postId获取帖子评论数量的接口   供其他service来调用
	CountByPostId(videoId int64) (int64, error)
	// SendComment 发送评论
	SendComment(comment models.TableComment) error
	// DeleteComment 删除评论
	DeleteComment(commentId int64) error
	// GetCommentList 获取帖子的所有评论列表
	GetCommentList(postId int64, userId int64) ([]FmtComment, error)
	// GetReplyList 获取帖子所有回复列表
	GetReplyList(postId int64, parentCommentId int64, userId int64) ([]FmtComment, error)
}
