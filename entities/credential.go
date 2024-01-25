package entities

import (
	"kryptify/entities/validations"
	valueobjects "kryptify/entities/value-objects"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID        uuid.UUID
	Email     string
	Website   string
	Category  string
	OwnerID   string
	Password  valueobjects.Password
	Health    valueobjects.Health
	CreatedAt time.Time
}

func NewCredentialt(email, password, website, owner_id string) (*Credential, error) {
	if err := validations.ValidateCreateCredentialInput(email, password, website, owner_id); err != nil {
		return nil, err
	}

	hashedPassword := valueobjects.NewPassword(password)
	passwordHealth := valueobjects.NewPasswordHealth(password)

	credential := &Credential{
		ID:        uuid.New(),
		Email:     email,
		Website:   website,
		OwnerID:   owner_id,
		Health:    passwordHealth,
		Password:  hashedPassword,
		CreatedAt: time.Now().UTC(),
	}

	return credential, nil
}
