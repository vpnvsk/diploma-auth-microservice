package service

import (
	"github.com/google/uuid"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
	"log/slog"
)

type UserService struct {
	log  *slog.Logger
	repo repository.User
}

func NewUserService(log *slog.Logger, repo repository.User) *UserService {
	return &UserService{
		log:  log,
		repo: repo,
	}
}

func (u *UserService) GetUser(id uuid.UUID) (models.User, error) {
	return models.User{}, nil
}
