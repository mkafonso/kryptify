package appError

type ErrorTokenInvalid struct {
	Message string
	Action  string
}

func NewErrorTokenInvalid() *ErrorTokenInvalid {
	return &ErrorTokenInvalid{
		Message: "invalid token",
		Action:  "please provide a valid access token",
	}
}

func (e *ErrorTokenInvalid) Error() string {
	return e.Message
}

func (e *ErrorTokenInvalid) GetAction() string {
	return e.Action
}
