package service

import (
	"github.com/google/uuid"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUser(id uuid.UUID) (models.User, error) {
	return models.User{}, nil
}
