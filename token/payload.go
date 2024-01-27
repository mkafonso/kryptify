package token

import (
	"time"

	appError "kryptify/usecase/errors"

	"github.com/google/uuid"
)

// payload represents the data stored in a token.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	AccountID uuid.UUID `json:"account_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with the provided account ID and duration.
func NewPayload(accountID uuid.UUID, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		AccountID: accountID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// valid checks if the token is still valid (not expired).
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return appError.NewErrorTokenExpired()
	}

	return nil
}
