package domain

import "time"

type OrderProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     int
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
