package repository

import (
	"fmt"
	"online-purchase/models"

	"gorm.io/gorm"
)

type Order interface {
	CreateNewOrder(models.Orders) error
	ReadOrders(*[]models.Orders) error
	ReadOrderByID(string, *models.Orders) error
	UpdateOrderByID(string, models.Orders) error
	DeleteOrderByID(string) error
}

// Helps to create new order in orders table
func (d *GORM_Connection) CreateNewOrder(order models.Orders) error {
	if err := d.DB.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

// Helps to get all the Order available in the application
func (d *GORM_Connection) ReadOrders(Orders *[]models.Orders) error {
	if err := d.DB.Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read the Order by id
func (d *GORM_Connection) ReadOrderByID(id string, Order *models.Orders) error {
	var result *gorm.DB
	if result = d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Where("order_id = ?", id).Find(&Order); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Order id")
	}
	return nil
}

// Helps to Update the Order by id
func (d *GORM_Connection) UpdateOrderByID(id string, Order models.Orders) error {
	var result *gorm.DB
	Order.OrderID = id
	if result = d.DB.Model(&models.Orders{}).Where("order_id = ?", id).UpdateColumns(Order); result.Error != nil {
		return result.Error
	}

	if !Order.Active {
		if result = d.DB.Model(&models.Orders{}).Where("order_id = ?", id).Update("active", false); result.Error != nil {
			return result.Error
		}
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Order id")
	}
	return nil
}

// Helps to delete the Order
func (d *GORM_Connection) DeleteOrderByID(id string) error {
	var result *gorm.DB

	if result = d.DB.Model(&models.Orders{}).Where("order_id = ?", id).Delete(id); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Order id")
	}
	return nil
}
