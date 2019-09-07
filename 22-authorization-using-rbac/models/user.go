package models

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/jacky-htg/go-services/libraries/array"
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
func (u *User) Create(ctx context.Context, tx *sqlx.Tx) error {
	const query = `
		INSERT INTO users (username, password, email, is_active, created, updated)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	stmt, err := tx.PreparexContext(ctx, query)
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

	if len(u.Roles) > 0 {
		for _, r := range u.Roles {
			err = u.AddRole(ctx, tx, r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//Update : update user
func (u *User) Update(ctx context.Context, tx *sqlx.Tx) error {

	stmt, err := tx.PreparexContext(ctx, `
		UPDATE users 
		SET username = ?,
			password = ?,
			is_active = ?,
			updated = NOW()
		WHERE id = ?
	`)

	if err != nil {
		println("error query update users")
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Username, u.Password, u.IsActive, u.ID)
	if err != nil {
		return err
	}

	existingRoles, err := u.GetRoleIDs(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, r := range u.Roles {
		isExist, _ := array.InArray(r.ID, existingRoles)
		if !isExist {
			err = u.AddRole(ctx, tx, r)
			if err != nil {
				return err
			}
		} else {
			temp := array.Remove(existingRoles, r.ID)
			existingRoles = temp.([]uint32)
		}
	}

	for _, r := range existingRoles {
		err = u.DeleteRole(ctx, tx, r)
		if err != nil {
			return err
		}
	}

	return nil
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

// GetRoleIDs : get array of role id from user
func (u *User) GetRoleIDs(ctx context.Context, tx *sqlx.Tx) ([]uint32, error) {
	var roles []uint32

	rows, err := tx.QueryContext(ctx, "SELECT role_id FROM roles_users WHERE user_id = ?", u.ID)
	if err != nil {
		println("error query get role ids")
		return roles, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			return roles, err
		}
		roles = append(roles, id)
	}

	return roles, rows.Err()
}

//AddRole to user
func (u *User) AddRole(ctx context.Context, tx *sqlx.Tx, r Role) error {
	stmt, err := tx.PreparexContext(ctx, `INSERT INTO roles_users (role_id, user_id) VALUES (?, ?)`)
	if err != nil {
		println("error query add role")
		return err
	}

	_, err = stmt.ExecContext(ctx, r.ID, u.ID)
	return err
}

//DeleteRole from user
func (u *User) DeleteRole(ctx context.Context, tx *sqlx.Tx, roleID uint32) error {
	stmt, err := tx.PreparexContext(ctx, `DELETE FROM roles_users WHERE role_id=? AND user_id=?`)
	if err != nil {
		println("error query delete role")
		return err
	}

	_, err = stmt.ExecContext(ctx, roleID, u.ID)
	return err
}
