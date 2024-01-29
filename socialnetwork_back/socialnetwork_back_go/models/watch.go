package models

import (
	"errors"
	"log"
)

// TableWatch 表的结构。
type TableWatch struct {
	Id     int64 //自增主键
	PostId int64 //帖子id
	UserId int64 //观看帖子的用户的id
}

// TableName 修改表名映射
func (TableWatch) TableName() string {
	return "watchs"
}

// AddWatch 添加点赞数据
func AddWatch(postId int64, userId int64) error {
	watch := TableWatch{
		PostId: postId,
		UserId: userId,
	}
	err := DB.Create(&watch).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("添加浏览记录失败")
	}
	log.Println("插入了一条watch记录")
	return err
}

// QueryWatchInfo 查询是否存在watch记录
func QueryWatchInfo(postId int64, userId int64) (TableWatch, error) {
	var result TableWatch
	err := DB.Model(TableWatch{}).Where(map[string]interface{}{
		"post_id": postId,
		"user_id": userId,
	}).First(&result).Error
	if err != nil {
		//数据库中没有该数据
		if err.Error() == "record not found" {
			log.Println("QueryWatchInfo 数据库中没有该数据")
			return TableWatch{}, err
		} else {
			log.Println(err.Error())
			return result, errors.New("查询失败")
		}
	}
	return result, nil
}

// GetWatchUserIdList 获取浏览过该帖子的用户的ID列表
func GetWatchUserIdList(postId int64) ([]int64, error) {
	var result []int64
	err := DB.Model(TableWatch{}).Where(map[string]interface{}{
		"post_id": postId,
	}).Pluck("user_id", &result).Error
	if err != nil {
		if err.Error() == "record not found" {
			log.Println("该帖子没有被任何用户浏览过")
			return result, nil
		} else {
			log.Println(err.Error())
			return result, errors.New("查询失败")
		}
	}
	return result, nil
}
