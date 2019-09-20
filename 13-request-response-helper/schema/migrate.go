package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add users",
		Script: `
CREATE TABLE users (
	id   BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	username         CHAR(15) NOT NULL UNIQUE,
	password         varchar(255) NOT NULL,
	email     VARCHAR(255) NOT NULL UNIQUE,
	is_active TINYINT(1) NOT NULL DEFAULT '0',
	created TIMESTAMP NOT NULL DEFAULT NOW(),
	updated TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);`,
	},
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
