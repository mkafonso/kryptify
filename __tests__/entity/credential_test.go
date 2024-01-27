package entity_test

import (
	"kryptify/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialEntity_TestCreateNewCredential(t *testing.T) {
	credential, err := entity.NewCredential("john@email.com", "test1234", "https://my-website.com", "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
	assert.NoError(t, err)
	assert.NotNil(t, credential)
}

func TestCredentialEntity_TestEmptyEmail(t *testing.T) {
	_, err := entity.NewCredential("", "test1234", "https://my-website.com", "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestCredentialEntity_TestEmptyPassword(t *testing.T) {
	_, err := entity.NewCredential("john@email.com", "", "https://my-website.com", "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestCredentialEntity_TestEmptyWebsite(t *testing.T) {
	_, err := entity.NewCredential("john@email.com", "test1234", "", "c92fdcdb-8e4b-4b0a-865c-bbc646a467a5")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestCredentialEntity_TestEmptyOwnerID(t *testing.T) {
	_, err := entity.NewCredential("john@email.com", "test1234", "https://my-website.com", "")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}
