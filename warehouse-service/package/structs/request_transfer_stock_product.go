package structs

type RequestTransferStockProduct struct {
	ProductID       int64 `json:"product_id"`
	FromWarehouseID int64 `json:"from_warehouse_id"`
	ToWarehouseID   int64 `json:"to_warehouse_id"`
	Quantity        int64 `json:"quantity"`
}