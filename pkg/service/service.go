package service

import (
	"github.com/google/uuid"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
)

type User interface {
	GetUser(id uuid.UUID) (models.User, error)
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
	}
}
