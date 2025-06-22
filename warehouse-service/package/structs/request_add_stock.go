package structs

type RequestAddStock struct {
	WarehouseID int64 `json:"warehouse_id"`
	ProductID   int64 `json:"product_id"`
	Quantity    int64 `json:"quantity"`
}