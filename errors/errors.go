package errors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ServerError struct {
	ErrorCode      string
	ErrorMessage   string
	HttpStatusCode int
}

func (se ServerError) Error() string {
	return fmt.Sprintf("%s: %s", se.ErrorCode, se.ErrorMessage)
}

// ValidationErrosToErrorResponse converts ValidationErrors into the ServerError
func ValidationErrosToServerError(valErr validator.ValidationErrors, traceID string) *ServerError {
	errorMessage := "Request validation failed because of the following fields: "
	fieldErrMsgs := make([]string, len(valErr))
	for i, err := range valErr { // TODO - add more info
		// Gets rid of the topmost namespace. Eg. CreateUserReq.email -> email
		sturctNamespace := strings.SplitN(err.StructNamespace(), ".", 2)[0] + "."
		fieldErrMsgs[i] = strings.TrimLeft(err.Namespace(), sturctNamespace)
	}
	errorMessage += strings.Join(fieldErrMsgs, ", ")

	return &ServerError{
		ErrorCode:    "AUTH-VALIDATION",
		ErrorMessage: errorMessage,
	}
}
