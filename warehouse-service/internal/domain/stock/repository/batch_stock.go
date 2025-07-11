package repository

import (
	"context"
	"warehouse-service/package/structs"
)

type ResponseBatchStock struct {
	AllAvailable bool `json:"all_available"`
	Products     []StockProduct `json:"products"`
}

type StockProduct struct {
	ProductID   int64    `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	WarehouseID int64  `json:"warehouse_id"`
	Available   bool `json:"available"`
}

func (r StokRepository) BatchStock(ctx context.Context, productIDs []int64) ([]*structs.Stock, error) {
	var stocks []*structs.Stock
	err :=r.db.Model(&structs.Stock{}).
	Joins("JOIN warehouses ON warehouses.id = stocks.warehouse_id").
	Where("stocks.product_id IN (?) AND warehouses.is_active = ?", productIDs, true).
	Select("stocks.product_id, stocks.quantity, stocks.warehouse_id").Scan(&stocks).Error
	if err != nil {
		return []*structs.Stock{}, err
	}
	return stocks, nil
}