package database

import (
	"database/sql"

	"github.com/Dall06/go-cleanarch-template/config"
	_ "github.com/go-sql-driver/mysql"
)

type mySQLConn struct {}

func NewMySQLConn() *mySQLConn{
	return &mySQLConn{}
}

func (m *mySQLConn) OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}
	/* err = db.Ping(); if err != nil return nil, err*/
	return db, nil
}

func (m *mySQLConn) CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}
