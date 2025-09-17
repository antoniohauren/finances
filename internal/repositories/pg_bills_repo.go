package repositories

import (
	"database/sql"
	"log/slog"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

type PgBillRepo struct {
	db *sql.DB
}

func NewPgBillRepo(db *sql.DB) *PgBillRepo {
	return &PgBillRepo{
		db: db,
	}
}

func (r *PgBillRepo) CreateBill(newBill models.Bill) (string, error) {
	var id string

	query := `
		INSERT INTO bills (name, due_date, type, category, frequency, payment_method, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(query,
		newBill.Name,
		newBill.DueDate,
		newBill.Type,
		newBill.Category,
		newBill.Frequency,
		newBill.PaymentMethod,
		newBill.UserID,
	).Scan(&id)

	if err != nil {
		slog.Error("create bill", "error", err.Error())
		return "", err
	}

	return id, nil
}

func (r *PgBillRepo) GetBillById(id uuid.UUID) (*models.Bill, error) {
	return nil, nil
}

func (r *PgBillRepo) GetAllBills() []models.Bill {
	return nil
}
