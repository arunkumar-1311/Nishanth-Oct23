package lookup

import (
	"online-purchase/models"

	"gorm.io/gorm"
)

func (*Empty) Lookup_003(db *gorm.DB) error {
	orderStatus := []models.OrderStatus{
		{OrderStatusID: "S-AS", Status: "Assembly"},
		{OrderStatusID: "S-UP", Status: "Underprocess"},
		{OrderStatusID: "S-SH", Status: "Shipped"},
		{OrderStatusID: "S-DD", Status: "Delivered"},
		{OrderStatusID: "S-NULL", Status: "null"},
	}

	if err := db.Create(orderStatus).Error; err != nil {
		return err
	}
	return nil
}
