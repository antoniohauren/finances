package services

import (
	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/utils"
	"github.com/google/uuid"
)

func (s *Services) CreateUser(newUser models.CreateUserRequest) (uuid.UUID, error) {
	hashedPassword, err := utils.HashPassword(newUser.Password)

	if err != nil {
		return uuid.Nil, err
	}

	dto := models.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hashedPassword,
	}

	id, err := s.repos.User.CreateUser(dto)

	if err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}
