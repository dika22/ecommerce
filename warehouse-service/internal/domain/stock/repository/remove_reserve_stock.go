package repository

import (
	"context"
	"warehouse-service/package/structs"
)


func (r StokRepository) RemoveReservation(ctx context.Context, req *structs.RequestReleaseStock) error {
	return r.db.Where("order_id = ? AND product_id = ? AND warehouse_id = ?", 
    req.OrderID, req.ProductID, req.WarehouseID).
    Delete(&structs.ReservedStock{}).Error
}