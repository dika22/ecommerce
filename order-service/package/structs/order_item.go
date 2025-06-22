package structs

import "time"

type OrderItem struct {
    ID           uint      `gorm:"primaryKey"`
    OrderID      int64     `gorm:"foreignKey:OrderID"`
    ProductID    int64	   `gorm:"foreignKey:ProductID"`
    Quantity     int64 	   `gorm:"quantity"`
    WarehouseID  int64     `gorm:"warehouse_id"`
	CreatedAt    time.Time `gorm:"created_at"`
}


func (o OrderProductItem) NewOrderItem(orderID int64) OrderItem  {
	return OrderItem{
		OrderID:      orderID,
		ProductID:    o.ProductID,
		Quantity:     o.Quantity,
		WarehouseID:  o.WarehouseID,
		CreatedAt:    time.Now(),
	}
}