package repository

import (
	"blog_post/adaptor"
	"blog_post/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateCategory(data models.Category) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Create(data); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("can't create category try again later")
	}
	return nil
}

// Helps to read all data
func ReadCategories(dest *[]models.Category) error {
	if result := adaptor.GetConn().Find(&dest); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to delete the category
func DeleteCategory(id string) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Category{}).Where("category_id = ?", id).Unscoped().Delete(&models.Category{}); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid category id")
	}
	if result = adaptor.GetConn().Model(&models.Post{}).Where("? = ANY(category_id)", id).Update("category_id", gorm.Expr("array_remove(category_id, ?)", id)); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to update the category
func UpdateCategory(id string, content models.Category) error {
	var result *gorm.DB
	if result = adaptor.GetConn().Where("category_id = ?", id).Updates(content); result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid category id")
	}
	return nil
}

// Helps to read the category by its id
func ReadCategoryByID(id string) (string, error) {
	var name string
	var result *gorm.DB
	if result = adaptor.GetConn().Model(&models.Category{}).Where("category_id = ?", id).Select("name").Scan(&name); result.Error != nil {
		return name, result.Error
	}
	if result.RowsAffected == 0 {
		return "",fmt.Errorf("invalid category id of %v", id)
	}
	return name, nil
}
