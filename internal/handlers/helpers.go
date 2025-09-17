package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/antoniohauren/finances/internal/models"
)

// func respondJSON(w http.ResponseWriter, status int, message any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(message)
// }

func respondJSONError(w http.ResponseWriter, status int, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ErrorResponse{Reason: reason})
}

func (h Handlers) ExtractUserFromToken(w http.ResponseWriter, r *http.Request) (*models.UserClaims, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		slog.Error("auth-header", "error", "missing auth header")
		return nil, fmt.Errorf("invalid Authorization header")
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		slog.Error("bearer", "error", "invalid auth header")
		return nil, fmt.Errorf("invalid Authorization header")
	}

	token := parts[1]

	user, err := h.services.GetUserFromToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
