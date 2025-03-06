package controller

import (
	"errors"
	"strconv"

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
	// 绑定body中的参数
	var requesrCategory model.Category
	ctx.Bind(&requesrCategory)
	if requesrCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称是必填项")
		return
	}
	// 获取Path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//需要更新的分类
	var updateCategory model.Category
	// RecordNotFound 方法已经被移除
	result := c.DB.First(&updateCategory, categoryId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类
	// Update() 三种类型参数
	// 1. map
	// 2. struct
	// 3. name value
	c.DB.Model(&updateCategory).Update("name", requesrCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取Path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	//需要查询的分类
	var Category model.Category
	// RecordNotFound 方法已经被移除
	result := c.DB.First(&Category, categoryId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": Category}, "")
}
func (c CategoryController) Remove(ctx *gin.Context) {
	// 获取Path中的参数
	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		response.Fail(ctx, nil, "无效的分类ID")
		return
	}

	// 直接根据ID删除，不需要先查询再删除
	result := c.DB.Delete(&model.Category{}, categoryId)

	// 检查是否有行被影响（即删除成功）
	if result.RowsAffected == 0 {
		response.Fail(ctx, nil, "删除失败，分类不存在")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
