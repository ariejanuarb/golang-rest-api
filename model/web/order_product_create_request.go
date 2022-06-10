package web

type OrderProductCreateRequest struct {
	OrderId   int `validate:"required" json:"order_id"`
	ProductId int `validate:"required" json:"product_id"`
	Qty       int `validate:"required" json:"qty"`
	Price     int `validate:"required" json:"price"`
	Amount    int `validate:"required" json:"amount"`
	//CreatedAt  time.Time `validate:"required,min=1,max=100" json:"created_at"`
	//UpdatedAt  time.Time `validate:"required,min=1,max=100" json:"updated_at"`
}
