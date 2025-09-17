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

func (r *PgBillRepo) GetAllBills(userID uuid.UUID) ([]models.Bill, error) {
	query := `
		SELECT id, name, due_date, type, category, frequency, payment_method
		FROM bills
		WHERE user_id = $1
		ORDER BY due_date ASC;
	`

	rows, err := r.db.Query(query, userID)

	if err != nil {
		slog.Error("get all bills", "error", err.Error())
		return nil, err
	}
	defer rows.Close()

	var bills []models.Bill

	for rows.Next() {
		var b models.Bill

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.DueDate,
			&b.Type,
			&b.Category,
			&b.Frequency,
			&b.PaymentMethod,
		)

		if err != nil {
			slog.Error("scan bill", "error", err.Error())
			return nil, err
		}

		b.UserID = userID
		bills = append(bills, b)
	}

	return bills, nil
}
