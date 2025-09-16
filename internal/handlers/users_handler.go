package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/utils"
)

func (h Handlers) registerUsersEndpoints() {
	http.HandleFunc("POST /auth/signup", h.createAuthSignUpEndpoint)
	http.HandleFunc("POST /auth/signin", h.createAuthSigninEndpoint)
	http.HandleFunc("POST /verify", h.createVerifyEndpoint)
}

func (h Handlers) createAuthSignUpEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "Bad Request"})
		return
	}

	id, jwtToken, err := h.services.CreateUser(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: err.Error()})
		return
	}

	if jwtToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "something went wrong"})
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
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "Bad Request"})
		return
	}

	id, jwtToken, err := h.services.SignIn(req)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "Unauthorized"})
		return
	}

	if jwtToken == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "something went wrong"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.AuthSignInResponse{
		UserID:      id,
		AccessToken: jwtToken,
	})
}

func (h Handlers) createVerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	type DTO struct {
		Token string `json:"token"`
	}

	var dto DTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "Bad Request"})
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
