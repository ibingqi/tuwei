package routers

import (
	"gin-i18n/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) {

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/users", controllers.CreateUser)
		authRoutes.POST("/login", controllers.Login)
	}

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", controllers.CreateTask)    // 创建用户
		taskRoutes.GET("/:id", controllers.GetTaskByID) // 根据 ID 获取用户
		taskRoutes.PUT("/:id", controllers.UpdateTask)  // 更新用户
	}
}
