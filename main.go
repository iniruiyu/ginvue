package main

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"iniyou.com/common"
	"iniyou.com/routes"
)

func main() {
	InitConfig()
	db := common.InitDB()
	dbclose, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbclose.Close()

	r := gin.Default()

	r = routes.CollectRoute(r)

	server := viper.GetString("server.host")
	port := viper.GetString("server.port")
	err = r.Run(server + ":" + port) // 监听并在 0.0.0.0:8080 上启动服务
	panic(err)
}

func InitConfig() {
	//获取当前工作目录
	workDir, _ := os.Getwd()
	viper.SetConfigName("application") // 指定配置文件名称，不需要文件扩展名
	viper.SetConfigType("yaml")        // 指定配置文件类型
	path := path.Join(workDir, "config")
	viper.AddConfigPath(path) // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
