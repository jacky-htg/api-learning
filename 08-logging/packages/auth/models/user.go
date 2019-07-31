package models

import (
	"database/sql"
)

//User : struct of User
type User struct {
	ID       uint
	Username string
	Password string
	Email    string
	IsActive bool
}

const qUsers = `SELECT id, username, password, email, is_active FROM users`

//List : List of users
func (u *User) List(db *sql.DB) ([]User, error) {
	return scanUsers(qUsers, db)
}

func scanUsers(q string, db *sql.DB) ([]User, error) {
	list := []User{}
	rows, err := db.Query(q)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsActive); err != nil {
			return list, err
		}
		list = append(list, user)
	}

	if err := rows.Err(); err != nil {
		return list, err
	}

	return list, nil
}
