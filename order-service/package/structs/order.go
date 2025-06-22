package structs

import "time"

type Order struct {
    ID        int64      `json:"id" gorm:"primaryKey"`
    Status    int64      `json:"status"`  // 1: pending, 2: paid 3: failed
    CreatedAt time.Time  `json:"created_at"`
}


func (p RequestCreateOrder) NewOrder() Order {
	return Order{
		Status:    1,
		CreatedAt: time.Now(),
	}
}