package usecase

import (
	"auth-service/package/structs"
	"context"
)	

func (u *SellerUsecase) UpdateSeller(ctx context.Context, req structs.RequestUpdateSeller) error {
	newUpdateSeller := req.NewUpdateSeller()
	if err := u.repo.UpdateSeller(ctx, newUpdateSeller); err != nil {
		return err
	}
	return nil
}