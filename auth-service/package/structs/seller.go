package structs

import "time"

type Seller struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	UserID     int64     `json:"user_id" gorm:"foreignKey:UserRefer"`
	NameSeller string    `json:"name_seller"`
	Address    string    `json:"address"`
	ImageURL   string    `json:"image_url"`
	Status     bool      `json:"status"`
	IsPremium  bool      `json:"is_premium"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (s RequestCreateSeller) NewSeller() Seller {
	return Seller{
		NameSeller: s.NameSeller,
		Address:    s.Address,
		ImageURL:   s.Image,
		Status:     true,
		IsPremium:  false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (r RequestUpdateSeller) NewUpdateSeller() Seller {	
	if r.IsPremium {
		r.IsPremium = true
	}

	if r.Status {
		r.Status = true
	}
	return Seller{
		ID:         r.Id,
		NameSeller: r.NameSeller,
		Address:    r.Address,
		UpdatedAt:  time.Now(),
	}
}