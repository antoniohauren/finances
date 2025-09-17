package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
)

func (h Handlers) registerBillsEndpoints() {
	http.HandleFunc("POST /bills", h.createBillEndpoint)
	http.HandleFunc("GET /bills", h.getAllBillsEndpoint)
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
		slog.Error("create-bill", "error", err.Error())
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, http.StatusCreated, models.CreateBillResponse{
		BillID: id,
	})
}

func (h Handlers) getAllBillsEndpoint(w http.ResponseWriter, r *http.Request) {
	user, err := h.ExtractUserFromToken(w, r)

	if err != nil {
		slog.Error("extract-user", "error", err.Error())
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	items := h.services.GetAllBills(user.ID)

	respondJSON(w, http.StatusOK, models.GetAllBillsResponse{
		Items: items,
	})
}
