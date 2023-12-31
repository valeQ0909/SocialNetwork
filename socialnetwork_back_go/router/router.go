package router

import (
	"github.com/gin-gonic/gin"
	"socialnetwork_back_go/controller"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./static")
	
	r.GET("/test/", controller.Test)
}
