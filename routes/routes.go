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

	CategoryRouters := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	CategoryRouters.POST("", categoryController.Create)
	CategoryRouters.PUT("/:id", categoryController.Update) // 传一个更新分类的ID
	CategoryRouters.GET("/:id", categoryController.Show)
	CategoryRouters.DELETE("/:id", categoryController.Remove)

	//CategoryRouters.PATCH() // PATCH局部修改，put是替换，一个模型替换另一个模型，patch补丁，只修改其中的一部分
	return r
}
