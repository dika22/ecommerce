package structs

type RequestCreateOrder struct {
	Items []OrderProductItem `json:"items"`
}

type OrderProductItem struct {
	ProductID   int64    `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	WarehouseID int64  `json:"warehouse_id"`
}