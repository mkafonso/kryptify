package entities

import (
	"kryptify/entities/validations"
	valueobjects "kryptify/entities/value-objects"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID           uuid.UUID
	Email        string
	Website      string
	Category     string
	OwnerID      string
	PasswordHash valueobjects.Password
	Health       valueobjects.Health
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCredential(email, password, website, owner_id string) (*Credential, error) {
	if err := validations.ValidateCreateCredentialInput(email, password, website, owner_id); err != nil {
		return nil, err
	}

	hashedPassword := valueobjects.NewPassword(password)
	passwordHealth := valueobjects.NewPasswordHealth(password)

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
