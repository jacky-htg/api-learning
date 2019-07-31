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
	list := []User{}

	rows, err := db.Query(qUsers)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return list, err
		}

		list = append(list, user)
	}

	if err := rows.Err(); err != nil {
		return list, err
	}

	return list, nil
}

//Get : get user by id
func (u *User) Get(db *sql.DB, id int64) (User, error) {
	var user User

	rows, err := db.Query(qUsers+" WHERE id=?", id)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		user, err = scanUser(rows)
		if err != nil {
			return user, err
		}
	}

	if err := rows.Err(); err != nil {
		return user, err
	}

	return user, nil
}

func scanUser(rows *sql.Rows) (User, error) {

	var user User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsActive); err != nil {
		return user, err
	}

	return user, nil
}
