package validation

import (
	appError "kryptify/usecase/error"
	"time"
)

func ValidateCreateSessionInput(
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
