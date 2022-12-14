package repository

import "time"

type Task struct {
	ID       uint   `gorm:"primaryKey"` //ID
	TaskTime string //任务日期 yyyy-MM-dd
	Total    uint   //售出套数
}

func (Task) TableName() string {
	return "task"
}

type TaskMapper struct {
}

func (TaskMapper) AddTask() uint {
	var task Task
	task.TaskTime = time.Now().Format("2006-01-02")
	DB.Create(&task)
	return task.ID
}

func (TaskMapper) UpdateTotalById(id uint, total int) {
	DB.Exec("update task set total = ? where id = ?", total, id)
}

func (TaskMapper) FindAll() []Task {
	var tasks []Task
	DB.Raw("select * from task").Scan(&tasks)
	return tasks
}

func (TaskMapper) FindLatestTaskIds() []uint {
	var ids []uint
	DB.Raw("select id from task order by id desc limit 2").Scan(&ids)
	return ids
}
