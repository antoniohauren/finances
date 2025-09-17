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
		respondJSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	id, jwtToken, err := h.services.CreateUser(req)

	if err != nil {
		respondJSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if jwtToken == "" {
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, http.StatusCreated, models.CreateUserResponse{
		UserID:      id,
		AccessToken: jwtToken,
	})
}

func (h Handlers) createAuthSigninEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.AuthSignInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	id, jwtToken, err := h.services.SignIn(req)

	if err != nil {
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if jwtToken == "" {
		respondJSONError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	respondJSON(w, http.StatusCreated, models.AuthSignInResponse{
		UserID:      id,
		AccessToken: jwtToken,
	})
}

func (h Handlers) createConfirmUserEndpoint(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		respondJSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		respondJSON(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token := parts[1]

	code := r.PathValue("code")
	err := h.services.ConfirmUser(token, code)

	if err != nil {
		slog.Error("confirm-user", "error", err.Error())
		respondJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	respondJSON(w, http.StatusOK, models.SuccessResponse{Message: "User confirmed Successfully"})
}

func (h Handlers) createVerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	type DTO struct {
		Token string `json:"token"`
	}

	var dto DTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		slog.Error("user-decoder", "error", err.Error())
		respondJSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	token := dto.Token

	valid, err := h.services.VerifyUser(token)

	if err != nil || !valid {
		respondJSON(w, http.StatusUnauthorized, models.VerifyUserResponse{
			IsValid: false,
		})
		return
	}

	respondJSON(w, http.StatusUnauthorized, models.VerifyUserResponse{
		IsValid: true,
	})
}
