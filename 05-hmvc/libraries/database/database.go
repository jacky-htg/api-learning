package database

import "database/sql"

//Open : open database
func Open() (*sql.DB, error) {
	return sql.Open("mysql", "root:pass@tcp(localhost:3306)/go-services?parseTime=true")
}
