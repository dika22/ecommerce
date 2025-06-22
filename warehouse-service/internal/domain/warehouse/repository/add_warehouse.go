package repository

import (
	"context"
	"warehouse-service/package/structs"
)


func (r WarehouseRepository)AddWarehouse(ctx context.Context, warehouse *structs.Warehouse) error {
	return r.db.Create(warehouse).Error
}