package mocks

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

type Mock struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func NewMock() *Mock {
	db, m, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return &Mock{
		db:   db,
		mock: m,
	}
}
