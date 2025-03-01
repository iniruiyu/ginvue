package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"iniyou.com/common"
	"iniyou.com/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// vaildate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") { // Bearer 前端hearder 附带的token格式
			ctx.JSON(400, gin.H{
				"code": 401,
				"msg":  "token格式错误",
			})
			ctx.Abort() //将请求抛弃掉
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(400, gin.H{
				"code": 401,
				"msg":  "token解析出错",
			})
			ctx.Abort() //将请求抛弃掉
			return
		}
		// 证明Token通过了验证
		// 获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
			ctx.Abort() //将请求抛弃掉
			return
		}
		// 用户存在  将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next() //继续执行后续代码
	}

}
