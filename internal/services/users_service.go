package services

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/antoniohauren/finances/internal/auth"
	"github.com/antoniohauren/finances/internal/mailer"
	"github.com/antoniohauren/finances/internal/models"
	"github.com/antoniohauren/finances/internal/utils"
	"github.com/google/uuid"
)

func (s *Services) CreateUser(newUser models.CreateUserRequest) (uuid.UUID, string, error) {
	hashedPassword, err := auth.HashPassword(newUser.Password)

	if err != nil {
		return uuid.Nil, "", err
	}

	code := utils.GenerateUserCode()

	dto := models.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Code:     sql.NullString{String: code, Valid: true},
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

	accessToken, _, err := s.jwtToken.CreateToken(uid, newUser.Email, time.Hour)

	if err != nil {
		return uuid.Nil, "", err
	}

	body := fmt.Sprintf("Code: %s", code)
	err = mailer.SendEmail([]string{newUser.Email}, "Please Confirm Your Email", body)

	// need to revert in case of error, or even delete the user
	if err != nil {
		slog.Error("failed to send email")
		return uuid.Nil, "", err
	}

	return uid, accessToken, nil
}
