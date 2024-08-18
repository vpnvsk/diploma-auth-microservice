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

func (a *AuthDB) SignUp(tx *sqlx.Tx, email, username, authMethod string, passwordHash []byte) (uuid.UUID, error) {
	op := "repository.auth.SignIn"
	var id uuid.UUID

	log := a.log.With(slog.String("op", op))
	query := fmt.Sprintf(`INSERT INTO "%s" (email, username, auth_method, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`, userTable)
	row := tx.QueryRow(query, email, username, authMethod, passwordHash)
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

func (a *AuthDB) LogIn(email string) (models.UserGet, error) {
	op := "repository.LogIn"
	var user models.UserGet
	log := a.log.With(slog.String("op", op))

	query := fmt.Sprintf(`SELECT id, email, password_hash FROM "%s" WHERE email=$1`, userTable)
	if err := a.db.Get(&user, query, email); err != nil {
		log.Info("user not found", err)
		return user, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (a *AuthDB) Transactional(txFunc func(tx *sqlx.Tx) error) error {
	tx, err := a.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return err
}

func (a *AuthDB) UpdateRefreshTokenTransaction(tx *sqlx.Tx, userId uuid.UUID, refreshToken []byte) error {
	op := "repository.UpdateRefreshToken"
	log := a.log.With("op", op)
	query := `UPDATE "user" SET refresh_token = $1 WHERE id = $2`
	if _, err := tx.Exec(query, refreshToken, userId); err != nil {
		log.Error("error occurred while updating token")
		return err
	}
	return nil
}

func (a *AuthDB) UpdateRefreshToken(userId uuid.UUID, refreshToken []byte) error {
	op := "repository.UpdateRefreshToken"
	log := a.log.With("op", op)
	query := `UPDATE "user" SET refresh_token = $1 WHERE id = $2`
	if _, err := a.db.Exec(query, refreshToken, userId); err != nil {
		log.Error("error occurred while updating token")
		return err
	}
	return nil
}
