package routers

import (
	"gin-i18n/controllers"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// todo!!!
		ctx.Next()
	}
}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) {

	rateLimiter := NewRateLimiter(2, 5)
	r.Use(RateLimitMiddleware(rateLimiter))

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/users", controllers.CreateUser)
		authRoutes.POST("/login", controllers.Login)
	}

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("/", authMiddleware(), controllers.CreateTask)
		taskRoutes.POST("/:id/translate", authMiddleware(), controllers.TranslateTask)
		taskRoutes.GET("/:id", authMiddleware(), controllers.GetTaskByID)
		taskRoutes.GET("/:id/download", authMiddleware(), controllers.DownloadTask)
	}
}
