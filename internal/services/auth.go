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

func (s *Services) GetUserFromToken(token string) (*models.UserClaims, error) {
	claims, err := s.jwtToken.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	if !claims.IsVerified {
		claims.IsVerified = s.repos.User.IsUserVerified(claims.ID)
	}

	return claims, err
}

func (s *Services) ConfirmUser(token string, code string) (string, error) {
	claims, err := s.jwtToken.VerifyToken(token)

	if err != nil {
		return "", err
	}

	email := claims.Email

	if email == "" {
		return "", fmt.Errorf("email not found in token")
	}

	user, err := s.repos.User.GetUserByEmail(email)

	if err != nil {
		return "", fmt.Errorf("Unauthorized")
	}

	if !user.Code.Valid || user.Code.String == "" {
		return "", fmt.Errorf("code already used")
	}

	if user.Code.String != code {
		slog.Error("Invalid confirmation code")
		return "", fmt.Errorf("Unauthorized")
	}

	err = s.repos.User.ConfirmUser(user.Email)

	if err != nil {
		return "", fmt.Errorf("Unauthorized")
	}

	isVerified := true
	accessToken, _, err := s.jwtToken.CreateToken(user.ID, user.Email, isVerified, time.Hour)

	if err != nil {
		return "", err
	}

	return accessToken, nil
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

	isVerified := user.Code.String == ""

	accessToken, _, err := s.jwtToken.CreateToken(uid, user.Email, isVerified, time.Hour)

	if err != nil {
		return uuid.Nil, "", err
	}

	return uid, accessToken, nil
}
