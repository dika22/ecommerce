package structs

import "time"

type Warehouse struct {
    ID        	uint      	`gorm:"primaryKey"`
	ShopeID   	int64     	`json:"shop_id"`
    Name      	string 		`json:"name"`
	Address 	string 		`json:"address"`
    IsActive  	bool      	`gorm:"default:true"`
    CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}


func (p RequestAddWarehouse) NewWarehouse() *Warehouse  {
	return &Warehouse{
		ShopeID:    p.ShopID,
		Name:       p.Name,
		Address:    p.Address,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}