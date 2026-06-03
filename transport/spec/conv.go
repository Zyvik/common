package spec

import "github.com/Zyvik/common/errors"

// ErrToErrorResponse converts error to ErrorResponse
func ErrToErrorResponse(err error) (resp *ErrorResponse, status int) {
	if se, ok := err.(errors.ServerError); ok {
		return &ErrorResponse{
			ErrorCode:    se.ErrorCode,
			ErrorMessage: se.ErrorMessage,
		}, se.HttpStatusCode
	}
	return &ErrorResponse{ErrorCode: "UNKNOWN", ErrorMessage: err.Error()}, 500
}
