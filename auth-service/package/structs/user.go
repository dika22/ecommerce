package structs

import (
	"auth-service/internal/constant"
	"auth-service/package/utils"
	"time"
)

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Role      int64     `json:"role"`
	IsSeller  bool      `json:"is_seller"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p RequestSignUp) NewUser() User {
	return User{
		Email:     p.Email,
		Password:  utils.HashPassword(p.Password),
		Name:      p.Name,
		Role:      constant.RoleUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r RequestUpdateUser) NewUpdateUser() User {
	user := User{}
	user.ID = r.ID
	if r.Password != "" {
		r.Password = utils.HashPassword(r.Password)
	}
	if r.Email != "" {
		r.Email = user.Email
	}
	if r.Name != "" {
		r.Name = user.Name
	}
	return user
}
