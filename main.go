package main

import (
	"github.com/gin-gonic/gin"
	"iniyou.com/common"
	"iniyou.com/routes"
)

func main() {
	db := common.InitDB()
	dbclose, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbclose.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	err = r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	panic(err)
}
