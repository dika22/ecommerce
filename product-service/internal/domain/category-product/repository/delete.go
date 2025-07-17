package repository

import (
	"context"
	"product-service/package/structs"
)

func (r CategoryProductRepository) Delete(ctx context.Context, id int64) error {
	return r.db.Where("category_id = ?", id).Delete(&structs.CategoryProduct{}).Error
}