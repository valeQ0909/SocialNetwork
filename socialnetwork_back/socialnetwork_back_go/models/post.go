package models

import (
	"log"
	"socialnetwork_back_go/config"
	"time"
)

// TablePost 映射字段名
type TablePost struct {
	Id          int64     `json:"id"`
	AuthorId    int64     `json:"author_id"`
	Category    string    `json:"category"`
	PostHtml    string    `json:"post_html"`
	PostText    string    `json:"post_text"`
	PublishTime time.Time `json:"publish_time"`
}

// TableName 表名映射
func (u TablePost) TableName() string {
	return "posts"
}

// GetPostsByLastTime
// 依据一个时间，来获取这个时间之前的一些帖子
func GetPostsByLastTime(lastTime time.Time) ([]TablePost, error) {
	posts := make([]TablePost, config.PostCount)
	result := DB.Where("publish_time<?", lastTime).Order("publish_time desc").Limit(config.PostCount).Find(&posts)
	if result.Error != nil {
		log.Println("查询posts有错误", result.Error)
		return posts, result.Error
	}
	return posts, nil
}

// GetPostsByAuthorId
// 根据作者的id来查询对应数据库数据，返回Post切片
func GetPostsByAuthorId(authorId int64) ([]TablePost, error) {
	var post []TablePost
	result := DB.Where(&TablePost{AuthorId: authorId}).Find(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}

// GetPostByPostId
// 依据VideoId来获得视频信息
func GetPostByPostId(videoId int64) (TablePost, error) {
	var post TablePost
	post.Id = videoId
	result := DB.First(&post)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

// InsertTablePost 创建一条帖子
func InsertTablePost(authorId int64, category string, postHtml string, postText string) error {
	newPost := TablePost{
		AuthorId:    authorId,
		Category:    category,
		PostHtml:    postHtml,
		PostText:    postText,
		PublishTime: time.Now(),
	}
	result := DB.Create(&newPost)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetPostIdList
// 通过作者id来查询发布的帖子id切片集合
func GetPostIdList(authorId int64) ([]int64, error) {
	var id []int64
	//通过pluck来获得单独的切片
	result := DB.Model(&TablePost{}).Where("author_id", authorId).Pluck("id", &id)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return id, nil
}
