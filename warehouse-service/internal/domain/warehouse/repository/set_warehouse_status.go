package repository

import (
	"context"
	"warehouse-service/package/structs"
)


func (r WarehouseRepository) SetWarehouseStatus(c context.Context, IdWarehouse int64) error{
	return r.db.Model(&structs.Warehouse{}).Where("id = ?", IdWarehouse).Update("is_active", false).Error
}