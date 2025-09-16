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

func (r *PgUsersRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := `
		SELECT id, name, email, password_hash
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		slog.Error("get by email", "error", err.Error())
		return nil, err
	}

	return &user, nil
}
