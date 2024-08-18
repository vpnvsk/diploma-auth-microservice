package service

import (
	"github.com/google/uuid"
	"github.com/vpnvsk/amunet_auth_microservices"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
	"log/slog"
)

type User interface {
	GetUser(id uuid.UUID) (models.User, error)
}

type Auth interface {
	SignUp(email, username, authMethod, password string) (string, string, error)
	LogIn(email, password string) (string, string, error)
}

type Service struct {
	log      *slog.Logger
	settings *amunet_auth_microservices.Config
	User
	Auth
}

func NewService(log *slog.Logger, repo *repository.Repository, settings *amunet_auth_microservices.Config) *Service {
	return &Service{
		log:      log,
		User:     NewUserService(log, repo.User),
		Auth:     NewAuthService(log, repo.Auth, settings),
		settings: settings,
	}
}
