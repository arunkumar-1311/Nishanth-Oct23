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
	ReadInactiveOrders(Orders *[]models.Orders) error
	DeleteOrderByID(string) error
	ReadAllOrderStatus(*[]models.OrderStatus) error
	UpdateOrderStatusByID(string, string) error
	ReadOrdersByUser(string, *[]models.Orders) error
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
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Where("active != ?", false).Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to fetch orders by user id
func (d *GORM_Connection) ReadOrdersByUser(id string, Orders *[]models.Orders) error {
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Where("user_id", id).Find(&Orders).Error; err != nil {
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

// Helps to read all inactive orders
func (d *GORM_Connection) ReadInactiveOrders(Orders *[]models.Orders) error {
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Unscoped().Where("active = ?", false).Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to delete the Order
func (d *GORM_Connection) DeleteOrderByID(id string) error {

	if result := d.DB.Model(&models.Orders{}).Where("order_id = ?", id).UpdateColumn("active", false); result.Error != nil {
		return result.Error
	}

	if result := d.DB.Model(&models.Orders{}).Where("order_id = ?", id).Delete(id); result.Error != nil {
		return result.Error
	}

	return nil
}

// Helps to read all order status
func (d *GORM_Connection) ReadAllOrderStatus(orderStatus *[]models.OrderStatus) error {

	if result := d.DB.Find(&orderStatus); result.Error != nil {
		return result.Error
	}
	return nil
}

// Helps to update the order status by id
func (d *GORM_Connection) UpdateOrderStatusByID(id string, status string) error {
	var result *gorm.DB
	if result = d.DB.Model(&models.Orders{}).Where("order_id = ?", id).Update("order_status_id", status); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no such order available")
	}

	if id == "S-DD" {
		if result = d.DB.Model(&models.Orders{}).Where("order_id = ?", id).UpdateColumn("active", false); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
