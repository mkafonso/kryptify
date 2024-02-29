package appError

type ErrorEmailAlreadyTaken struct {
	Message string
	Action  string
}

func NewErrorEmailAlreadyTaken() *ErrorEmailAlreadyTaken {
	return &ErrorEmailAlreadyTaken{
		Message: "email already taken",
		Action:  "try using a different email address",
	}
}

func (e *ErrorEmailAlreadyTaken) Error() string {
	return e.Message
}

func (e *ErrorEmailAlreadyTaken) GetAction() string {
	return e.Action
}
