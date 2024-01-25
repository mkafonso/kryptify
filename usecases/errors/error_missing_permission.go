package usecases

type ErrorMissingPermission struct {
	Message string
	Action  string
}

func NewErrorMissingPermission() *ErrorMissingPermission {
	return &ErrorMissingPermission{
		Message: "missing permission",
		Action:  "you don't have the required permission to perform this action",
	}
}

func (e *ErrorMissingPermission) Error() string {
	return e.Message
}

func (e *ErrorMissingPermission) GetAction() string {
	return e.Action
}
