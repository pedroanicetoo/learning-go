package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro_r := sql.Open("mysql", stringConnection)
	if erro_r != nil {
		return nil, erro_r
	}

	if erro_r = db.Ping(); erro_r != nil {
		return nil, erro_r
	}

	return db, nil
}
