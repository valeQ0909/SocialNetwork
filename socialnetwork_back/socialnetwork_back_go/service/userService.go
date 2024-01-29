package service

import "socialnetwork_back_go/models"

// FmtUser 最终封装后,controller返回的FmtUser结构体
type FmtUser struct {
	Id           int64  `json:"id"`
	UserName     string `json:"username"`
	Avatar       string `json:"avatar"`    //头像
	Signature    string `json:"signature"` //个性签名
	IsFollow     bool   `json:"is_follow"`
	IsMe         bool   `json:"is_me"`    // 是否是当前用户
	PostCnt      int64  `json:"post_cnt"` //发布的帖子数量
	LikeCnt      int64  `json:"like_cnt"`
	FollowerCnt  int64  `json:"follower_cnt"`  //粉丝数量
	FollowingCnt int64  `json:"following_cnt"` //关注数量
}

type UserService interface {
	// GetUserList 获得全部User对象
	GetUserList() []models.TableUser
	// GetUserByUsername 根据username获得User对象
	GetUserByUsername(name string) (models.TableUser, error)
	// GetFmtUserByUsername 根据username获得FmtUser对象
	GetFmtUserByUsername(name string) (FmtUser, error)
	// InsertUser 将user插入表内
	InsertUser(user *models.TableUser) bool
	// GetFmtUserById 未登录情况下,根据user_id获得User对象
	GetFmtUserById(id int64) (FmtUser, error)
	// GetFmtUserByIdWithCurId 已登录(curID)情况下,根据user_id获得User对象
	GetFmtUserByIdWithCurId(id int64, curId int64) (FmtUser, error)
	// UpdateAvatar 更新用户头像
	UpdateAvatar(userId int64, avatar string) error
}
