package repository

import (
	"fmt"
	"online-purchase/models"

	"gorm.io/gorm"
)

type Ram interface {
	CreateNewRam(models.Ram) error
	ReadRAMs(*[]models.Ram) error
	ReadRAMByID(string, *models.Ram) error
	UpdateRAMByID(string, models.Ram) error
	DeleteRAMByID(string) error
}

// Helps to create new RAM in RAM table
func (d *GORM_Connection) CreateNewRam(ram models.Ram) error {
	if err := d.DB.Create(&ram).Error; err != nil {
		return err
	}
	return nil
}

// Helps to get all the RAM available in the application
func (d *GORM_Connection) ReadRAMs(ram *[]models.Ram) error {
	if err := d.DB.Find(&ram).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read the RAM by id
func (d *GORM_Connection) ReadRAMByID(id string, ram *models.Ram) error {
	var result *gorm.DB
	if result = d.DB.Where("ram_id = ?", id).Find(&ram); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid RAM id")
	}
	return nil
}

// Helps to Update the RAM by id
func (d *GORM_Connection) UpdateRAMByID(id string, ram models.Ram) error {
	var result *gorm.DB
	ram.RamID = id
	if result = d.DB.Model(&models.Ram{}).Where("ram_id = ?", id).UpdateColumns(ram); result.Error != nil {
		return result.Error
	}

	if !ram.Status {
		if result = d.DB.Model(&models.Ram{}).Where("ram_id = ?", id).Update("status", false); result.Error != nil {
			return result.Error
		}
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid RAM id")
	}
	return nil
}

// Helps to delete the RAM
func (d *GORM_Connection) DeleteRAMByID(id string) error {
	var result *gorm.DB

	if result = d.DB.Model(&models.Ram{}).Where("ram_id = ?", id).Delete(id); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid RAM id")
	}
	return nil
}
