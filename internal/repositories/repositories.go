package repositories

import (
	"database/sql"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

type Repositories struct {
	User interface {
		CreateUser(newUser models.User) (string, error)
		GetUserByEmail(email string) (*models.User, error)
		ConfirmUser(email string) error
	}
	Bill interface {
		CreateBill(newBill models.Bill) (string, error)
		GetBillByID(id uuid.UUID) (*models.Bill, error)
		GetAllBills(userId uuid.UUID) ([]models.Bill, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: NewPgUsersRepo(db),
		Bill: NewPgBillRepo(db),
	}
}
