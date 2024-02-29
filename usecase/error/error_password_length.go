package appError

type ErrorPasswordLength struct {
	Message string
	Action  string
}

func NewErrorPasswordLength() *ErrorPasswordLength {
	return &ErrorPasswordLength{
		Message: "password length must be at least 8 characters",
		Action:  "please choose a longer password",
	}
}

func (e *ErrorPasswordLength) Error() string {
	return e.Message
}

func (e *ErrorPasswordLength) GetAction() string {
	return e.Action
}
