package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	BaseEntity
	Name     string
	Email    string
	Password string
	Code     sql.NullString
	Bills    []Bill
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	UserID       uuid.UUID `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type VerifyUserResponse struct {
	IsValid bool `json:"is_valid"`
}

type ConfirmUserResponse struct {
	Token string `json:"token"`
}
