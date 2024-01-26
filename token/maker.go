package token

import (
	"time"

	"github.com/google/uuid"
)

type MakerInterface interface {
	CreateToken(accountID uuid.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
