package services

import (
	"github.com/antoniohauren/finances/internal/auth"
	"github.com/antoniohauren/finances/internal/repositories"
)

type Services struct {
	repos    *repositories.Repositories
	jwtToken *auth.Jwt
}

func New(repos *repositories.Repositories, secretKey string) *Services {
	return &Services{
		repos:    repos,
		jwtToken: auth.NewJwt(secretKey),
	}
}
