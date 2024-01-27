package valueobject_test

import (
	valueobject "kryptify/entity/value-object"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword_ShouldHashThePassword(t *testing.T) {
	password := "myVerySecurePassword"

	hashedPassword := valueobject.NewPassword(password)

	assert.NotEmpty(t, hashedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	assert.NoError(t, err)
}

func TestPassword_TestIsPasswordValidWithValidPassword(t *testing.T) {
	password := "myVerySecurePassword"
	hashedPassword := valueobject.NewPassword(password)

	isValid := hashedPassword.IsPasswordValid(password)
	assert.True(t, isValid)
}

func TestPassword_TestIsPasswordValidWithInValidPassword(t *testing.T) {
	password := "myVerySecurePassword"
	hashedPassword := valueobject.NewPassword(password)

	invalidPassword := "wrongPassword"
	isValid := hashedPassword.IsPasswordValid(invalidPassword)
	assert.False(t, isValid)
}

func TestPassword_SamePasswordsShouldProduceDifferentHashedPasswords(t *testing.T) {
	password1 := "myVerySecurePassword"
	hashedPassword1 := valueobject.NewPassword(password1)

	password2 := "myVerySecurePassword"
	hashedPassword2 := valueobject.NewPassword(password2)

	assert.NotEqual(t, hashedPassword1, hashedPassword2)
}
