package repository

import (
	"fmt"
	"to-do/models"

	"gorm.io/gorm"
)

type Task interface {
	AddTask(task models.Tasks) error

	ReadTaskByID(id string, task *models.Tasks) error
	ReadDeletedTask(id string, tasks *[]models.Tasks) error
	ReadAllTask(id string, dest *[]models.Tasks, status ...bool) error

	UpdateTask(id string, content string) error
	UpdateTaskStatus(id string, status bool) error
	UpdateAllStatus(id string) error

	DeleteTask(id string) error
	ClearCompleted(id string) error

	CountActiveTasks(id string, dest *int64) error
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
	return nil
}

// helps to get all deleted task
func (d *DB_Connection) ReadDeletedTask(id string, tasks *[]models.Tasks) error {
	return d.DB.Model(models.Tasks{}).Unscoped().Where("user_id = ? AND deleted_at IS NOT NULL", id).Find(&tasks).Error
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
	if status {
		if err := d.DB.Model(models.Tasks{}).Unscoped().Where("task_id = ?", id).Update("active", false).Error; err != nil {
			return err
		}
		return nil
	}
	return d.DB.Model(models.Tasks{}).Unscoped().Where("task_id = ?", id).Update("active", true).Error
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
		return nil
	}
	return d.DB.Model(models.Tasks{}).Where("user_id = ?", id).Update("active", true).Error

}

// Helps to read all user task
func (d *DB_Connection) ReadAllTask(id string, dest *[]models.Tasks, status ...bool) error {
	if len(status) > 0 {
		if err := d.DB.Model(&models.Tasks{}).Where("user_id = ? AND active = ?", id, status[0]).Find(&dest).Error; err != nil {
			return err
		}
		return nil
	}

	return d.DB.Model(&models.Tasks{}).Where("user_id = ?", id).Find(&dest).Error
}

// helps to count All active tasks
func (d *DB_Connection) CountActiveTasks(id string, dest *int64) error {
	return d.DB.Model(&models.Tasks{}).Where("user_id = ? AND active = ?", id, true).Count(dest).Error
}

// Helps to delete all completed tasks
func (d *DB_Connection) ClearCompleted(id string) error {
	return d.DB.Model(&models.Tasks{}).Unscoped().Where("active = ?", false).Or("deleted_at IS NOT NULL").Delete(id).Error
}
