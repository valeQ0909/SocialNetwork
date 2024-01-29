package service

type WatchService interface {
	// PostWatchCnt 计算当前帖子的浏览量
	PostWatchCnt(postId int64) (int64, error)

	// WatchAction 用户的浏览动作
	WatchAction(postId int64, userId int64) error
}
