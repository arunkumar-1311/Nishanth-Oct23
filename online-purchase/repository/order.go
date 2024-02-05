package repository

import (
	"fmt"
	"online-purchase/models"
	"time"

	"gorm.io/gorm"
)

type Order interface {
	CreateNewOrder(models.Orders) error
	ReadOrders(*[]models.Orders) error
	ReadOrderByID(string, *models.Orders) error
	CancelOrderByID(string) error
	ReadAllOrderStatus(*[]models.OrderStatus) error
	UpdateOrderStatusByID(string, string) error
	ReadOrdersByUser(string, *[]models.Orders) error
	ReadByStatus(string, time.Time, time.Time, *[]models.Orders) error
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
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Where("order_status_id != ? AND order_status_id != ?", "S-DD", "S-CAN").Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to fetch orders by user id
func (d *GORM_Connection) ReadOrdersByUser(id string, Orders *[]models.Orders) error {
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Unscoped().Where("user_id", id).Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to read the Order by id
func (d *GORM_Connection) ReadOrderByID(id string, Order *models.Orders) error {
	var result *gorm.DB
	if result = d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Unscoped().Where("order_id = ?", id).Find(&Order); result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invalid Order id")
	}
	return nil
}


// Helps to read the order by status
func (d *GORM_Connection) ReadByStatus(id string, fromDate, toDate time.Time, Orders *[]models.Orders) error {
	if err := d.DB.Preload("Brand").Preload("Ram").Preload("OrderStatus").Unscoped().Where("updated_at between ? and ?", fromDate, toDate).Where("order_status_id", id).Find(&Orders).Error; err != nil {
		return err
	}
	return nil
}

// Helps to delete the Order
func (d *GORM_Connection) CancelOrderByID(id string) error {

	if err := d.UpdateOrderStatusByID(id, "S-CAN"); err != nil {
		return err
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
	
	return nil
}
