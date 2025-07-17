package repository

import (
	"context"
	"product-service/package/structs"
)

func (r CategoryProductRepository) Update(ctx context.Context, payload interface{}) error {
	return r.db.Where("category_id = ?", payload.(*structs.CategoryProduct).ID).Updates(payload).Error
}