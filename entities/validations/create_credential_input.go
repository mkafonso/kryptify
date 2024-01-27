package validations

import (
	appError "kryptify/usecase/errors"
)

func ValidateCreateCredentialInput(email, password, website, owner_id string) error {
	if email == "" || password == "" || website == "" || owner_id == "" {
		return appError.NewErrorMissingFields()
	}

	return nil
}
