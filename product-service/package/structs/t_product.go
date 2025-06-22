package structs

import "time"

type Product struct {
	ID       int64    	`json:"id"`
	Name     string 	`json:"name"`
	Quantity int64    	`json:"quantity"`
	Price    int64    	`json:"price"`
	CreatedAt time.Time	`json:"created_at"`
}