package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type OrdersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders)
	FindById(ctx context.Context, tx *sql.Tx, orders int) (domain.Orders, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Orders
}
