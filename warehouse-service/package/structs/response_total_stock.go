package structs

type ResponseTotalStock struct {
	ProductID  int64 `json:"product_id"`
	TotalStock int64 `json:"total_stock"`
}