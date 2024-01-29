package service

import (
	"log"
	"socialnetwork_back_go/models"
)

// FollowServiceImpl 该结构体继承FollowService接口。
type FollowServiceImpl struct {
}

// IsFollow 根据当前用户id和目标用户id来判断当前用户是否关注了目标用户
// 思路：查询用户所有关注的用户列表，然后查看里面是否有目标用户的id
func (fsi *FollowServiceImpl) IsFollow(userId int64, curId int64) (bool, error) {
	//获取所有该用户关注的用户的id
	isFollow := false
	followUserList, _ := models.GetFollowingsId(curId)

	for i := 0; i < len(followUserList); i++ {
		if followUserList[i] == userId {
			isFollow = true
		}
	}

	return isFollow, nil
}

// GetTotalFollowersCnt 根据用户id查询粉丝数量
// 思路：查询该用户的所有粉丝，然后计算数量
func (fsi *FollowServiceImpl) GetTotalFollowersCnt(userId int64) (int64, error) {
	followersList, _ := models.GetFollowersId(userId)
	return int64(len(followersList)), nil
}

// GetTotalFollowingsCnt 根据用户id查询该用户的关注数量
// 施工ing
func (fsi *FollowServiceImpl) GetTotalFollowingsCnt(userId int64) (int64, error) {
	followingsList, _ := models.GetFollowingsId(userId)
	return int64(len(followingsList)), nil
}

// FollowAction 当前用户关注目标用户 操作
func (fsi *FollowServiceImpl) FollowAction(userId int64, followerId int64, actionType int32) error {
	var err error

	if actionType == 0 { // 关注
		// 先查有没有这条关注记录
		_, err = models.QueryFollowInfo(userId, followerId)
		if err != nil {
			if err.Error() == "record not found" { // 没有该记录
				log.Println("数据库中没有该数据")
				// 没有点赞记录，增加该关注记录
				err = models.AddFollow(userId, followerId)
				if err != nil {
					log.Println("插入关注数据失败")
				}
			} else { // 发生了其他错误
				log.Println(err.Error())
			}
		} else { // 表中存在记录
			err = models.UpdateFollowAction(userId, followerId, actionType)
			if err != nil {
				log.Println("方法FollowAction失败")
			}
		}
	} else if actionType == 1 { // 取消点赞
		err = models.UpdateFollowAction(userId, followerId, actionType)
		if err != nil {
			log.Println("方法FavouriteAction失败")
		}
	}

	return err
}

// GetFollowingsList 获取当前用户的关注列表
func (fsi *FollowServiceImpl) GetFollowingsList(userId int64) ([]FmtUser, error) {

	return []FmtUser{}, nil
}

// GetFollowersList 获取当前用户的粉丝列表
func (fsi *FollowServiceImpl) GetFollowersList(userId int64) ([]FmtUser, error) {

	return []FmtUser{}, nil
}
