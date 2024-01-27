package entity

import (
	"kryptify/entity/validation"
	valueobject "kryptify/entity/value-object"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID           uuid.UUID
	Email        string
	Website      string
	Category     string
	OwnerID      string
	PasswordHash valueobject.Password
	Health       valueobject.Health
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCredential(email, password, website, owner_id string) (*Credential, error) {
	if err := validation.ValidateCreateCredentialInput(email, password, website, owner_id); err != nil {
		return nil, err
	}

	hashedPassword := valueobject.NewPassword(password)
	passwordHealth := valueobject.NewPasswordHealth(password)

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