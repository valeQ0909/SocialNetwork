package service

// FmtFriend 格式化的好友列表
type FmtFriend struct {
	FmtUser
	Message string `json:"message,omitempty"`
	MsgType int64  `json:"msgType"`
}

type FollowService interface {
	// IsFollow 根据当前用户id和目标用户id来判断当前用户是否关注了目标用户
	IsFollow(userId int64, curId int64) (bool, error)
	// GetTotalFollowersCnt 根据用户id查询粉丝数量
	GetTotalFollowersCnt(userId int64) (int64, error)
	// GetTotalFollowingsCnt 根据用户id查询该用户的关注数量
	GetTotalFollowingsCnt(userId int64) (int64, error)

	// FollowAction 当前用户关注目标用户 操作
	FollowAction(userId int64, curId int64) error

	// GetFollowingsList 获取当前用户的关注列表
	GetFollowingsList(userId int64) ([]FmtUser, error)
	// GetFollowersList 获取当前用户的粉丝列表
	GetFollowersList(userId int64) ([]FmtUser, error)
}
