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
		IsUserVerified(id uuid.UUID) bool
		ConfirmUser(email string) error
	}
	Bill interface {
		CreateBill(newBill models.Bill) (string, error)
		GetBillByID(id uuid.UUID) (*models.Bill, error)
		GetAllBills(userId uuid.UUID) ([]models.Bill, error)
	}
	Payment interface {
		CreatePayment(newPayment models.Payment) (string, error)
		GetPaymentByID(id uuid.UUID) (*models.Payment, *models.Upload, error)
		GetAllPayments(userId uuid.UUID) ([]models.Payment, error)
		GetAllPaymentsByBill(userId uuid.UUID, billID uuid.UUID) ([]models.Payment, error)
		AttatchReceipt(paymentID uuid.UUID, uploadID uuid.UUID) error
	}
	Upload interface {
		UploadFile(upload models.Upload) (string, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User:    NewPgUsersRepo(db),
		Bill:    NewPgBillRepo(db),
		Payment: NewPaymentRepo(db),
		Upload:  NewPgUploadRepo(db),
	}
}
