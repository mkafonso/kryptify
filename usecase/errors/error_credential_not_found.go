package usecase

import (
	"fmt"
)

type ErrorCredentialNotFound struct {
	Message string
	Action  string
}

func NewErrorCredentialNotFound(credentialID string) *ErrorCredentialNotFound {
	return &ErrorCredentialNotFound{
		Message: fmt.Sprintf("credential not found for: %s", credentialID),
		Action:  "please provide a valid credential ID",
	}
}

func (e *ErrorCredentialNotFound) Error() string {
	return e.Message
}

func (e *ErrorCredentialNotFound) GetAction() string {
	return e.Action
}
