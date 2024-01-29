package service

// LikeService 定义点赞状态和点赞数量
type LikeService interface {
	// IsFavourite 根据当前帖子id判断是否点赞了该帖子。
	IsFavourite(videoId int64, userId int64) (bool, error)
	//PostFavouriteCount 根据当前帖子id获取当前帖子点赞数量。
	PostFavouriteCount(videoId int64) (int64, error)

	//TotalFavourite 根据userId获取这个用户总共被点赞数量
	TotalFavourite(userId int64) (int64, error)
	//FavouritePostCount 根据userId获取这个用户点赞视频数量
	FavouritePostCount(userId int64) (int64, error)

	// FavouriteAction 当前用户对视频的点赞操作 ,并把这个行为更新到like表中。
	// 当前操作行为，1点赞，2取消点赞。
	FavouriteAction(userId int64, postId int64, actionType int32) error
	// GetLikeFmtPostList 获取当前用户的所有点赞视频，调用postService的方法
	GetLikeFmtPostList(userId int64, curId int64) ([]FmtPost, error)
}
