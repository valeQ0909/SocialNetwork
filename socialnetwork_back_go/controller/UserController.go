package controller

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(200, "Hello World !!!")
}
