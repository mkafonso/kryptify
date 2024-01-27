package usecase

import (
	"fmt"
)

type ErrorAccountNotFound struct {
	Message string
	Action  string
}

func NewErrorAccountNotFound(accountIDOREmail string) *ErrorAccountNotFound {
	return &ErrorAccountNotFound{
		Message: fmt.Sprintf("account not found for: %s", accountIDOREmail),
		Action:  "please provide a valid account ID or Email",
	}
}

func (e *ErrorAccountNotFound) Error() string {
	return e.Message
}

func (e *ErrorAccountNotFound) GetAction() string {
	return e.Action
}
