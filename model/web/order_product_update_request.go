package web

type OrderProductUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required" json:"order_id"`
	ProductId int `validate:"required" json:"product_id"`
	Qty       int `validate:"required" json:"qty"`
	Price     int `validate:"required" json:"price"`
	Amount    int `validate:"required" json:"amount"`
	//CreatedAt  time.Time `validate:"required"`
	//UpdatedAt  time.Time `validate:"required"`
}
