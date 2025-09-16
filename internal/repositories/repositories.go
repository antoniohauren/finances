package repositories

import (
	"database/sql"

	"github.com/antoniohauren/finances/internal/models"
)

type Repositories struct {
	User interface {
		CreateUser(newUser models.User) (string, error)
		GetUserByEmail(email string) (*models.User, error)
		ConfirmUser(email string) error
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: NewPgUsersRepo(db),
	}
}
