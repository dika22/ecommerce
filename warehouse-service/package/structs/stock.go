package structs

import "time"

type Stock struct {
    ID          int64 `gorm:"primaryKey"`
    ProductID   int64
    WarehouseID int64
    Quantity    int64
    CreatedAt   time.Time
	UpdatedAt   time.Time
}


func (p RequestAddStock) NewStock() Stock {
	return Stock{
		ProductID:   p.ProductID,
		WarehouseID: p.WarehouseID,
		Quantity:    p.Quantity,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	
}