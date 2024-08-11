package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"log/slog"
)

type User interface {
	GetUser(id uuid.UUID) (models.User, error)
}

type Auth interface {
	SignUp(email, username, authMethod string, passwordHash []byte) (uuid.UUID, error)
	LogIn(email string) (models.User, error)
}

type Repository struct {
	log *slog.Logger
	User
	Auth
}

func NewRepository(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{log: log, User: NewUserDB(log, db), Auth: NewAuthDB(log, db)}
}
