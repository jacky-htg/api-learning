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
			url := routings[1][2:(len(routings[1]) - 1)]
			//store access except login route
			if !(url == "/login") {
				urls := strings.Split(url, "/")
				controller := urls[1]
				access := httpMethod + " " + url
				existingAccess, err = storeAccess(existingAccess, tx, controller, access)
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

func storeAccess(existingAccess []uint32, tx *sqlx.Tx, controller string, access string) ([]uint32, error) {
	ctx := context.Background()
	// get or store parent access
	existingAccess, id, err := storeController(existingAccess, ctx, tx, controller)
	if err != nil {
		return existingAccess, err
	}

	u := models.Access{ParentID: id, Name: access}
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
	u := models.Access{Name: controller}
	err := u.GetByName(ctx, tx)
	if err != sql.ErrNoRows && err != nil {
		return existingAccess, 0, err
	}

	if err == sql.ErrNoRows {
		u.ParentID = 1
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
