package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"socialnetwork_back_go/config"
	"strings"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Auth 若用户携带的token正确,解析token,将userId放入上下文context中并放行;否则,返回错误信息
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 { // 如果未携带token
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		}
		auth = strings.Fields(auth)[1]
		log.Println("auth: ", auth)
		token, err := parseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Token Error",
			})
		}
		c.Set("userId", token.Id)
		c.Next()
	}
}

// AuthWithoutLogin 未登录情况下,若携带token,则解析出用户id并放入context;若未携带,则放入用户id默认值0
func AuthWithoutLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Query("token")
		log.Println("用户token!!!!!!: ", auth)
		var userId string
		if len(auth) == 0 { // 若未携带token，设置userId为0
			userId = "0"
		} else { // 若携带token，则解析
			auth = strings.Fields(auth)[1]
			token, err := parseToken(auth)
			if err != nil {
				context.Abort()
				context.JSON(http.StatusUnauthorized, Response{
					StatusCode: -1,
					StatusMsg:  "Token Error",
				})
			} else {
				userId = token.Id
				println("token 正确")
			}
		}
		context.Set("userId", userId)
		context.Next()
	}
}

// parseToken 解析token
func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
