package repositories

import (
	"database/sql"

	"github.com/antoniohauren/finances/internal/models"
)

type Repositories struct {
	User interface {
		CreateUser(newUser models.User) (string, error)
		GetUserByEmail(email string) (*models.User, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: NewPgUsersRepo(db),
	}
}
