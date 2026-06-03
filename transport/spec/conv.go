package spec

import "github.com/Zyvik/common/errors"

// ErrToErrorResponse converts error to ErrorResponse
func ErrToErrorResponse(err error) *ErrorResponse {
	if se, ok := err.(errors.ServerError); ok {
		return &ErrorResponse{
			ErrorCode:    se.ErrorCode,
			ErrorMessage: se.ErrorMessage,
		}
	}
	return &ErrorResponse{ErrorCode: "UNKNOWN", ErrorMessage: err.Error()}
}
