package models

import (
	"database/sql"
	"errors"
)

//User : struct of User
type User struct {
	ID       uint64
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

	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(getArgs(&user)...)
		if err != nil {
			return list, err
		}

		list = append(list, user)
	}

	if err := rows.Err(); err != nil {
		return list, err
	}

	if len(list) <= 0 {
		return list, errors.New("Users not found")
	}

	return list, nil
}

//Get : get user by id
func (u *User) Get(db *sql.DB, id int64) error {
	return db.QueryRow(qUsers+" WHERE id=?", id).Scan(getArgs(u)...)
}

func getArgs(user *User) []interface{} {
	var args []interface{}
	args = append(args, &user.ID)
	args = append(args, &user.Username)
	args = append(args, &user.Password)
	args = append(args, &user.Email)
	args = append(args, &user.IsActive)
	return args
}
