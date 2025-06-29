package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"log/slog"
)

type UserDB struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewUserDB(log *slog.Logger, db *sqlx.DB) *UserDB {
	return &UserDB{
		log: log,
		db:  db,
	}
}

func (u *UserDB) GetUser(id uuid.UUID) (models.User, error) {
	return models.User{}, nil
}
