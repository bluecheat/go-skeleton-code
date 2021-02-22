package errors

import "google.golang.org/grpc/codes"

type ErrorCode string

func (code ErrorCode) String() string {
	return string(code)
}

type ComponentError struct {
	Message   string     `json:"message"`
	Code      codes.Code `json:"code"`
	ErrorCode ErrorCode  `json:"errorCode"`
}

func Error(message string, errorCode ErrorCode) *ComponentError {
	return &ComponentError{
		Message:   message,
		ErrorCode: errorCode,
	}
}

func Convert(err error) *ComponentError {
	switch ve := err.(type) {
	case *ComponentError:
		if ve.Code == 0 {
			ve.Code = codes.Unknown
		}
		return ve
	}
	return &ComponentError{
		Message:   err.Error(),
		Code:      codes.Unknown,
		ErrorCode: ErrCode,
	}
}

func (c ComponentError) Error() string {
	return c.Message
}
