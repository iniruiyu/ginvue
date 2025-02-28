package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;"`
	Telephone string `gorm:"type:varchar(110);not null;"`
	Password  string `gorm:"size:255;not null;"`
}

func main() {
	db := InitDB()
	dbclose, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbclose.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		// 数据验证
		if len(name) == 0 {
			name = RandomString(10)
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

		if isTelephoneExist(db, telephone) {
			// 不允许注册
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"code": "422",
					"msg": "用户已存在"})
			return
		}
		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		// 返回结果

		c.JSON(200, gin.H{

			"message": "注册成功",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func isTelephoneExist(db *gorm.DB, telephone string) (b bool) {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		b = true
	}
	return
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnwoqrstwvuxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	// 创建一个新的随机数生成器
	seed := time.Now().UnixNano() // 使用当前时间作为种子
	rng := rand.New(rand.NewSource(seed))

	for i := range result {
		result[i] = letters[rng.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {

	//dirverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin"
	username := "root"
	password := "root"
	charset := "utf8"
	// 构造 DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	// 使用 mysql.New 创建 Dialector 实例
	dialector := mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN 数据源名称
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failded to connect database,err= " + err.Error())
	}

	/*db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})*/

	// 自动创建数据表
	db.AutoMigrate(&User{})

	return db
}
