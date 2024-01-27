package validation

import (
	appError "kryptify/usecase/error"
)

func ValidateCreateCredentialInput(email, password, website, owner_id string) error {
	if email == "" || password == "" || website == "" || owner_id == "" {
		return appError.NewErrorMissingFields()
	}

	return nil
}
