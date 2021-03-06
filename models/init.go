package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化postgresql链接
func Database(connString string) {
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		fmt.Println(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
