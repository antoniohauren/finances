package repositories

import (
	"database/sql"
	"log/slog"

	"github.com/antoniohauren/finances/internal/models"
)

type PgUsersRepo struct {
	db *sql.DB
}

func NewPgUsersRepo(db *sql.DB) *PgUsersRepo {
	return &PgUsersRepo{
		db: db,
	}
}

func (r *PgUsersRepo) CreateUser(newUser models.User) (string, error) {
	var id string

	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		newUser.Name,
		newUser.Email,
		newUser.Password,
	).Scan(&id)

	if err != nil {
		slog.Error("create user", "error", err.Error())
		return "", err
	}

	return id, nil
}
