package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/internal/storage"
	"github.com/google/uuid"
)

func (h Handlers) registerPaymentsEndpoints() {
	http.HandleFunc("POST /payments", h.createPaymentEndpoint)
	http.HandleFunc("GET /payments", h.getAllPaymentsEndpoint)
	http.HandleFunc("GET /payments/{payment_id}", h.getPaymentByIdEndpoint)
	http.HandleFunc("GET /payments/bill/{bill_id}", h.getPaymentByBillEndpoint)
	http.HandleFunc("POST /payments/upload-receipt", h.uploadReceiptEndpoint)
}

func (h Handlers) createPaymentEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.CreatePaymentRequest

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

	if !user.IsVerified {
		respondJSONError(w, http.StatusUnauthorized, "please confirm your email")
		return
	}

	req.UserID = user.ID

	id, err := h.services.CreatePayment(req)

	if err != nil {
		slog.Error("create-payment", "error", err.Error())
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, http.StatusCreated, models.CreatePaymentResponse{
		PaymentID: id,
	})
}

func (h Handlers) getPaymentByIdEndpoint(w http.ResponseWriter, r *http.Request) {
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

	id := r.PathValue("payment_id")

	uid, err := uuid.Parse(id)

	if err != nil {
		slog.Error("payment-id", "error", err.Error())
		respondJSONError(w, http.StatusBadRequest, "BadRequest")
		return
	}

	body, err := h.services.GetPaymentByID(user.ID, uid)

	if err != nil {
		slog.Error("get-payment-by-id", "error", err.Error())
		respondJSONError(w, http.StatusNotFound, "Payment Not Found")
		return
	}

	respondJSON(w, http.StatusOK, body)
}

func (h Handlers) getAllPaymentsEndpoint(w http.ResponseWriter, r *http.Request) {
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

	items := h.services.GetAllPayments(user.ID)

	respondJSON(w, http.StatusOK, models.GetAllPaymentsResponse{
		Items: items,
	})
}

func (h Handlers) getPaymentByBillEndpoint(w http.ResponseWriter, r *http.Request) {
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

	id := r.PathValue("bill_id")

	billID, err := uuid.Parse(id)

	if err != nil {
		slog.Error("payment-id", "error", err.Error())
		respondJSONError(w, http.StatusBadRequest, "BadRequest")
		return
	}

	items := h.services.GetAllPaymentsByBill(user.ID, billID)

	respondJSON(w, http.StatusOK, models.GetAllPaymentsResponse{
		Items: items,
	})
}

func (h Handlers) uploadReceiptEndpoint(w http.ResponseWriter, r *http.Request) {
	// 10 MB
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		slog.Error("file size", "error", err)
		respondJSONError(w, http.StatusBadRequest, "max file size is 10 MB")
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		respondJSONError(w, http.StatusBadRequest, "file not found")
		return
	}

	defer file.Close()

	data, err := storage.UploadFile(file, header.Filename, "receipt", "images")

	if err != nil {
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, http.StatusCreated, data)
}
