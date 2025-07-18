package structs

type RequestGetStock struct {
	WarehouseID int64 `json:"warehouse_id"`
	ProductID   int64 `json:"product_id"`
}