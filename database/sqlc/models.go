// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	AvatarUrl         sql.NullString `json:"avatar_url"`
	IsAccountVerified sql.NullBool   `json:"is_account_verified"`
	PasswordHash      string         `json:"password_hash"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

type Credential struct {
	ID           uuid.UUID      `json:"id"`
	Email        string         `json:"email"`
	Website      string         `json:"website"`
	Category     sql.NullString `json:"category"`
	OwnerID      uuid.UUID      `json:"owner_id"`
	PasswordHash string         `json:"password_hash"`
	Health       int16          `json:"health"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID      `json:"id"`
	AccountID    string         `json:"account_id"`
	RefreshToken string         `json:"refresh_token"`
	UserAgent    sql.NullString `json:"user_agent"`
	ClientIp     sql.NullString `json:"client_ip"`
	IsBlocked    sql.NullBool   `json:"is_blocked"`
	ExpiresAt    time.Time      `json:"expires_at"`
	CreatedAt    time.Time      `json:"created_at"`
}
