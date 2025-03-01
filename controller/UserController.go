package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"iniyou.com/common"
	"iniyou.com/dto"
	"iniyou.com/model"
	"iniyou.com/response"
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
		response.Response(c, http.StatusOK, 200, nil, "name长度为0，已分配随机昵称")
		//return
	}
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号错误")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码过短")
		return
	}

	// 判断手机号是否存在

	if isTelephoneExist(telephone) {
		// 不允许注册
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "密码加密出错")
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	// 返回结果

	/*c.JSON(200, gin.H{

		"message": "注册成功",
	})*/
	response.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证

	if len(name) == 0 {
		//name = utils.RandomString(10)
		response.Response(c, http.StatusOK, 200, nil, "name长度为0，已分配随机昵称")
		//return
	}
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号错误")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码过短")
		return
	}

	// 判断手机号是否存在
	if isTelephoneExist(telephone) {

	} else {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	db := common.GetDB()
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密码错误")
		return
	}

	// 发放Token给前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token 发放失败")
		return
	}
	// 返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
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
