package entities

import (
	"kryptify/entities/validations"
	valueobjects "kryptify/entities/value-objects"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID
	Name              string
	Email             string
	AvatarUrl         string
	IsAccountVerified bool
	PasswordHash      valueobjects.Password
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewAccount(name, email, password string) (*Account, error) {
	if err := validations.ValidateCreateAccountInput(name, email, password); err != nil {
		return nil, err
	}

	hashedPassword := valueobjects.NewPassword(password)

	account := &Account{
		ID:                uuid.New(),
		Name:              name,
		Email:             email,
		IsAccountVerified: false,
		PasswordHash:      hashedPassword,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}

	return account, nil
}
