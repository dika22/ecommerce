package structs

import "time"

type ReservedStock struct {
    ID          int64      `gorm:"primaryKey"`
    OrderID     int64      // dari Order Service
    ProductID   int64      // dari Stock Service
    WarehouseID int64      // dari Warehouse Service
    Quantity    int64
    CreatedAt   time.Time
}


func (s RequestReserveStock) NewReservedStock() ReservedStock {
	return ReservedStock{
		OrderID:     s.OrderID,
		ProductID:   s.ProductID,
		WarehouseID: s.WarehouseID,
		Quantity:    s.Quantity,
		CreatedAt:   time.Now(),
	}
}
