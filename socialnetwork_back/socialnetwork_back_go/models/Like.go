package models

import (
	"errors"
	"log"
)

// TableLike 表的结构。
type TableLike struct {
	Id     int64 //自增主键
	UserId int64 //点赞用户id
	PostId int64 //帖子id
	Cancel int8  //是否点赞，0为点赞，1为取消赞
}

// TableName 修改表名映射
func (TableLike) TableName() string {
	return "likes"
}

// AddLike 添加点赞数据
func AddLike(userId int64, postId int64) error {
	like := TableLike{
		UserId: userId,
		PostId: postId,
		Cancel: 0,
	}
	err := DB.Create(&like).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("添加点赞数据失败")
	}
	log.Println("插入了一条like记录")
	return err
}

// QueryLikeInfo 根据usrId,postId查询具体的一条点赞信息
func QueryLikeInfo(userId int64, postId int64) (TableLike, error) {
	var result TableLike
	err := DB.Model(TableLike{}).Where(map[string]interface{}{
		"user_id": userId,
		"post_id": postId,
	}).First(&result).Error
	if err != nil {
		//数据库中没有该数据
		if err.Error() == "record not found" {
			log.Println("QueryLikeInfo 数据库中没有该数据")
			return TableLike{}, err
		} else {
			log.Println(err.Error())
			return result, errors.New("查询失败")
		}
	}
	return result, nil
}

// GetLikePostIdList 根据userId,查询该用户点赞过的全部postId
func GetLikePostIdList(userId int64) ([]int64, error) {
	var result []int64
	err := DB.Model(TableLike{}).Where(map[string]interface{}{
		"user_id": userId,
		"cancel":  0,
	}).Pluck("post_id", &result).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("该用户没有点赞过任何帖子")
			return result, nil
		} else {
			log.Println(err.Error())
			return result, errors.New("查询失败")
		}
	}
	return result, nil
}

// GetUserIdList 根据postId查询所有点赞过该帖子的用户userId
func GetUserIdList(postId int64) ([]int64, error) {
	var result []int64
	err := DB.Model(TableLike{}).Where(map[string]interface{}{
		"post_id": postId,
		"cancel":  0,
	}).Pluck("user_id", &result).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("该帖子没有被任何用户点赞过")
			return result, nil
		} else {
			log.Println(err.Error())
			return result, errors.New("查询失败")
		}
	}
	return result, nil
}

// UpdateLikeAction 根据userId,postId,actionType修改用户点赞的状态
func UpdateLikeAction(userId int64, postId int64, actionType int32) error {
	err := DB.Model(TableLike{}).Where(map[string]interface{}{
		"user_id": userId,
		"post_id": postId,
	}).Update("cancel", actionType).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("更新失败")
	}
	return nil
}
