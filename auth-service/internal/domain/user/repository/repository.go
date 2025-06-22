package repository

import (
	"auth-service/package/structs"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	StoreUser(ctx context.Context, payload structs.User) error
	GetByID(ctx context.Context, id uint) error
	GetByEmail(ctx context.Context, email string) error
}

type UploadRepository struct {
	db *gorm.DB
}

func NewUserRepository(g *gorm.DB) UploadRepository {
	return UploadRepository{
		db: g,
	}
}
