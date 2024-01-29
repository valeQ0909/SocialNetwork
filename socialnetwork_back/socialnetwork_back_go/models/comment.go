package models

import (
	"errors"
	"log"
	"time"
)

// TableComment
// 评论信息-数据库中的结构体-models层使用
type TableComment struct {
	Id              int64     `json:"id"`                //评论id
	CommenterId     int64     `json:"commenter_id"`      //评论用户id
	PostId          int64     `json:"post_id"`           //帖子id
	ParentCommentId int64     `json:"parent_comment_id"` //回复的评论id，如果没有则为0
	CommentText     string    `json:"comment_text"`      //评论内容
	PublishTime     time.Time `json:"publish_time"`      //评论发布的日期mm-dd
	Cancel          int32     `json:"cancel"`            //取消评论为1，发布评论为0
}

// TableName 修改表名映射
func (TableComment) TableName() string {
	return "comments"
}

// CommentCount
// 1、使用Post id 查询Comment数量
func CommentCount(postId int64) (int64, error) {
	var count int64
	// 数据库中查询评论数量
	err := DB.Model(TableComment{}).Where(map[string]interface{}{"post_id": postId, "cancel": 0}).Count(&count).Error
	if err != nil {
		return -1, errors.New("查询评论数量失败")
	}
	return count, nil
}

// GetCommentIdListByPostId 根据帖子id获取评论id列表
func GetCommentIdListByPostId(postId int64) ([]string, error) {
	var commentIdList []string
	err := DB.Model(TableComment{}).Select("id").Where("post_id = ?", postId).Find(&commentIdList).Error
	if err != nil {
		log.Println("CommentIdList:", err)
		return nil, err
	}
	return commentIdList, nil
}

// InsertComment
// 2、发表评论
func InsertComment(comment TableComment) error {
	// 数据库中插入一条评论信息
	err := DB.Model(TableComment{}).Create(&comment).Error
	if err != nil {
		return errors.New("发表评论存入数据失败")
	}
	return nil
}

// DeleteComment
// 3、删除评论，传入评论id
func DeleteComment(id int64) error {
	var commentInfo TableComment
	// 先查询是否有此评论
	result := DB.Model(TableComment{}).Where(map[string]interface{}{"id": id, "cancel": 0}).First(&commentInfo)
	if result.RowsAffected == 0 { // 查询到此评论数量为0则返回无此评论
		return errors.New("此评论不存在")
	}
	// 数据库中删除评论-更新评论状态为-1
	err := DB.Model(TableComment{}).Where("id = ?", id).Update("cancel", 1).Error
	if err != nil {
		return errors.New("删除评论失败")
	}
	return nil
}

// GetCommentList
// 4.根据帖子id查询所属评论全部列表信息
func GetCommentList(postId int64, parentCommentId int64) ([]TableComment, error) {
	// 数据库中查询评论信息list
	var commentList []TableComment
	log.Println("models comment postid: ", postId)
	result := DB.Model(&TableComment{}).Where(map[string]interface{}{"post_id": postId, "parent_comment_id": parentCommentId, "cancel": 0}).
		Order("publish_time desc").
		Find(&commentList)
	// 若此视频没有评论信息，返回空列表
	if result.RowsAffected == 0 {
		log.Println("此帖子无评论") // 函数返回提示无评论
		return nil, nil
	}
	// 若获取评论列表出错
	if result.Error != nil {
		log.Println(result.Error.Error())
		return commentList, errors.New("获取评论列表失败")
	}
	return commentList, nil
}
