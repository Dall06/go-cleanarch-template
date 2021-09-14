package database

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

type Mock struct {
	DB   *sql.DB
	Sqlmock sqlmock.Sqlmock
}

func NewMock() *Mock {
	db, m, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return &Mock{
		DB:   db,
		Sqlmock: m,
	}
}
