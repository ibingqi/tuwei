package main

import (
	"gin-i18n/controllers"
	"gin-i18n/models"
	"gin-i18n/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	controllers.InitDB()
	defer controllers.CloseDB()

	// 自动迁移
	if err := controllers.DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 创建 Gin 实例
	r := gin.Default()

	// 注册路由
	routers.RegisterRoutes(r)

	// 启动服务器
	r.Run(":8080") // http://localhost:8080
}
