package migrations

import (
	"database/sql"
	"fmt"
)

func MigratePayment(db *sql.DB) error {
	paymentsTable := `
		CREATE TABLE IF NOT EXISTS payments (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			date TIMESTAMP NOT NULL,
			amount NUMERIC(10,2) NOT NULL,
    	method VARCHAR(50) NOT NULL,
			receipt_id UUID REFERENCES uploads(id),
			bill_id UUID NOT NULL REFERENCES bills(id) ON DELETE CASCADE,
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (now()),
			updated_at TIMESTAMPTZ
		);
	`

	if _, err := db.Exec(paymentsTable); err != nil {
		return fmt.Errorf("creating bills table: %w", err)
	}

	return nil
}
