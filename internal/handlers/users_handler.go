package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/antoniohauren/finances/internal/models"
)

func (h Handlers) registerUsersEndpoints() {
	http.HandleFunc("POST /auth/signup", h.createAuthSignUpEndpoint)
	http.HandleFunc("POST /auth/signin", h.createAuthSigninEndpoint)
	http.HandleFunc("POST /auth/confirm-user/{code}", h.createConfirmUserEndpoint)
	http.HandleFunc("POST /verify", h.createVerifyEndpoint)
}

func (h Handlers) createAuthSignUpEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Bad Request"})
		return
	}

	id, jwtToken, err := h.services.CreateUser(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	if jwtToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{
		UserID:      id,
		AccessToken: jwtToken,
	})
}

func (h Handlers) createAuthSigninEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.AuthSignInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Bad Request"})
		return
	}

	id, jwtToken, err := h.services.SignIn(req)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Unauthorized"})
		return
	}

	if jwtToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.AuthSignInResponse{
		UserID:      id,
		AccessToken: jwtToken,
	})
}

func (h Handlers) createConfirmUserEndpoint(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "missing Authorization header"})
		return
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid Authorization header"})
		return
	}

	token := parts[1]

	code := r.PathValue("code")
	err := h.services.ConfirmUser(token, code)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		slog.Error("confirm-user", "error", err.Error())
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Unauthorized"})
		return
	}

	json.NewEncoder(w).Encode(models.SuccessResponse{Message: "User confirmed Successfully"})
}

func (h Handlers) createVerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	type DTO struct {
		Token string `json:"token"`
	}

	var dto DTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "Bad Request"})
		return
	}

	token := dto.Token

	valid, err := h.services.VerifyUser(token)

	if err != nil || !valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{
			"valid": false,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"valid": true,
	})
}
