package structs

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (p RequestSignUp) NewUser() User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	return User{
		Email:     p.Email,
		Password:  string(hashed),
		Name:      p.Name,
		CreatedAt: time.Now(),
	}
}
