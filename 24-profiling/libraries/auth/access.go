package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jacky-htg/go-services/libraries/array"
	"github.com/jacky-htg/go-services/models"
	"github.com/jmoiron/sqlx"
)

func ScanAccess(db *sqlx.DB) error {
	var existingAccess []uint32
	var err error

	// get existing access
	{
		a := models.Access{}
		existingAccess, err = a.GetIDs(context.Background(), db)
	}

	if err != nil {
		return err
	}

	// read routing file
	data, err := ioutil.ReadFile("routing/route.go")
	if err != nil {
		return err
	}

	// set transaction
	tx := db.MustBegin()

	// convert routing to access field
	datas := strings.Split(string(data), "\n")
	for _, env := range datas {
		env = strings.TrimSpace(env)
		if len(env) > 11 && env[:11] == "app.Handle(" {
			routings := strings.Split(env[11:(len(env)-1)], ",")
			httpMethod := routings[0][11:]
			url := strings.TrimSpace(routings[1])
			url = url[1:(len(url) - 1)]
			alias := strings.TrimSpace(routings[2])

			//store access except login route
			if !(url == "/login") {
				urls := strings.Split(url, "/")
				controller := urls[1]
				access := strings.ToUpper(httpMethod) + " " + url
				existingAccess, err = storeAccess(existingAccess, tx, controller, access, alias)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	// remove existing access
	err = removeAccess(tx, existingAccess)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func storeAccess(existingAccess []uint32, tx *sqlx.Tx, controller string, access string, alias string) ([]uint32, error) {
	ctx := context.Background()
	// get or store parent access
	existingAccess, id, err := storeController(existingAccess, ctx, tx, controller)
	if err != nil {
		return existingAccess, err
	}
	nullID := sql.NullInt64{Int64: int64(id), Valid: true}

	u := models.Access{ParentID: nullID, Name: access, Alias: alias}
	err = u.GetByName(ctx, tx)
	if err != sql.ErrNoRows && err != nil {
		return existingAccess, err
	}

	if err == sql.ErrNoRows {
		err = u.Create(ctx, tx)
		if err != nil {
			return existingAccess, err
		}
		println("store " + u.Name)
	} else {
		arr := array.Remove(existingAccess, u.ID)
		existingAccess = arr.([]uint32)
	}

	return existingAccess, nil
}

func storeController(existingAccess []uint32, ctx context.Context, tx *sqlx.Tx, controller string) ([]uint32, uint32, error) {
	u := models.Access{Name: controller, Alias: controller}
	err := u.GetByName(ctx, tx)
	if err != sql.ErrNoRows && err != nil {
		return existingAccess, 0, err
	}

	if err == sql.ErrNoRows {
		u.ParentID = sql.NullInt64{Int64: 1, Valid: true}
		err = u.Create(ctx, tx)
		if err != nil {
			return existingAccess, 0, err
		}
		println("store " + u.Name)
	} else {
		arr := array.Remove(existingAccess, u.ID)
		existingAccess = arr.([]uint32)
	}

	return existingAccess, u.ID, nil
}

func removeAccess(tx *sqlx.Tx, existingAccess []uint32) error {
	var err error
	ctx := context.Background()

	for _, i := range existingAccess {
		var isSuccess bool
		u := models.Access{ID: i}
		err = u.Get(ctx, tx)
		if err != nil {
			return err
		}

		name := u.Name

		isSuccess, err = u.Delete(ctx, tx)
		if err != nil {
			return err
		}

		if !isSuccess {
			return errors.New("Deleted failed")
		}

		fmt.Println("Deleted " + name)
	}

	return err
}
