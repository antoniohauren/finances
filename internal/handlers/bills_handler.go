package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
)

func (h Handlers) registerBillsEndpoints() {
	http.HandleFunc("POST /bills", h.createBillEndpoint)
}

func (h Handlers) createBillEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBillRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("body-decode", "error", err.Error())
		respondJSONError(w, http.StatusBadRequest, "Bad request")
		return
	}

	user, err := h.ExtractUserFromToken(w, r)

	if err != nil {
		slog.Error("extract-user", "error", err.Error())
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	req.UserID = user.ID

	id, err := h.services.CreateBill(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "something went wrong"})
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateBillResponse{
		BillID: id,
	})
}
