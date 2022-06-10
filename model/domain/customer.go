package domain

import "time"

type Customer struct {
	Id          int
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
