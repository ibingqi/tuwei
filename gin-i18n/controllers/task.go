package controllers

import (
	"fmt"
	"gin-i18n/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTask 创建任务
func CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

// GetTaskByID 根据 ID 获取任务
func GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// UpdateTask 更新任务信息
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
