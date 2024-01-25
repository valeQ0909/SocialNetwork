package service

import "socialnetwork_back_go/models"

// FmtUser 最终封装后,controller返回的FmtUser结构体
type FmtUser struct {
	Id        int64  `json:"id,omitempty"`
	UserName  string `json:"username,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Signature string `json:"signature,omitempty"`
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
