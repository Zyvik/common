package fiber

import (
	"github.com/Zyvik/common/transport/spec"
	"github.com/gofiber/fiber/v3"
)

// ServerErrorHandler is a fiber error handler that always returns ErrorResponse
func ServerErrorHandler(c fiber.Ctx, err error) error {
	resp, statusCode := spec.ErrToErrorResponse(err)
	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}
	return c.Status(statusCode).JSON(resp)
}
