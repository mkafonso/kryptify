package utils

import (
	"kryptify/token"
	"time"

	"github.com/google/uuid"
)

func GenerateRefreshToken(accountID uuid.UUID) (string, *token.Payload, error) {
	tokenSymmetricKey := "12345678123456781234567812345678"
	tokenMaker, err := token.NewPasetoMaker(tokenSymmetricKey)
	if err != nil {
		return "", nil, err
	}

	expirationDuration, err := time.ParseDuration("24h")
	if err != nil {
		return "", nil, err
	}

	accessToken, accessTokenPayload, err := tokenMaker.CreateToken(accountID, expirationDuration)
	if err != nil {
		return "", nil, err
	}

	return accessToken, accessTokenPayload, nil
}
