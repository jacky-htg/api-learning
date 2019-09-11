package models

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//Role : struct of Role
type Role struct {
	ID   uint32 `db:"id"`
	Name string `db:"name"`
}

const qRoles = `SELECT id, name FROM roles`

//List of roles
func (u *Role) List(ctx context.Context, db *sqlx.DB) ([]Role, error) {
	list := []Role{}
	err := db.SelectContext(ctx, &list, qRoles)
	return list, err
}

//Get role by id
func (u *Role) Get(ctx context.Context, db *sqlx.DB) error {
	return db.GetContext(ctx, u, qRoles+" WHERE id=?", u.ID)
}

//Create new role
func (u *Role) Create(ctx context.Context, db *sqlx.DB) error {
	const query = `
		INSERT INTO roles (name, created)
		VALUES (?, NOW())
	`
	stmt, err := db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.Name)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint32(id)

	return nil
}

//Update role
func (u *Role) Update(ctx context.Context, db *sqlx.DB) error {

	stmt, err := db.PreparexContext(ctx, `
		UPDATE roles 
		SET name = ?
		WHERE id = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Name, u.ID)
	return err
}

//Delete role
func (u *Role) Delete(ctx context.Context, db *sqlx.DB) (bool, error) {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM roles WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

//Grant access to role
func (u *Role) Grant(ctx context.Context, db *sqlx.DB, accessID uint32) error {
	stmt, err := db.PreparexContext(ctx, `INSERT INTO access_roles (access_id, role_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, accessID, u.ID)
	return err
}

//Revoke access from role
func (u *Role) Revoke(ctx context.Context, db *sqlx.DB, accessID uint32) error {
	stmt, err := db.PreparexContext(ctx, `DELETE FROM access_roles WHERE access_id= ? AND role_id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, accessID, u.ID)
	return err
}
