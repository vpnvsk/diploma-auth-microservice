package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/vpnvsk/amunet_auth_microservices/internal/models"
	"log/slog"
)

type AuthDB struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewAuthDB(log *slog.Logger, db *sqlx.DB) *AuthDB {
	return &AuthDB{
		log: log,
		db:  db,
	}
}

func (a *AuthDB) SignUp(email, username, authMethod string, passwordHash []byte) (uuid.UUID, error) {
	op := "repository.auth.SignIn"
	var id uuid.UUID

	log := a.log.With(slog.String("op", op))
	query := fmt.Sprintf(`INSERT INTO "%s" (email, username, auth_method, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`, userTable)
	row := a.db.QueryRow(query, email, username, authMethod, passwordHash)
	if err := row.Scan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // 23505 is the PostgreSQL error code for unique violation
				a.log.Info("SignUp error: user already exists, email: %s, username: %s", email, username)
				return uuid.Nil, fmt.Errorf("%s: %w", op, ErrUserExists)
			}
		}
		log.Info("error while registering user", err)
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (a *AuthDB) LogIn(email string) (models.User, error) {
	return models.User{}, nil
}
