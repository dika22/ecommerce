package structs


type RequestReserveStock struct {
	OrderID     int64 `json:"order_id"`
	ProductID   int64 `json:"product_id"`
	WarehouseID int64 `json:"warehouse_id"`
	Quantity    int64  `json:"quantity"`
}