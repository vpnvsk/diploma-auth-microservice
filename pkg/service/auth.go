package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/vpnvsk/amunet_auth_microservices"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type AuthService struct {
	log      *slog.Logger
	repo     repository.Auth
	settings *amunet_auth_microservices.Config
}

func NewAuthService(log *slog.Logger, repo repository.Auth, settings *amunet_auth_microservices.Config) *AuthService {
	return &AuthService{log: log, repo: repo, settings: settings}
}

func (a *AuthService) SignUp(email, username, authMethod, password string) (string, string, error) {
	op := "service.auth.SignUp"
	log := a.log.With(slog.String("op", op))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	id, err := a.repo.SignUp(email, username, authMethod, hashedPassword)
	if err != nil {
		log.Error("failed to save user")
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	accessToken, err := a.generateToken(id, time.Duration(a.settings.AccessTTL), a.settings.AccessSecret)
	if err != nil {
		log.Error("failed to generate access token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	refreshToken, err := a.generateToken(id, time.Duration(a.settings.RefreshTTL), a.settings.RefreshSecret)
	if err != nil {
		log.Error("failed to generate access token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	return accessToken, refreshToken, nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

func (a *AuthService) generateToken(userId uuid.UUID, tokenTTL time.Duration, secret string) (string, error) {
	const op = "service.generateToken"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix()},
		userId,
	})
	return token.SignedString([]byte(secret))
}

func (a *AuthService) LogIn(email, password string) (string, string, error) {
	op := "service.auth.LogIn"
	log := a.log.With(slog.String("op", op))
	user, err := a.repo.LogIn(email)
	if err != nil {
		log.Error("failed to find user", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Error("wrong password")
			return "", "", errors.New("bad credentials")
		}
	}
	accessToken, err := a.generateToken(user.Id, time.Duration(a.settings.AccessTTL), a.settings.AccessSecret)
	if err != nil {
		log.Error("failed to generate access token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	refreshToken, err := a.generateToken(user.Id, time.Duration(a.settings.RefreshTTL), a.settings.RefreshSecret)
	if err != nil {
		log.Error("failed to generate access token", err)
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	return accessToken, refreshToken, nil
}
