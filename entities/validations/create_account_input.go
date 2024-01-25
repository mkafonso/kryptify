package validations

import (
	appError "kryptify/usecases/errors"
)

func ValidateCreateAccountInput(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return appError.NewErrorMissingFields()
	}

	if len(password) < 8 {
		return appError.NewErrorPasswordLength()
	}

	return nil
}
