package structs

type RequestCreateSeller struct {
	NameSeller string `json:"name_seller"`
	Address    string `json:"address"`
	Image      string `json:"image"`
}