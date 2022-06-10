package web

type OrdersResponse struct {
	Id int `json:"id"`
	// OrderDate   time.Time `json:"order_date"`
	CustomerId  int `json:"customer_id"`
	TotalAmount int `json:"total_amount"`
}
