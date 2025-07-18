package structs

type RequestUpdateSeller struct {
	Id         int64  `json:"id"`
	NameSeller string `json:"name_seller"`
	Address    string `json:"address"`
	Status     bool   `json:"status"`
	IsPremium  bool   `json:"is_premium"`
}