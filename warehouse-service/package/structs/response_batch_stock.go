package structs

type ResponseBatchStock struct {
	AllAvailable bool `json:"all_available"`
	StockProducts     []StockProduct `json:"products"`
}

type StockProduct struct {
	ProductID   int64    `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	WarehouseID int64  `json:"warehouse_id"`
	Available   bool `json:"available"`
}
