package usecase

import (
	"auth-service/internal/domain/seller/repository"
	userRepo "auth-service/internal/domain/user/repository"
	"auth-service/package/structs"
	"context"
)

type ISeller interface {
	CreateSeller(ctx context.Context, req structs.RequestCreateSeller) error
	UpdateSeller(ctx context.Context, req structs.RequestUpdateSeller) error
}

type SellerUsecase struct {
	repo repository.SellerRepository
	userRepo userRepo.Repository
}

func NewSellerUsecase(repo repository.SellerRepository, userRepo userRepo.Repository) ISeller {
	return &SellerUsecase{
		repo: repo,
		userRepo: userRepo,
	}
}