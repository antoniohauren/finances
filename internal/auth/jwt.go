package auth

import (
	"fmt"
	"time"

	"github.com/antoniohauren/finances/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Jwt struct {
	secretKey string
}

func NewJwt(secretKey string) *Jwt {
	return &Jwt{
		secretKey: secretKey,
	}
}

func (j *Jwt) CreateToken(id uuid.UUID, email string, duration time.Duration) (string, *models.UserClaims, error) {
	claims, err := models.NewUserClaims(id, email, duration)

	if err != nil {
		return "", nil, fmt.Errorf("failed to create claims %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign token %w", err)
	}

	return tokenStr, claims, nil
}

func (j *Jwt) VerifyToken(tokenStr string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.UserClaims{}, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parsing token: %w", err)
	}

	claims, ok := token.Claims.(*models.UserClaims)

	if !ok {
		return nil, fmt.Errorf("invalid token claims: %w", err)
	}

	return claims, nil
}
