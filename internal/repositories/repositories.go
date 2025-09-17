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
		GetBillById(id uuid.UUID) (*models.Bill, error)
		GetAllBills() []models.Bill
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: NewPgUsersRepo(db),
		Bill: NewPgBillRepo(db),
	}
}
