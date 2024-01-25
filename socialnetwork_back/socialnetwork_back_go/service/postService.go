package service

import (
	"socialnetwork_back_go/models"
	"time"
)

type FmtPost struct {
	models.TablePost
	Author         FmtUser `json:"author,omitempty"`
	FavoriteCount  int64   `json:"favorite_count,omitempty"`
	CommentCount   int64   `json:"comment_count,omitempty"`
	IsFavorite     bool    `json:"is_favorite,omitempty"`
	FmtPublishTime string  `json:"fmt_publish_time,omitempty"`
}

type PostService interface {
	// Feed 通过传入时间戳，当前用户的id返回对应的帖子切片数组，以及帖子数组中最早的发布时间
	Feed(lastTime time.Time, userId int64) ([]FmtPost, time.Time, error)
	// GetPost 传入帖子id获取对应的视频对象
	GetPost(postId int64, userId int64) (FmtPost, error)
	// SendPost 发送帖子
	SendPost(authorId int64, category string, postHtml string, postText string) error
	// GetPostList 获取帖子列表
	GetPostList(userId int64, curId int64) ([]FmtPost, error)
}
