package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/utils"
)

func (c Controller) registerUsersEndpoints() {
	http.HandleFunc("POST /user", c.createUserEndpoint)
}

func (c Controller) createUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: "Bad Request"})
		return
	}

	id, err := c.services.CreateUser(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{UserID: id})
}
