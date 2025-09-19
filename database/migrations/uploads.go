package migrations

import (
	"database/sql"
	"fmt"
)

func MigrateUploads(db *sql.DB) error {
	uploadsTable := `
		CREATE TABLE IF NOT EXISTS uploads (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			bucket_name VARCHAR(50),
			file_key VARCHAR(50),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMPTZ DEFAULT (now()),
			updated_at TIMESTAMPTZ
		);
	`

	if _, err := db.Exec(uploadsTable); err != nil {
		return fmt.Errorf("creating bills table: %w", err)
	}

	return nil
}
