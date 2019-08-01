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

//Create new user
func (u *User) Create(db *sql.DB) error {
	const query = `
		INSERT INTO users (username, password, email, is_active, created, updated)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Username, u.Password, u.Email, u.IsActive)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint64(id)

	return nil
}

//Update : update user
func (u *User) Update(db *sql.DB) error {

	stmt, err := db.Prepare(`
		UPDATE users 
		SET username = ?,
			password = ?,
			is_active = ?,
			updated = NOW()
		WHERE id = ?
	`)
	_, err = stmt.Exec(u.Username, u.Password, u.IsActive, u.ID)
	return err
}

//Delete : delete user
func (u *User) Delete(db *sql.DB) (bool, error) {
	stmt, err := db.Prepare(`DELETE FROM users WHERE id = ?`)
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
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
