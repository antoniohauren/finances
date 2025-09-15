package services

import "github.com/antoniohauren/finances/internal/repositories"

type Services struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *Services {
	return &Services{
		repos: repos,
	}
}
