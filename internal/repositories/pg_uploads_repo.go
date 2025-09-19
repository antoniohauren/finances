package repositories

import (
	"database/sql"
	"log/slog"

	"github.com/antoniohauren/finances/internal/models"
)

type PgUploadRepo struct {
	db *sql.DB
}

func NewPgUploadRepo(db *sql.DB) *PgUploadRepo {
	return &PgUploadRepo{
		db: db,
	}
}

func (r *PgUploadRepo) UploadFile(upload models.Upload) (string, error) {
	var id string

	query := `
		INSERT INTO uploads (bucket_name, file_key, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(query,
		upload.BucketName,
		upload.Key,
		upload.UserID,
	).Scan(&id)

	if err != nil {
		slog.Error("create upload", "error", err.Error())
		return "", err
	}

	return id, nil
}
