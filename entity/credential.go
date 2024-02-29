package entity

import (
	"kryptify/util"
	"kryptify/val"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID           uuid.UUID
	Email        string
	Website      string
	Category     string
	OwnerID      string
	PasswordHash util.Password
	Health       util.Health
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCredential(email, password, website, owner_id string) (*Credential, error) {
	if err := val.CreateCredentialInput(email, password, website, owner_id); err != nil {
		return nil, err
	}

	hashedPassword := util.NewPassword(password)
	passwordHealth := util.NewPasswordHealth(password)

	credential := &Credential{
		ID:           uuid.New(),
		Email:        email,
		Website:      website,
		OwnerID:      owner_id,
		Health:       passwordHealth,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	return credential, nil
}
