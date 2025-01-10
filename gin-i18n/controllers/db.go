package controllers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局数据库实例

// InitDB 初始化数据库连接
func InitDB() {
	// MySQL 配置
	user := "demos"          // 数据库用户名
	password := "123123"     // 数据库密码
	host := "dev.asdfin.com" // 数据库地址
	port := "3306"           // 数据库端口
	database := "tuwei_i18n" // 数据库名称

	// 构建 DSN（数据源名称）
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}
	log.Println("MySQL Database connected successfully")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get raw DB connection: %v", err)
		return
	}
	sqlDB.Close()
}
