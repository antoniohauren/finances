package services

import (
	"github.com/antoniohauren/finances/internal/repositories"
	"github.com/antoniohauren/finances/utils"
)

type Services struct {
	repos    *repositories.Repositories
	jwtToken *utils.Jwt
}

func New(repos *repositories.Repositories, secretKey string) *Services {
	return &Services{
		repos:    repos,
		jwtToken: utils.NewJwt(secretKey),
	}
}
