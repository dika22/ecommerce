package repository

import "context"

func (r CategoryProductRepository) Store(ctx context.Context, payload interface{}) error {
	if err := r.db.Create(payload).Error; err != nil {
		return err
	}
	return nil
}