package entity

import (
	"kryptify/entity/validation"
	valueobject "kryptify/entity/value-object"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID
	Name              string
	Email             string
	AvatarUrl         string
	IsAccountVerified bool
	PasswordHash      valueobject.Password
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewAccount(name, email, password string) (*Account, error) {
	if err := validation.ValidateCreateAccountInput(name, email, password); err != nil {
		return nil, err
	}

	hashedPassword := valueobject.NewPassword(password)

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
