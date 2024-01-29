package service

import (
	"log"
	"socialnetwork_back_go/models"
)

type LikeServiceImpl struct {
}

// IsFavourite 根据当前帖子id判断是否点赞了该帖子。
// 思路：查询用户所有点赞视频的id列表，然后查看里面是否有当前帖子的id
func (lsi *LikeServiceImpl) IsFavourite(postId int64, userId int64) (bool, error) {
	//获取用户所有喜欢的帖子的id
	isLike := false
	likePostList, _ := models.GetLikePostIdList(userId)

	for i := 0; i < len(likePostList); i++ {
		if likePostList[i] == postId {
			isLike = true
		}
	}

	return isLike, nil
}

// PostFavouriteCount 根据当前帖子id获取当前帖子点赞数量。
// 思路：查询所有点赞过当前帖子的用户，然后计算数量
func (lsi *LikeServiceImpl) PostFavouriteCount(postId int64) (int64, error) {
	userIds, _ := models.GetUserIdList(postId)
	return int64(len(userIds)), nil
}

// TotalFavourite 根据userId获取这个用户总共被点赞数量
// 思路：先获取该用户所有的帖子，然后累加每个帖子的点赞数量
func (lsi *LikeServiceImpl) TotalFavourite(userId int64) (int64, error) {
	postIdList, _ := models.GetPostIdList(userId)
	num := len(postIdList)
	sum := int64(0)
	var userIds []int64
	for i := 0; i < num; i++ {
		userIds, _ = models.GetUserIdList(postIdList[i])
		sum = sum + int64(len(userIds))
	}
	return sum, nil
}

// FavouriteAction 当前用户对视频的点赞操作 ,并把这个行为更新到like表中。
// 当前操作行为，0点赞，1取消点赞。
func (lsi *LikeServiceImpl) FavouriteAction(userId int64, postId int64, actionType int32) error {
	var err error

	if actionType == 0 { // 点赞
		// 先查有没有这条点赞记录
		_, err = models.QueryLikeInfo(userId, postId)
		if err != nil {
			if err.Error() == "record not found" { // 没有该记录
				log.Println("数据库中没有该数据")
				// 没有点赞记录，增加该点赞记录
				err = models.AddLike(userId, postId)
				if err != nil {
					log.Println("插入点赞数据失败")
				}
			} else { // 发生了其他错误
				log.Println(err.Error())
			}
		} else { // 表中存在记录
			err = models.UpdateLikeAction(userId, postId, actionType)
			if err != nil {
				log.Println("方法FavouriteAction失败")
			}
		}
	} else if actionType == 1 { // 取消点赞
		err = models.UpdateLikeAction(userId, postId, actionType)
		if err != nil {
			log.Println("方法FavouriteAction失败")
		}
	}

	return err
}

// GetLikeFmtPostList 获取当前用户的所有点赞的帖子并加工成FmtPost，调用postService的方法
// 施工ing
func (lsi *LikeServiceImpl) GetLikeFmtPostList(userId int64, curId int64) ([]FmtPost, error) {

	psi := PostServiceImpl{}
	var fmtPosts []FmtPost
	fmtPost := FmtPost{}

	likePostIdList, _ := models.GetLikePostIdList(userId)
	num := len(likePostIdList)
	for i := 0; i < num; i++ {
		fmtPost, _ = psi.GetFmtPostByPostId(likePostIdList[i], curId)
		fmtPosts = append(fmtPosts, fmtPost)
	}

	return fmtPosts, nil
}
