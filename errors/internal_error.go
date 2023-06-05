package errors

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return "Error: " + e.Message
}

func GenerateInternalError(message string) error {
	return &InternalError{
		Message: message,
	}
}
