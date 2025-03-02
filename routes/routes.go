package routes

import (
	"github.com/gin-gonic/gin"
	"iniyou.com/controller"
	"iniyou.com/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Reigster)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
