package errors

type ErrorCode string

type ComponentError struct {
	message   string    `json:"message"`
	errorCode ErrorCode `json:"errorCode"`
}

func Error(message string, errorCode ErrorCode) error {
	return &ComponentError{
		message:   message,
		errorCode: errorCode,
	}
}

func (c ComponentError) Error() string {
	return c.message
}
