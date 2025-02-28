package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iniyou.com/model"
)

var DB *gorm.DB

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
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
