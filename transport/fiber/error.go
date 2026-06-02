package fiber

import "common/errors"

var (
	ErrorMissingTokenClaims = errors.ServerError{
		ErrorCode:    "JWT-001",
		ErrorMessage: "Token does not contain expected claims",
	}
)
