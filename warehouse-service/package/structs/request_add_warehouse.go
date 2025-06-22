package structs

type RequestAddWarehouse struct {
	ShopID int64 `json:"shop_id"`
	Name string `json:"name"`
	Address string `json:"address"`
}