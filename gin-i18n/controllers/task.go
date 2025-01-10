package controllers

import (
	"encoding/json"
	"fmt"
	"gin-i18n/controllers/provider"
	"gin-i18n/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateTask 创建任务
func CreateTask(ctx *gin.Context) {
	task := models.Task{}
	title := ctx.PostForm("title")
	if title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	task.Title = title
	// 获取上传的文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	// 将文件保存到本地临时目录
	tempFile := fmt.Sprintf("./%s", file.Filename) // todo 根据userid + 时间戳 切割目录
	if err := ctx.SaveUploadedFile(file, tempFile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	task.SourceFile = tempFile
	// defer os.Remove(tempFile) // 使用后删除临时文件

	// 打开并读取文件内容
	fileContent, err := os.ReadFile(tempFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// 解析 JSON 文件内容
	var texts interface{}
	if err := json.Unmarshal(fileContent, &texts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
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
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// UpdateTask 更新任务信息
func DownloadTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	filePath := fmt.Sprintf("./%s.json", id) // 文件在服务器上的路径
	fileName := fmt.Sprintf("%s.json", id)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
	}

	// 下载文件
	ctx.FileAttachment(filePath, fileName)

	ctx.JSON(http.StatusOK, gin.H{"message": "download successfully"})
}

// UpdateTask 更新任务信息
func TranslateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := DB.First(&task, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// 打开并读取文件内容
	fileContent, err := os.ReadFile(task.SourceFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// 解析 JSON 文件内容
	var texts map[string]string
	if err := json.Unmarshal(fileContent, &texts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	transferClient, err := provider.GetTransProvider("tencent")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get translate provider"})
		return
	}

	translatedTexts, err := transferClient.Translate(texts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate task"})
		return
	}

	if err := DB.Model(&models.Task{}).Where("id = ?", id).Update("status", models.TaskStatusDoing).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate task"})
		return
	}
	targetFile := fmt.Sprintf("./%s.json", id)
	file, err := os.Create(targetFile)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// 3. 将 map 转为 JSON 并写入文件
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 格式化 JSON 输出
	if err := encoder.Encode(translatedTexts); err != nil {
		fmt.Printf("Failed to write JSON to file: %v\n", err)
		return
	}

	if err := DB.Model(&models.Task{}).Where("id = ?", id).Update("target_file", targetFile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to translate task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "task translated successfully"})
}
