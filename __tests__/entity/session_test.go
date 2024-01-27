package entity_test

import (
	"kryptify/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSessionEntity_TestCreateNewSession(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := "valid-refresh-token"
	userAgent := "valid-user-agent"
	clientIP := "valid-client-ip"
	isBlocked := false
	expiresAt := time.Now().Add(time.Hour)

	session, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.NoError(t, err)
	assert.NotNil(t, session)
}

func TestSessionEntity_TestBlockedAccount(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := "valid-refresh-token"
	userAgent := "valid-user-agent"
	clientIP := "valid-client-ip"
	isBlocked := true
	expiresAt := time.Now().Add(time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "cannot create a session with a blocked account")
}

func TestSessionEntity_TestInvalidExpiresAtDate(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := "valid-refresh-token"
	userAgent := "valid-user-agent"
	clientIP := "valid-client-ip"
	isBlocked := false
	expiresAt := time.Now().Add(-time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid time value")
}

func TestSessionEntity_TestEmptyAccountID(t *testing.T) {
	accountID := ""
	refreshToken := "valid-refresh-token"
	userAgent := "valid-user-agent"
	clientIP := "valid-client-ip"
	isBlocked := false
	expiresAt := time.Now().Add(time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestSessionEntity_TestEmptyRefreshToken(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := ""
	userAgent := "valid-user-agent"
	clientIP := "valid-client-ip"
	isBlocked := false
	expiresAt := time.Now().Add(time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestSessionEntity_TestEmptyUserAgent(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := "valid-refresh-token"
	userAgent := ""
	clientIP := "valid-client-ip"
	isBlocked := false
	expiresAt := time.Now().Add(time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}

func TestSessionEntity_TestEmptyClientIP(t *testing.T) {
	accountID := "valid-account-id"
	refreshToken := "valid-refresh-token"
	userAgent := "valid-user-agent"
	clientIP := ""
	isBlocked := false
	expiresAt := time.Now().Add(time.Hour)

	_, err := entity.NewSession(accountID, refreshToken, userAgent, clientIP, isBlocked, expiresAt)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing required fields in the JSON object")
}
