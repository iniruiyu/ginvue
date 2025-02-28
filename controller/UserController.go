package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	if isTelephoneExist(telephone) {
		// 不允许注册
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"code": "422",
				"msg": "用户已存在"})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "500",
				"msg": "密码加密出错"})
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 返回结果

	c.JSON(200, gin.H{

		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
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
	if isTelephoneExist(telephone) {

	} else {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": "422",
				"msg":  "用户不存在"})
		return
	}
	// 判断密码是否正确
	db := common.GetDB()
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": "400",
				"msg":  "密码错误"})
		return
	}

	// 发放Token给前端
	token := "1111111111111111111"
	// 返回结果
	c.JSON(200, gin.H{
		"code":    "200",
		"token":   token,
		"message": "登录成功",
	})
}

func isTelephoneExist(telephone string) (b bool) {
	db := common.GetDB()
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		b = true
	}
	return
}
