package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/config"
	"socialnetwork_back_go/models"
	"socialnetwork_back_go/service"
	"strconv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	User  service.FmtUser `json:"user"`
	Token string          `json:"token"`
}

type UserInfoResponse struct {
	Response
	User service.FmtUser `json:"user"`
}

type UpdateAvatarResponse struct {
	Response
	Avatar string `json:"avatar"`
}

type UpdatePersonalInfoResponse struct {
	Response
	Username  string `json:"username"`
	Signature string `json:"signature"`
}

// Test /test/ GET
// 测试专用
func Test(c *gin.Context) {
	log.Println("Test: 执行")
	c.String(200, "Hello World !!!")
}

// Register /socialnetwork/user/register/ POST
// 用户注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	usi := service.UserServiceImpl{}

	// 查询该用户名是否被注册过
	u, _ := usi.GetUserByUsername(username)

	if username == u.Username { // 如果该用户名被注册过
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "该用户名已被注册"},
		})
	} else { // 如果没有注册过，就注册一个
		log.Println("该用户名未被注册过")
		newUser := models.TableUser{
			Username:  username,
			Password:  service.EnCoder(password), //对用户密码进行加密
			Avatar:    config.DefaultAvatar,
			Signature: config.DefaultSign,
		}
		if usi.InsertUser(&newUser) != true {
			println("创建用户失败")
		}
		// 创建用户成功
		u, _ := usi.GetUserByUsername(username)
		fmtUser := service.FmtUser{
			Id:        u.Id,
			UserName:  username,
			Avatar:    u.Avatar,
			Signature: u.Signature,
		}
		token := service.GenerateToken(username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			User:     fmtUser,
			Token:    token,
		})
	}
}

// Login POST /socialnetwork/user/login/
// 用户登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	encoderPassword := service.EnCoder(password)

	usi := service.UserServiceImpl{}
	u, _ := usi.GetUserByUsername(username)

	if encoderPassword == u.Password { //如果密码正确
		u, _ := usi.GetUserByUsername(username)
		fmtUser := service.FmtUser{
			Id:        u.Id,
			UserName:  username,
			Avatar:    u.Avatar,
			Signature: u.Signature,
		}
		token := service.GenerateToken(username)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			User:     fmtUser,
			Token:    token,
		})
	} else {
		log.Println("用户名称：", u.Username)
		log.Println("提交密码错误")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or Password Error"},
		})
	}
}

// UserInfo Post /socialnetwork/user/info/
// 获取用户信息
func UserInfo(c *gin.Context) {
	username := c.PostForm("username") //返回的这个any类型我不是很懂，留着以后再学
	log.Println("username: ", username)
	userService := service.UserServiceImpl{}

	if fmtUser, err := userService.GetFmtUserByUsername(username); err != nil {
		// 获取用户信息失败
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User Doesn't Exist"},
		})
	} else {
		// 成功获取用户信息
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:     fmtUser,
		})
	}
}

// Token GET /socialnetwork/user/token/
// 获取用户信息
func Token(c *gin.Context) {
	tokenId, _ := c.Get("tokenId") //返回的这个any类型我不是很懂，留着以后再学
	tokenIdstr := tokenId.(string) // 这个any必须转换未string，我也不是很懂，留着以后再学
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	usi := service.UserServiceImpl{}

	if fmtUser, err := usi.GetFmtUserById(userId); err != nil {
		// 获取用户信息失败
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User Doesn't Exist"},
		})
	} else {
		// 成功获取用户信息
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:     fmtUser,
		})
	}
}

// UpdateAvatar POST /user/updateavatar/
// 更新用户头像
func UpdateAvatar(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	file, err := c.FormFile("image_file")
	//log.Println("file: ", file)
	if err != nil {
		log.Println("用户上传头像图片失败")
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户上传头像图片失败",
		})
	}
	user, _ := models.GetTableUserById(userId)
	imagePath := "./media/images/" + user.Username + ".jpg"
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		log.Println("保存图片时发生了错误：", err)
		return
	}

	usi := service.UserServiceImpl{}
	imgUrl := "http://127.0.0.1:3000/media/images/" + user.Username + ".jpg"
	err = usi.UpdateAvatar(userId, imgUrl)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "更新头像失败",
		})
		return
	}

	log.Printf("成功更新头像")
	c.JSON(http.StatusOK, UpdateAvatarResponse{
		Response: Response{StatusCode: 0, StatusMsg: "更新头像成功"},
		Avatar:   imgUrl,
	})
}

// UpdatePersonalInfo POST /user/updatepersonalinfo/
// 更新用户的名称和签名
func UpdatePersonalInfo(c *gin.Context) {
	tokenId, _ := c.Get("tokenId")
	tokenIdstr := tokenId.(string)
	userId, _ := strconv.ParseInt(tokenIdstr, 10, 64)

	username := c.PostForm("username")
	signature := c.PostForm("signature")

	err := models.UpdatePersonalInfo(userId, username, signature)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "更新用户名称和签名失败",
		})
		return
	}

	log.Printf("成功更新用户名称和签名")
	c.JSON(http.StatusOK, UpdatePersonalInfoResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "更新用户名称和签名成功"},
		Username:  username,
		Signature: signature,
	})

}
