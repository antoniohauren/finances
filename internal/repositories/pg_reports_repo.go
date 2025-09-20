package repositories

import (
	"database/sql"
	"time"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

type PgReportRepo struct {
	db *sql.DB
}

func NewReportRepo(db *sql.DB) *PgReportRepo {
	return &PgReportRepo{
		db: db,
	}
}

func (r *PgReportRepo) GetMonthyReport(userID uuid.UUID, month time.Time) ([]models.Report, error) {
	query := `
		SELECT
			b.id AS bill_id,
			b.name AS bill_name,
			b.frequency,
			b.payment_method,
			COALESCE(SUM(p.amount), 0) AS total_amount,
			COUNT(p.id) AS payments_count
		FROM bills b
		LEFT JOIN payments p
			ON p.bill_id = b.id
			AND DATE_TRUNC('month', p.date) = DATE_TRUNC('month', $2::timestamptz)
			AND p.user_id = $1
		WHERE b.user_id = $1
		GROUP BY b.id, b.name, b.frequency, b.payment_method
		ORDER BY b.name;
	`

	rows, err := r.db.Query(query, userID, month)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []models.Report

	for rows.Next() {
		var r models.Report

		err := rows.Scan(
			&r.BillID,
			&r.BillName,
			&r.Frequency,
			&r.BillMethod,
			&r.TotalAmount,
			&r.PaymentsCount,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, r)
	}

	return reports, nil
}
