package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"log"
	"socialnetwork_back_go/config"
	"socialnetwork_back_go/models"
	"strconv"
	"time"
)

type UserServiceImpl struct {
}

// GetUserList 获取全部的User对象
func (usi *UserServiceImpl) GetUserList() []models.TableUser {
	tableUsers, err := models.GetTableUserList()
	if err != nil {
		log.Println("Error: ", err.Error())
		return tableUsers
	}
	return tableUsers
}

// GetUserByUsername 通过name获取User对象
func (usi *UserServiceImpl) GetUserByUsername(name string) models.TableUser {
	tableUser, err := models.GetTableUserByUsername(name)
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("User Not Found")
		log.Println("tableUser: ", tableUser)
		return tableUser
	}
	log.Println("Query User Success")
	return tableUser
}

// InsertUser 创建一个User对象
func (usi *UserServiceImpl) InsertUser(user *models.TableUser) bool {
	flag := models.InsertTableUser(user)
	if flag == false {
		log.Println("插入失败")
		return false
	}
	return true
}

// GetFmtUserById 通过Id获取格式化的FmtUser对象
func (usi *UserServiceImpl) GetFmtUserById(id int64) (FmtUser, error) {
	fmtUser := FmtUser{
		Id:        0,
		UserName:  "",
		Avatar:    config.DefaultAvatar,
		Signature: config.DefaultSign,
	}

	user, err := models.GetTableUserById(id)
	if err != nil {
		log.Println("产生错误:", err.Error())
		log.Println("没有查到用户id为", id, "的用户")
		return fmtUser, err
	}
	fmtUser.UserName = user.Username
	fmtUser.Id = user.Id

	return fmtUser, nil
}

// GetFmtUserByIdWithCurId 已登录(curID)情况下,根据user_id获得User对象
func (usi *UserServiceImpl) GetFmtUserByIdWithCurId(id int64, curId int64) (FmtUser, error) {
	fmtUser := FmtUser{
		Id:        0,
		UserName:  "",
		Avatar:    config.DefaultAvatar,
		Signature: config.DefaultSign,
	}

	user, err := models.GetTableUserById(id)
	if err != nil {
		log.Println("产生错误:", err.Error())
		log.Println("没有查到用户id为", id, "的用户")
		return fmtUser, err
	}
	fmtUser.UserName = user.Username
	fmtUser.Id = user.Id
	fmtUser.Avatar = user.Avatar
	fmtUser.Signature = user.Signature

	return fmtUser, nil
}

// GenerateToken 根据username生成一个token
func GenerateToken(username string) string {
	u := UserService.GetUserByUsername(new(UserServiceImpl), username)
	token := NewToken(u)
	log.Println("token: ", token)
	return token
}

// NewToken 根据用户信息创建token
func NewToken(user models.TableUser) string {
	expiresTime := time.Now().Unix() + 60*60*24
	id64 := user.Id
	claims := jwt.StandardClaims{
		Audience:  user.Username,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		return token
	} else {
		return "fail"
	}
}

// EnCoder 密码加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
