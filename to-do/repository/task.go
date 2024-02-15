package repository

import (
	"fmt"
	"to-do/models"

	"gorm.io/gorm"
)

type Task interface {
	AddTask(task models.Tasks) error
	UpdateTask(id string, content string) error
	DeleteTask(id string) error
	ReadTaskByID(id string, task *models.Tasks) error
	ReadDeletedTask(id string, tasks *[]models.Tasks) error
	UpdateAllStatus(id string) error
}

// Helps to add task
func (d *DB_Connection) AddTask(task models.Tasks) error {
	return d.DB.Create(&task).Error
}

// helps to update the task
func (d *DB_Connection) UpdateTask(id string, content string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Tasks{}).Where("task_id = ?", id).Update("task", content); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such task exist invalid task id")
	}
	return nil
}

// Helps to delete the task
func (d *DB_Connection) DeleteTask(id string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Tasks{}).Where("task_id = ?", id).Delete(id); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such task exist invalid task id")
	}
	return d.UpdateTaskStatus(id, false)
}

// helps to get all deleted task
func (d *DB_Connection) ReadDeletedTask(id string, tasks *[]models.Tasks) error {
	return d.DB.Model(models.Tasks{}).Unscoped().Where("user_id = ? AND deleted_at != nil", id).Find(&tasks).Error
}

// Helps to read the task by its ID
func (d *DB_Connection) ReadTaskByID(id string, task *models.Tasks) error {
	var result *gorm.DB
	if result = d.DB.Model(models.Tasks{}).Where("task_id = ?", id).Find(&task); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such task exist invalid task id")
	}
	return nil

}

// Helps to update the task status
func (d *DB_Connection) UpdateTaskStatus(id string, status bool) error {
	return d.DB.Model(models.Tasks{}).Unscoped().Where("task_id = ?", id).Update("active", status).Error
}

// helps to update all tasks status
func (d *DB_Connection) UpdateAllStatus(id string) error {
	var flag bool
	if err := d.DB.Model(models.Tasks{}).Select("active").Where("user_id = ? AND active = ?", id, true).Scan(&flag).Error; err != nil {
		return err
	}

	if flag {
		if err := d.DB.Model(models.Tasks{}).Where("user_id = ?", id).Update("active", false).Error; err != nil {
			return err
		}
	}
	return d.DB.Model(models.Tasks{}).Where("user_id = ?", id).Update("active", true).Error

}
