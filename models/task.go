package models

import (
	"log"

	"github.com/hellojukay/todo-server/utils"
	"gorm.io/gorm"
)

// Task 任务实体
type Task struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"title" json:"title"`
	Desc  string `gorm:"desc" json:"desc"`
}

func ListAllTasks() []Task {
	var tasks []Task
	if err := Store.Find(&tasks).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Fatal(err)
	}
	return tasks
}

func SaveTask(task Task) (Task, error) {
	if task.ID == 0 {
		task.ID = utils.NextID()
	}
	var err = Store.Save(&task).Error
	return task, err
}

func RemoveTask(id int64) error {
	return Store.Delete(Task{ID: id}).Error
}
