package service

import (
	"log"
	"socialnetwork_back_go/models"
)

type WatchServiceImpl struct {
}

// PostWatchCnt 计算当前帖子的浏览量
func (wsi *WatchServiceImpl) PostWatchCnt(postId int64) (int64, error) {
	watchIds, _ := models.GetWatchUserIdList(postId)
	return int64(len(watchIds)), nil
}

// WatchAction 用户的浏览动作
func (wsi *WatchServiceImpl) WatchAction(postId int64, userId int64) error {
	var err error
	// 先查有没有这条浏览记录
	_, err = models.QueryWatchInfo(postId, userId)
	if err != nil {
		if err.Error() == "record not found" { // 没有该记录
			log.Println("数据库中没有该数据")
			// 没有浏览记录，增加该浏览记录
			err = models.AddWatch(postId, userId)
			if err != nil {
				log.Println("插入浏览数据失败")
			} else {
				return nil //成功插入返回nil
			}
		} else { // 发生了其他错误
			log.Println(err.Error())
		}
	}
	//如果有该记录返回nil， 发生其他错误也直接返回
	return err
}
