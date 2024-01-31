package repository

import (
	"blog_post/models"
	"fmt"
	"gorm.io/gorm"
)

// Helps to set and access gorm connection
type GORM_Connection struct {
	DB *gorm.DB
}

// Methods helps to manipulate categories
type Categories interface{
	CreateCategory(models.Category) error
	ReadCategories(*[]models.CategoryResponse) error
	DeleteCategory(string) error
	UpdateCategory(string, models.Category) error
	ReadCategoryByID(string) (string, error)
}
// Helps to create the new category
func (d *GORM_Connection) CreateCategory(data models.Category) error {
	var result *gorm.DB
	if result = d.DB.Create(data); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("can't create category try again later")
	}
	return nil
}

// Helps to read all data
func (d *GORM_Connection) ReadCategories(dest *[]models.CategoryResponse) error {
	if result := d.DB.Model(&models.Category{}).Find(&dest); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to delete the category
func (d *GORM_Connection) DeleteCategory(id string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Category{}).Where("category_id = ?", id).Unscoped().Delete(&models.Category{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid category id")
	}
	if result = d.DB.Model(&models.Post{}).Where("? = ANY(category_id)", id).Update("category_id", gorm.Expr("array_remove(category_id, ?)", id)); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to update the category
func (d *GORM_Connection) UpdateCategory(id string, content models.Category) error {
	var result *gorm.DB
	if result = d.DB.Where("category_id = ?", id).Updates(content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid category id")
	}
	return nil
}

// Helps to read the category by its id
func (d *GORM_Connection) ReadCategoryByID(id string) (string, error) {
	var name string
	var result *gorm.DB
	if result = d.DB.Model(&models.Category{}).Where("category_id = ?", id).Select("name").Scan(&name); result.Error != nil {
		return name, result.Error
	}
	if result.RowsAffected == 0 {
		return "", fmt.Errorf("invalid category id of %v", id)
	}
	return name, nil
}
