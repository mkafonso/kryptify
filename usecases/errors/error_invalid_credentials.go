package usecases

type ErrorInvalidCredentials struct {
	Message string
	Action  string
}

func NewErrorInvalidCredentials() *ErrorInvalidCredentials {
	return &ErrorInvalidCredentials{
		Message: "invalid credentials",
		Action:  "please check your email and password and try again",
	}
}

func (e *ErrorInvalidCredentials) Error() string {
	return e.Message
}

func (e *ErrorInvalidCredentials) GetAction() string {
	return e.Action
}
