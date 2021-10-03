package models

import (
	"task-tracker/src/config"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Team        int64     `json:"team_id"`
	Summary     string    `json:"task_summary"`
	Description string    `json:"task_description"`
	Deadline    time.Time `json:"task_deadline"`
	Completed   bool      `json:"task_completed"`
}

func GetTask(task *Task, id string) error {
	if err := config.DB.First(&task, id).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func GetTasks(tasks *[]Task, filter map[string]interface{}) error {
	if err := config.DB.Find(&tasks, filter).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func CreateTask(task *Task) error {
	if err := config.DB.Create(&task).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateTask(task *Task, id string) error {
	var temp Task
	if err := config.DB.First(&temp, id).Error; err != nil {
		return err
	}
	config.DB.Save(&task)
	return nil
}

func DeleteTask(task *Task, id string) error {
	if err := config.DB.Delete(&task, id).Error; err != nil {
		return err
	} else {
		return nil
	}
}
