package models

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

//User : struct of User
type User struct {
	ID       uint64 `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	IsActive bool   `db:"is_active"`
	Roles    []Role `db:"roles"`
}

const qUsers = `SELECT id, username, password, email, is_active FROM users`

//List : List of users
func (u *User) List(ctx context.Context, db *sqlx.DB) ([]User, error) {
	qUsers := `
	SELECT users.id, users.username, users.password, users.email, users.is_active, 
		JSON_ARRAYAGG(roles.id) as roles_id, JSON_ARRAYAGG(roles.name) as roles_name
	FROM users
	JOIN roles_users ON users.id=roles_users.user_id
	JOIN roles ON roles_users.role_id=roles.id
	GROUP BY users.id
	`
	list := []User{}
	rows, err := db.Query(qUsers)
	if err != nil {
		return list, err
	}

	defer rows.Close()

	for rows.Next() {
		var u User
		var roles_id string
		var roles_name string
		err = rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.IsActive, &roles_id, &roles_name)
		if err != nil {
			return list, err
		}

		var ids []int32
		err = json.Unmarshal([]byte(roles_id), &ids)
		if err != nil {
			return list, err
		}
		var names []string
		err = json.Unmarshal([]byte(roles_name), &names)
		if err != nil {
			return list, err
		}

		for i, v := range ids {
			u.Roles = append(u.Roles, Role{ID: uint32(v), Name: names[i]})
		}

		list = append(list, u)
	}

	return list, rows.Err()
}

//Get : get user by id
func (u *User) Get(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qUsers+" WHERE id=?", u.ID)
}

//GetByEmail : get user by email
func (u *User) GetByEmail(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qUsers+" WHERE email=?", u.Email)
}

//GetByUsername : get user by username
func (u *User) GetByUsername(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qUsers+" WHERE username=?", u.Username)
}

//Create new user
func (u *User) Create(ctx context.Context, db *sqlx.DB) error {
	const query = `
		INSERT INTO users (username, password, email, is_active, created, updated)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.Username, u.Password, u.Email, u.IsActive)
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
func (u *User) Update(ctx context.Context, db *sqlx.DB) error {

	stmt, err := db.PreparexContext(ctx, `
		UPDATE users 
		SET username = ?,
			password = ?,
			is_active = ?,
			updated = NOW()
		WHERE id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Username, u.Password, u.IsActive, u.ID)
	return err
}

//Delete : delete user
func (u *User) Delete(ctx context.Context, db *sqlx.DB) (bool, error) {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM users WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
