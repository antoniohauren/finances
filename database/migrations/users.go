package migrations

import (
	"database/sql"
	"fmt"
)

func MigrateUser(db *sql.DB) error {
	usersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			password_hash VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT (now()),
			updated_at TIMESTAMP
		)
	`

	if _, err := db.Exec(usersTable); err != nil {
		return fmt.Errorf("creating users table: %w", err)
	}

	return nil
}
