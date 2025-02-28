package routes

import (
	"github.com/gin-gonic/gin"
	"iniyou.com/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Reigster)

	return r
}
