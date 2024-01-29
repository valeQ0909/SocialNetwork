package models

import (
	"errors"
	"log"
)

// TableFollow 用户关系结构，对应用户关系表。
type TableFollow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
}

// TableName 设置Follow结构体对应数据库表名。
func (TableFollow) TableName() string {
	return "follows"
}

// AddFollow 添加关注数据
func AddFollow(userId int64, followerId int64) error {
	follow := TableFollow{
		UserId:     userId,
		FollowerId: followerId,
		Cancel:     0,
	}
	err := DB.Create(&follow).Error
	if err != nil {
		log.Println(err.Error())
		return errors.New("添加follow记录失败")
	}
	log.Println("插入了一条follow记录")
	return err
}

// QueryFollowInfo 给定当前用户id: curId，和目标用户id：userId，查询follow表中相应的记录。
func QueryFollowInfo(userId int64, curId int64) (TableFollow, error) {
	var result TableFollow
	err := DB.Model(TableFollow{}).Where(map[string]interface{}{
		"user_id":     userId,
		"follower_id": curId,
	}).First(&result).Error

	if err != nil {
		// 数据库中没有该数据
		if err.Error() == "record not found" {
			log.Println("QueryFollowInfo 数据库中没有记录")
			return TableFollow{}, err
		} else {
			log.Println(err.Error())
			return result, err
		}
	}

	return result, nil
}

// GetFollowingsId 给定用户id，查询该用户所有关注者的id。
func GetFollowingsId(userId int64) ([]int64, error) {
	var follwingsId []int64
	err := DB.Model(TableFollow{}).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Pluck("user_id", &follwingsId).Error
	if err != nil {
		log.Println("查询用户的关注用户id列表是发生错误：", err.Error())
		return nil, err
	}
	// 查询成功。
	return follwingsId, nil
}

// GetFollowersId 给定用户id，查询该用户所有的粉丝id
func GetFollowersId(userId int64) ([]int64, error) {
	var followersId []int64
	err := DB.Model(TableFollow{}).
		Where("user_id = ?", userId).
		Where("cancel = ?", 0).
		Pluck("follower_id", &followersId).Error

	if err != nil {
		log.Println("查询用户粉丝id时发生错误：", err.Error())
		return nil, err
	}

	return followersId, nil
}

// UpdateFollowAction 给定用户和目标用户的id，更新他们的关系为取消关注或再次关注。
func UpdateFollowAction(userId int64, targetId int64, cancel int32) error {
	log.Println("cancel：", cancel)
	// 更新失败，返回错误。
	err := DB.Model(TableFollow{}).
		Where("user_id = ?", userId).
		Where("follower_id = ?", targetId).
		Update("cancel", cancel).Error

	if err != nil {
		// 更新失败，打印错误日志。
		log.Println("更新关系为取消关注或者再次关注时发生错误： ", err.Error())
		return err
	}
	// 更新成功。
	return nil
}
