package models

const (
	TaskStatusFailed  = -1
	TaskStatusPending = 0
	TaskStatusDoing   = 1
	TaskStatusDone    = 2
)

type Task struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `json:"user_id"`
	Title      string `json:"title"`
	Status     int    `json:"status"`
	SourceFile string `json:"file"`
	TargetFile string `json:"target_file"`
}
