package models

import "github.com/google/uuid"

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserGet struct {
	Id             uuid.UUID `json:"id" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	HashedPassword []byte    `json:"password_hash" db:"password_hash" binding:"required"`
}

type UserLogIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
}
