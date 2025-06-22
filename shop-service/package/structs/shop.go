package structs

import "time"

type Shop struct {
	ID       	int64      `json:"id"`
	Name     	string 	   `json:"name"`
	Address 	string 	   `json:"address"`
	CreatedAt 	time.Time  `json:"created_at"`
	UpdatedAt 	time.Time  `json:"updated_at"`
}


func (p RequestCreateShop) NewShop() Shop {
	return Shop{
		Name:       p.Name,
		Address:    p.Address,
		CreatedAt:  time.Now(),
	}
}