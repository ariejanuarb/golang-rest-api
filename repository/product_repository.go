package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error)
	//FindById(ctx context.Context, tx *sql.Tx, productId int) ([]web.ProductResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse
}
