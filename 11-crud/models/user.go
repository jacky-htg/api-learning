package models

import (
	"github.com/jmoiron/sqlx"
)

//User : struct of User
type User struct {
	ID       uint64 `db:"id"`
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

//Get : get user by id
func (u *User) Get(db *sqlx.DB) error {
	return db.Get(u, qUsers+" WHERE id=?", u.ID)
}

//Create new user
func (u *User) Create(db *sqlx.DB) error {
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
func (u *User) Update(db *sqlx.DB) error {

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
func (u *User) Delete(db *sqlx.DB) (bool, error) {
	stmt, err := db.Prepare(`DELETE FROM users WHERE id = ?`)
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
