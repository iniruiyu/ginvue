package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iniyou.com/common"
	"iniyou.com/model"
	"iniyou.com/response"
)

// 借助编辑器生成一些常用的代码,定义一个接口
type ICategoryController interface {
	/*Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Remove(ctx *gin.Context)*/
	RestController
}

// 让结构体实现这个接口，编辑器自动生成代码,可惜，vscode不支持
type CategoryController struct {
	// 需要数据库连接池
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	// 自动迁移
	db.AutoMigrate(model.Category{})

	// 创建一个新的CategoryController实例
	return CategoryController{DB: db}

}

func (c CategoryController) Create(ctx *gin.Context) {
	var requesrCategory model.Category
	ctx.Bind(&requesrCategory)

	if requesrCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称是必填项")
		return
	}
	// 创建分类
	c.DB.Create(&requesrCategory)
	response.Success(ctx, gin.H{"category": requesrCategory}, "")
}
func (c CategoryController) Update(ctx *gin.Context) {

}
func (c CategoryController) Show(ctx *gin.Context) {

}
func (c CategoryController) Remove(ctx *gin.Context) {

}
