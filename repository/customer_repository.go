package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer)
	FindById(ctx context.Context, tx *sql.Tx, customer int) (domain.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Customer
}
