package errors

type InvalidArgument struct {
	Message string
}

func (e *InvalidArgument) Error() string {
	return "Invalid Argument Passed {" + e.Message + "}"
}

func GenerateInvalidArgumentError(message string) error {
	return &InvalidArgument{
		Message: message,
	}
}
