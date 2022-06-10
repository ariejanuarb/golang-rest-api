package web

type ProductCreateRequest struct {
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Price      int    `validate:"required" json:"price"`
	CategoryId int    `validate:"required" json:"category_id"`
	//CreatedAt  time.Time `validate:"required,min=1,max=100" json:"created_at"`
	//UpdatedAt  time.Time `validate:"required,min=1,max=100" json:"updated_at"`
}
