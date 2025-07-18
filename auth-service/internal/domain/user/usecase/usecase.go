package usecase

import (
	"auth-service/internal/domain/user/repository"
	"auth-service/package/structs"
	"context"
)

type IUser interface{
	SignUp(ctx context.Context, req structs.RequestSignUp) error
	Login(ctx context.Context, req structs.RequestLogin) (structs.ResponseLogin, error)
	GetByEmail(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, req structs.RequestUpdateUser) error 
}

type UserUsecase struct{
	repo repository.UserRepository
}


func NewUserUsecase(repo repository.UserRepository) IUser  {
	return &UserUsecase{
		repo: repo,
	}
}
