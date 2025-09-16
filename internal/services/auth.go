package services

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/antoniohauren/finances/internal/auth"
	"github.com/antoniohauren/finances/internal/models"
	"github.com/google/uuid"
)

func (s *Services) VerifyUser(token string) (bool, error) {
	claims, err := s.jwtToken.VerifyToken(token)

	if err != nil {
		return false, err
	}

	if claims == nil {
		return false, fmt.Errorf("invalid claims found")
	}

	return true, nil
}

func (s *Services) SignIn(dto models.AuthSignInRequest) (uuid.UUID, string, error) {
	user, err := s.repos.User.GetUserByEmail(dto.Email)

	if err != nil {
		slog.Error("signin", "error", err)
		return uuid.Nil, "", err
	}

	uid, err := uuid.Parse(user.ID.String())

	if err != nil {
		return uuid.Nil, "", err
	}

	if err := auth.CheckPassword(dto.Password, user.Password); err != nil {
		return uuid.Nil, "", err
	}

	accessToken, _, err := s.jwtToken.CreateToken(uid, user.Email, 15*time.Minute)

	if err != nil {
		return uuid.Nil, "", err
	}

	return uid, accessToken, nil
}
