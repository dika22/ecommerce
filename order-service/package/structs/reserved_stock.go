package structs

import "time"

type ReservedStock struct {
    ID          uint `gorm:"primaryKey"`
    ProductID   uint
    WarehouseID uint
    OrderID     uint
    Quantity    int
    CreatedAt   time.Time
}