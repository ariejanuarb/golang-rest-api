package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required,max=100,min=1" json:"name"`
	Price      int    `validate:"required" json:"price"`
	CategoryId int    `validate:"required" json:"category_id"`
	//CreatedAt  time.Time `validate:"required"`
	//UpdatedAt  time.Time `validate:"required"`
}
