package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iniyou.com/common"
	"iniyou.com/model"
	"iniyou.com/utils"
)

func Reigster(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(name) == 0 {
		name = utils.RandomString(10)
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": "ok",
				"msg":  name})
		//return
	}
	if len(telephone) != 11 {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": "422",
				"msg":  "手机号错误",
				"len":  telephone})
		return
	}
	if len(password) < 6 {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": "422",
				"msg":  "密码过短"})
		return
	}

	// 判断手机号是否存在

	if isTelephoneExist(DB, telephone) {
		// 不允许注册
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"code": "422",
				"msg": "用户已存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	// 返回结果

	c.JSON(200, gin.H{

		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) (b bool) {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		b = true
	}
	return
}
