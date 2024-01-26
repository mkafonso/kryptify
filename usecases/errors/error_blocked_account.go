package usecases

type ErrorBlockedAccount struct {
	Message string
	Action  string
}

func NewErrorBlockedAccount() *ErrorBlockedAccount {
	return &ErrorBlockedAccount{
		Message: "cannot create a session with a blocked account",
		Action:  "please unblock the account before creating a session",
	}
}

func (e *ErrorBlockedAccount) Error() string {
	return e.Message
}

func (e *ErrorBlockedAccount) GetAction() string {
	return e.Action
}
