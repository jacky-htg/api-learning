package models

import (
	"database/sql"
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

//Create new user
func (u *User) Create(db *sql.DB) (*User, error) {
	const query = `
		INSERT INTO users (username, password, email, is_active, created, updated)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		return u, err
	}

	res, err := stmt.Exec(u.Username, u.Password, u.Email, u.IsActive)
	if err != nil {
		return u, err
	}

	id, err := res.LastInsertId()

	u.ID = uint64(id)

	return u, nil
}

//Update : update user
func (u *User) Update(db *sql.DB) (*User, error) {

	stmt, err := db.Prepare(`
		UPDATE users 
		SET username = ?,
			password = ?,
			is_active = ?,
			updated = NOW()
		WHERE id = ?
	`)
	_, err = stmt.Exec(u.Username, u.Password, u.IsActive, u.ID)
	if err != nil {
		return u, err
	}

	return u, nil
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

func scanUser(rows *sql.Rows) (User, error) {

	var user User
	if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsActive); err != nil {
		return user, err
	}

	return user, nil
}
