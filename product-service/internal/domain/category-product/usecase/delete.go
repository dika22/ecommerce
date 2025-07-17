package usecase

import (
	"context"
)

func (u *CategoryProductUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}