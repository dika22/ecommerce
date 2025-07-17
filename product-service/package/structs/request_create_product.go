package structs

type RequestCreateProduct struct {
	Name               string `json:"name"`
	Quantity           int64  `json:"quantity"`
	Price              int64  `json:"price"`
	Image              string `json:"image"`
	CategoryID         int64  `json:"category_id"`
	SampleImageProduct string `json:"sample_image_product"`
}