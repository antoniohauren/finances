package migrations

import (
	"database/sql"
	"fmt"
)

func MigrateBill(db *sql.DB) error {
	billsTable := `
		CREATE TABLE IF NOT EXISTS bills (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(100) NOT NULL,
			due_date TIMESTAMPTZ NOT NULL,
			type VARCHAR(50) NOT NULL,
			category VARCHAR(50) NOT NULL,
			frequency VARCHAR(50) NOT NULL,
			payment_method VARCHAR(50) NOT NULL,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (now()),
			updated_at TIMESTAMPTZ
		);
	`
	if _, err := db.Exec(billsTable); err != nil {
		return fmt.Errorf("creating bills table: %w", err)
	}

	return nil
}
