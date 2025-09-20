package handlers

import (
	"log/slog"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
)

func (h Handlers) registerReportEndpoints() {
	http.HandleFunc("GET /report", h.getMonthReportEndpoint)
}

func (h Handlers) getMonthReportEndpoint(w http.ResponseWriter, r *http.Request) {
	user, err := h.ExtractUserFromToken(w, r)

	if err != nil {
		slog.Error("extract-user", "error", err.Error())
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !user.IsVerified {
		respondJSONError(w, http.StatusUnauthorized, "please confirm your email")
		return
	}

	data, err := h.services.GetMonthyReport(user.ID)

	if err != nil {
		slog.Error("", "error", err)
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, 201, models.ReportResponse{
		Items: data,
	})
}
