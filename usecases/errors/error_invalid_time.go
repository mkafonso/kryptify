package usecases

type ErrorInvalidTime struct {
	Message string
	Action  string
}

func NewErrorInvalidTime() *ErrorInvalidTime {
	return &ErrorInvalidTime{
		Message: "invalid time value",
		Action:  "please provide a valid time",
	}
}

func (e *ErrorInvalidTime) Error() string {
	return e.Message
}

func (e *ErrorInvalidTime) GetAction() string {
	return e.Action
}
