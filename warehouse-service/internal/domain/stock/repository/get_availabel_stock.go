package repository

import (
	"context"
	"warehouse-service/package/structs"
)

func (r StokRepository) GetAvailableStockByProductId(ctx context.Context, productID int64) (int64, error) {
	var totalStock int64
	err :=r.db.Model(&structs.Stock{}).
	Joins("JOIN warehouses ON warehouses.id = stocks.warehouse_id").
	Where("stocks.product_id = ? AND warehouses.is_active = ?", productID, true).
	Select("SUM(stocks.quantity)").Scan(&totalStock).Error
	if err != nil {
		return 0, err
	}
	return totalStock, nil
}