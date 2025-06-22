package structs


type ResponseLogin struct {
	Token string `json:"token"`
	Message string `json:"message"`
}