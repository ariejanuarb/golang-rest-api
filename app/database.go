package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_rest_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
