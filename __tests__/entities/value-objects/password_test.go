package valueobjects_test

import (
	valueobjects "kryptify/entities/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword_ShouldHashThePassword(t *testing.T) {
	password := "myVerySecurePassword"

	hashedPassword := valueobjects.NewPassword(password)

	assert.NotEmpty(t, hashedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	assert.NoError(t, err)
}

func TestPassword_TestIsPasswordValidWithValidPassword(t *testing.T) {
	password := "myVerySecurePassword"
	hashedPassword := valueobjects.NewPassword(password)

	isValid := hashedPassword.IsPasswordValid(password)
	assert.True(t, isValid)
}

func TestPassword_TestIsPasswordValidWithInValidPassword(t *testing.T) {
	password := "myVerySecurePassword"
	hashedPassword := valueobjects.NewPassword(password)

	invalidPassword := "wrongPassword"
	isValid := hashedPassword.IsPasswordValid(invalidPassword)
	assert.False(t, isValid)
}

func TestPassword_SamePasswordsShouldProduceDifferentHashedPasswords(t *testing.T) {
	password1 := "myVerySecurePassword"
	hashedPassword1 := valueobjects.NewPassword(password1)

	password2 := "myVerySecurePassword"
	hashedPassword2 := valueobjects.NewPassword(password2)

	assert.NotEqual(t, hashedPassword1, hashedPassword2)
}
