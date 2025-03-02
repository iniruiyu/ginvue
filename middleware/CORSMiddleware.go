package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过中间件给header写入允许访问的域名或者是方法
		// *表示允许所有域名跨域
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// 设置缓存时间
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		// 设置允许请求的头部
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// 设置是否允许携带cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 判断是否是OPTIONS请求，是的话直接返回200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
