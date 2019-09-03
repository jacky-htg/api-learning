package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//Open : open database
func Open() (*sql.DB, error) {
	return sql.Open("mysql", "root:pass@tcp(localhost:3306)/go-services?parseTime=true")
}

//Openx : open database
func Openx() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "root:pass@tcp(localhost:3306)/go-services?parseTime=true")
}
