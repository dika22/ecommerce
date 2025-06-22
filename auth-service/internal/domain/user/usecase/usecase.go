package usecase

import (
	"auth-service/internal/domain/user/repository"
	"auth-service/package/structs"
	"context"
)

type IUser interface{
	SignUp(ctx context.Context, req structs.RequestSignUp) error
	Login(ctx context.Context, req structs.RequestLogin) (structs.ResponseLogin, error)
}

type UserUsecase struct{
	repo repository.UploadRepository
}


func NewUserUsecase(repo repository.UploadRepository) IUser  {
	return &UserUsecase{
		repo: repo,
	}
}
