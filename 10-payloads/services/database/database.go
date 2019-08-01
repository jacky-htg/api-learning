package database

import (
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"
)

//Open : open database
func Open() (*sql.DB, error) {
	return sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
}

//Openx : open database
func Openx() (*sqlx.DB, error) {
	return sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
}
