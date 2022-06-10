package web

type CustomerResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
