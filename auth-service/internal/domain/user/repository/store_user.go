package repository

import (
	"auth-service/package/structs"
	"context"
)

func (r UploadRepository) StoreUser(ctx context.Context, payload structs.User) error {
	return r.db.Table("users").Create(&payload).Error
}
