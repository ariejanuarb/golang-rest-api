package web

type OrdersUpdateRequest struct {
	Id int `validate:"required"`
	//OrderDate   time.Time `validate:"required" json:"order_date"`
	CustomerId  int `validate:"required" json:"customer_id"`
	TotalAmount int `validate:"required" json:"total_amount"`
	//CreatedAt  time.Time `validate:"required"`
	//UpdatedAt  time.Time `validate:"required"`
}
