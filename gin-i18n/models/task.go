package models

// User 用户模型
type Task struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}
