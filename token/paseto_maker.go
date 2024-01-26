package token

import (
	"fmt"
	appError "kryptify/usecases/errors"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// pasetoMaker implements the Maker interface using Paseto.
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker instance with the provided symmetric key.
func NewPasetoMaker(symmetricKey string) (MakerInterface, error) {
	if len(symmetricKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid secret key length: expected %d bytes, got %d", chacha20poly1305.KeySize, len(symmetricKey))
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// createToken generates a new Paseto token for the given account ID and duration.
func (m *PasetoMaker) CreateToken(accountID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	pasetoInstance := paseto.NewV2()

	payload, err := NewPayload(accountID, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := pasetoInstance.Encrypt(m.symmetricKey, payload, nil)
	return token, payload, err
}

// verifyToken verifies and decodes a Paseto token, returning the token's payload if valid.
func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	pasetoInstance := paseto.NewV2()

	payload := &Payload{}

	err := pasetoInstance.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, appError.NewErrorTokenInvalid()
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
