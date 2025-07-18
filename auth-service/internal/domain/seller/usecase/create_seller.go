package usecase

import (
	"auth-service/internal/constant"
	"auth-service/package/structs"
	"context"
	"fmt"
	"time"
)

func (u *SellerUsecase) CreateSeller(ctx context.Context, req structs.RequestCreateSeller) error {
	fmt.Println("DEBUG", req)
	newSeller := req.NewSeller()
	if err := u.repo.StoreSeller(ctx, newSeller); err != nil {
		return err
	}

	if err := u.userRepo.UpdateStatusUser(ctx, structs.RequestUpdateStatusUser{
		ID:        newSeller.UserID,
		IsSeller:  true,
		Role: constant.RoleSeller,
		UpdatedAt: time.Now(),
	}); err != nil {
		return err
	}
	return  nil
}