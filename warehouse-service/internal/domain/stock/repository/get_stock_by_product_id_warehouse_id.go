package repository

import (
	"context"
	"warehouse-service/package/structs"
)


func (r StokRepository)  GetAvailableStockByProductIdWarehouseId(ctx context.Context, productID int64, warehouseID int64) (int64, error) {
	var totalStock int64
	err :=r.db.Model(&structs.Stock{}).
	Joins("JOIN warehouses ON warehouses.id = stocks.warehouse_id").
	Where("stocks.product_id = ? AND stocks.warehouse_id = ? AND warehouses.is_active = ?", productID, warehouseID, true).
	Select("SUM(stocks.quantity)").Scan(&totalStock).Error
	if err != nil {
		return 0, err
	}
	return totalStock, nil
	
}