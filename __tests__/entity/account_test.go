package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountEntity_TestCreateNewAccount(t *testing.T) {
	account, err := entity.NewAccount("John Doe", "john@email.com", "test1234")
	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestAccountEntity_TestEmptyName(t *testing.T) {
	_, err := entity.NewAccount("", "john@email.com", "test1234")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestAccountEntity_TestEmptyEmail(t *testing.T) {
	_, err := entity.NewAccount("John Doe", "", "test1234")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestAccountEntity_TestEmptyPassword(t *testing.T) {
	_, err := entity.NewAccount("John Doe", "john@email.com", "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestAccountEntity_TestShortPassword(t *testing.T) {
	_, err := entity.NewAccount("John Doe", "john@email.com", "test")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "password length must be at least 8 characters")
}
