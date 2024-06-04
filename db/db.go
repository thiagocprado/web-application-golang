package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	s := "user=postgres dbname=golang_loja password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", s)

	if err != nil {
		panic(err.Error())
	}

	return db
}
