package repository

import "time"

type Task struct {
	ID       uint   `gorm:"primaryKey"` //ID
	TaskTime string //任务日期 yyyy-MM-dd
}

func (Task) TableName() string {
	return "task"
}

type TaskMapper struct {
}

func (TaskMapper) addTask() uint {
	var task Task
	task.TaskTime = time.Now().Format("2006-01-02")
	DB.Create(&task)
	return task.ID
}
