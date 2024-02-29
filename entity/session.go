package entity

import (
	"kryptify/val"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID
	AccountID    string
	RefreshToken string
	UserAgent    string
	ClientIP     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

func NewSession(accountID, refreshToken, userAgent, clientIP string, isBlocked bool, expiresAt time.Time) (*Session, error) {
	if err := val.CreateSessionInput(
		accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt,
	); err != nil {
		return nil, err
	}

	session := &Session{
		ID:           uuid.New(),
		AccountID:    accountID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		ClientIP:     clientIP,
		IsBlocked:    isBlocked,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now().UTC(),
	}

	return session, nil
}
