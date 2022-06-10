package web

type OrdersCreateRequest struct {
	// OrderDate   time.Time `validate:"required" json:"order_date"`
	CustomerId  int `validate:"required" json:"customer_id"`
	TotalAmount int `validate:"required" json:"total_amount"`
}
