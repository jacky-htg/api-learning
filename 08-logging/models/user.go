package models

import "github.com/jmoiron/sqlx"

//User : struct of User
type User struct {
	ID       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	IsActive bool   `db:"is_active"`
}

const qUsers = `SELECT id, username, password, email, is_active FROM users`

//List : List of users
func (u *User) List(db *sqlx.DB) ([]User, error) {
	list := []User{}
	err := db.Select(&list, qUsers)
	return list, err
}
