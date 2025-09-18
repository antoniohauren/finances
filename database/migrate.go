package database

import (
	"database/sql"

	"github.com/antoniohauren/finances/database/migrations"
)

func MigrateAll(db *sql.DB) error {
	if err := migrations.MigrateUser(db); err != nil {
		return err
	}

	if err := migrations.MigrateBill(db); err != nil {
		return err
	}

	if err := migrations.MigratePayment(db); err != nil {
		return err
	}

	return nil
}
