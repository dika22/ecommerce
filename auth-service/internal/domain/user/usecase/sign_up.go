package usecase

import (
	"auth-service/package/structs"
	"context"
)

func (u *UserUsecase) SignUp(ctx context.Context, req structs.RequestSignUp) error {
	payload := req.NewUser()
	if err := u.repo.StoreUser(ctx, payload); err != nil {
		return err
	}
	return nil
}
