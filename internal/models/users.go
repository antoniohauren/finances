package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}
