package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (h Handlers) registerBillsEndpoints() {
	http.HandleFunc("POST /bills", h.createBillEndpoint)
	http.HandleFunc("GET /bills", h.getAllBillsEndpoint)
	http.HandleFunc("GET /bills/{bill_id}", h.getBillByIdEndpoint)
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

func (h Handlers) getBillByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	user, err := h.ExtractUserFromToken(w, r)

	if err != nil {
		slog.Error("extract-user", "error", err.Error())
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id := r.PathValue("bill_id")

	uid, err := uuid.Parse(id)

	if err != nil {
		slog.Error("bill-id", "error", err.Error())
		respondJSONError(w, http.StatusBadRequest, "BadRequest")
		return
	}

	bill, err := h.services.GetBillByID(user.ID, uid)

	if err != nil {
		slog.Error("get-bill-by-id", "error", err.Error())
		respondJSONError(w, http.StatusNotFound, "Bill Not Found")
		return
	}

	respondJSON(w, http.StatusOK, bill)
}
