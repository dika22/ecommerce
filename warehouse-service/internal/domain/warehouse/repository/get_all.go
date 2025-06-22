package repository

import (
	"context"
)

func (r WarehouseRepository) GetAll(ctx context.Context, dest interface{}) error{
	return r.db.Find(dest).Error	
}