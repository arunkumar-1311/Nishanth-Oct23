package repository

import (
	"fmt"
	"online-purchase/models"

	"gorm.io/gorm"
)

type Brand interface {
	CreateNewBrand(models.Brand) error
	ReadBrands(*[]models.Brand) error
	ReadBrandByID(string, *models.Brand) error
	UpdateBrandByID(string, models.Brand) error
	DeleteBrandByID(string) error
}

// Helps to create new brand in brands table
func (d *GORM_Connection) CreateNewBrand(brand models.Brand) error {
	if err := d.DB.Create(&brand).Error; err != nil {
		return err
	}
	return nil
}

// Helps to get all the brands available in the application
func (d *GORM_Connection) ReadBrands(brands *[]models.Brand) error {
	if err := d.DB.Find(&brands).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read the brand by id
func (d *GORM_Connection) ReadBrandByID(id string, brand *models.Brand) error {
	var result *gorm.DB
	if result = d.DB.Where("brand_id = ?", id).Find(&brand); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid brand id")
	}
	return nil
}

// Helps to Update the brand by id
func (d *GORM_Connection) UpdateBrandByID(id string, brand models.Brand) error {
	var result *gorm.DB
	brand.BrandID = id
	if result = d.DB.Model(&models.Brand{}).Where("brand_id = ?", id).UpdateColumns(brand); result.Error != nil {
		return result.Error
	}

	if !brand.Status {
		if result = d.DB.Model(&models.Brand{}).Where("brand_id = ?", id).Update("status", false); result.Error != nil {
			return result.Error
		}
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid brand id")
	}
	return nil
}

// Helps to delete the brand
func (d *GORM_Connection) DeleteBrandByID(id string) error {
	var result *gorm.DB

	if result = d.DB.Model(&models.Brand{}).Where("brand_id = ?", id).Delete(id); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid brand id")
	}
	return nil
}
