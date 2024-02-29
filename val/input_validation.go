package val

import (
	appError "kryptify/usecase/error"
	"time"
)

func CreateAccountInput(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return appError.NewErrorMissingFields()
	}

	if len(password) < 8 {
		return appError.NewErrorPasswordLength()
	}

	return nil
}

func CreateCredentialInput(email, password, website, owner_id string) error {
	if email == "" || password == "" || website == "" || owner_id == "" {
		return appError.NewErrorMissingFields()
	}

	return nil
}

func CreateSessionInput(
	accountID, refreshToken, userAgent, clientIP string, isBlocked bool, expiresAt time.Time,
) error {
	if accountID == "" || refreshToken == "" || userAgent == "" || clientIP == "" {
		return appError.NewErrorMissingFields()
	}

	if isBlocked {
		return appError.NewErrorBlockedAccount()
	}

	if expiresAt.IsZero() || expiresAt.Before(time.Now()) {
		return appError.NewErrorInvalidTime()
	}

	return nil
}
