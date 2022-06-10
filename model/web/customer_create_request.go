package web

type CustomerCreateRequest struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Address     string `validate:"required,min=1,max=200" json:"address"`
	Email       string `validate:"required,min=1,max=100" json:"email"`
	PhoneNumber string `validate:"required,min=1,max=100" json:"phone_number"`
}
