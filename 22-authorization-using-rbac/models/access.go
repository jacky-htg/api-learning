package models

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//Access : struct of Access
type Access struct {
	ID       uint32        `db:"id"`
	ParentID sql.NullInt64 `db:"parent_id"`
	Name     string        `db:"name"`
}

const qAccess = `SELECT id, parent_id, name FROM access`

//List : List of access
func (u *Access) List(ctx context.Context, tx *sqlx.Tx) ([]Access, error) {
	list := []Access{}
	err := tx.SelectContext(ctx, &list, qAccess)
	return list, err
}

//GetByName : get access by name
func (u *Access) GetByName(ctx context.Context, tx *sqlx.Tx) error {
	return tx.GetContext(ctx, u, qAccess+" WHERE name=?", u.Name)
}

//Get : get access by id
func (u *Access) Get(ctx context.Context, tx *sqlx.Tx) error {
	return tx.GetContext(ctx, u, qAccess+" WHERE id=?", u.ID)
}

//Create new Access
func (u *Access) Create(ctx context.Context, tx *sqlx.Tx) error {
	const query = `
		INSERT INTO access (parent_id, name, created)
		VALUES (?, ?, NOW())
	`
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.ParentID, u.Name)
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

//Delete : delete user
func (u *Access) Delete(ctx context.Context, tx *sqlx.Tx) (bool, error) {
	stmt, err := tx.PreparexContext(ctx, `DELETE FROM access WHERE id = ?`)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *Access) GetIDs(ctx context.Context, db *sqlx.DB) ([]uint32, error) {
	var access []uint32

	rows, err := db.Query("SELECT id FROM access WHERE name != 'root'")
	if err != nil {
		return access, err
	}

	for rows.Next() {
		var id uint32
		err = rows.Scan(&id)
		if err != nil {
			return access, err
		}
		access = append(access, id)
	}

	return access, rows.Err()
}
