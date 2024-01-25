package usecases

type ErrorMissingFields struct {
	Message string
	Action  string
}

func NewErrorMissingFields() *ErrorMissingFields {
	return &ErrorMissingFields{
		Message: "missing required fields in the JSON object",
		Action:  "please ensure you include all the required fields",
	}
}

func (e *ErrorMissingFields) Error() string {
	return e.Message
}

func (e *ErrorMissingFields) GetAction() string {
	return e.Action
}
