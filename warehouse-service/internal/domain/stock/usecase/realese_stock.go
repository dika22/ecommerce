package usecase

import (
	"context"
	"warehouse-service/package/structs"
)	

func (u *UsecaseStock) ReleaseStock(ctx context.Context, req *structs.RequestReleaseStock) error {
	return u.repo.RemoveReservation(ctx, req)
}