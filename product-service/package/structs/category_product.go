package structs

type CategoryProduct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (cp RequestCreateCategoryProduct) NewCategoryProduct() CategoryProduct {
	return CategoryProduct{
		Name: cp.Name,
	}
}