package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/utils"
)

func (h Handlers) registerUsersEndpoints() {
	http.HandleFunc("POST /user", h.createUserEndpoint)
}

func (h Handlers) createUserEndpoint(w http.ResponseWriter, r *http.Request) {
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
