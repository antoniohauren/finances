package repositories

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

type PgPaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PgPaymentRepo {
	return &PgPaymentRepo{
		db: db,
	}
}

func (r *PgPaymentRepo) CreatePayment(newPayment models.Payment) (string, error) {
	var id string

	query := `
		INSERT INTO payments (date, amount, method, bill_id, user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	err := r.db.QueryRow(query,
		newPayment.Date,
		newPayment.Amount,
		newPayment.Method,
		newPayment.BillID,
		newPayment.UserID,
	).Scan(&id)

	if err != nil {
		slog.Error("create payment", "error", err.Error())
		return "", nil
	}

	return id, nil
}

func (r *PgPaymentRepo) GetPaymentByID(id uuid.UUID) (*models.Payment, *models.Upload, error) {
	query := `
		SELECT p.id, p.amount, p.date, p.method, p.user_id, p.bill_id,
					 u.bucket_name, u.file_key, u.user_id
		FROM payments p
		LEFT JOIN uploads u ON p.receipt_id = u.id
		WHERE p.id = $1;
	`

	var payment models.Payment
	// upload
	var bucketName sql.NullString
	var fileKey sql.NullString
	var userID sql.NullString

	err := r.db.QueryRow(query, id).Scan(
		&payment.ID,
		&payment.Amount,
		&payment.Date,
		&payment.Method,
		&payment.UserID,
		&payment.BillID,
		// upload
		&bucketName,
		&fileKey,
		&userID,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			slog.Error("get-payment-by-id", "error", err.Error())
		}

		return nil, nil, err
	}

	if bucketName.Valid && fileKey.Valid && userID.Valid {
		id, _ := uuid.Parse(userID.String)

		upload := models.Upload{
			BucketName: bucketName.String,
			Key:        fileKey.String,
			UserID:     id,
		}

		return &payment, &upload, nil
	}

	return &payment, nil, nil
}

func (r *PgPaymentRepo) GetAllPayments(userId uuid.UUID) ([]models.Payment, error) {
	query := `
		SELECT id, amount, date, method, user_id, bill_id
		FROM payments
		WHERE user_id = $1
		ORDER BY date ASC;
	`

	rows, err := r.db.Query(query, userId)

	if err != nil {
		slog.Error("get all payments", "error", err.Error())
		return nil, err
	}

	defer rows.Close()

	var payments []models.Payment

	for rows.Next() {
		var p models.Payment

		err := rows.Scan(
			&p.ID,
			&p.Amount,
			&p.Date,
			&p.Method,
			&p.UserID,
			&p.BillID,
		)

		if err != nil {
			slog.Error("scan payment", "error", err.Error())
			return nil, err
		}

		payments = append(payments, p)
	}

	return payments, nil
}

func (r *PgPaymentRepo) GetAllPaymentsByBill(userID uuid.UUID, billID uuid.UUID) ([]models.Payment, error) {
	query := `
		SELECT id, amount, date, method, user_id, bill_id
		FROM payments
		WHERE user_id = $1 AND bill_id = $2
		ORDER BY date ASC;
	`

	rows, err := r.db.Query(query, userID, billID)

	if err != nil {
		slog.Error("get all payments", "error", err.Error())
		return nil, err
	}

	defer rows.Close()

	var payments []models.Payment

	for rows.Next() {
		var p models.Payment

		err := rows.Scan(
			&p.ID,
			&p.Amount,
			&p.Date,
			&p.Method,
			&p.UserID,
			&p.BillID,
		)

		if err != nil {
			slog.Error("scan payment", "error", err.Error())
			return nil, err
		}

		payments = append(payments, p)
	}

	return payments, nil
}

func (r *PgPaymentRepo) AttatchReceipt(paymentID uuid.UUID, uploadID uuid.UUID) error {
	query := `
		UPDATE payments
		SET receipt_id = $1, updated_at = now()
		WHERE id = $2;
	`

	_, err := r.db.Exec(query, uploadID, paymentID)

	if err != nil {
		return fmt.Errorf("attach receipt: %w", err)
	}

	return nil
}
