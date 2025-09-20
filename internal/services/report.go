package services

import (
	"time"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (s *Services) GetMonthyReport(userID uuid.UUID) ([]models.ReportItemResponse, error) {
	now := time.Now()

	data, err := s.repos.Report.GetMonthyReport(userID, now)

	if err != nil {
		return nil, err
	}

	var items []models.ReportItemResponse
	for _, r := range data {
		items = append(items, models.ReportItemResponse{
			BillID:        r.BillID,
			BillName:      r.BillName,
			Frequency:     r.Frequency,
			BillMethod:    r.BillMethod,
			TotalAmount:   r.TotalAmount,
			PaymentsCount: r.PaymentsCount,
		})
	}

	return items, nil
}
