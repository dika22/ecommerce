package repository

import (
	"context"
	"warehouse-service/package/structs"
)

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