package usecase

type ErrorTokenExpired struct {
	Message string
	Action  string
}

func NewErrorTokenExpired() *ErrorTokenExpired {
	return &ErrorTokenExpired{
		Message: "expired token",
		Action:  "please generate a new token and try again",
	}
}

func (e *ErrorTokenExpired) Error() string {
	return e.Message
}

func (e *ErrorTokenExpired) GetAction() string {
	return e.Action
}
