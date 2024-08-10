package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (u *UserDB) GetUser(id uuid.UUID) (models.User, error) {
	return models.User{}, nil
}
