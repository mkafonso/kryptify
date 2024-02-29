package token

import (
	"time"

	"github.com/google/uuid"
)

func GenerateAccessToken(accountID uuid.UUID) (string, *Payload, error) {
	tokenSymmetricKey := "12345678123456781234567812345678"
	tokenMaker, err := NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		return "", nil, err
	}

	expirationDuration, err := time.ParseDuration("15m")
	if err != nil {
		return "", nil, err
	}

	accessToken, accessTokenPayload, err := tokenMaker.CreateToken(accountID, expirationDuration)
	if err != nil {
		return "", nil, err
	}

	return accessToken, accessTokenPayload, nil
}
