package services

import (
	"time"

	"github.com/antoniohauren/finances/internal/auth"
	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (s *Services) CreateUser(newUser models.CreateUserRequest) (uuid.UUID, string, error) {
	hashedPassword, err := auth.HashPassword(newUser.Password)

	if err != nil {
		return uuid.Nil, "", err
	}

	dto := models.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hashedPassword,
	}

	id, err := s.repos.User.CreateUser(dto)

	if err != nil {
		return uuid.Nil, "", err
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil, "", err
	}

	accessToken, _, err := s.jwtToken.CreateToken(uid, newUser.Email, 15*time.Minute)

	if err != nil {
		return uuid.Nil, "", err
	}

	return uid, accessToken, nil
}
