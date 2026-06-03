package fiber

import (
	"crypto/ecdsa"
	"time"

	"github.com/Zyvik/common/auth"
	"github.com/Zyvik/common/logging"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

const (
	TraceIDHeader = "TraceID"
	SpanIDHeader  = "SpanID"
)

// NewLoggingMiddleware adds a logger and tracing data into the context for later use.
// It also logs out requests before and after processing.
func NewLoggingMiddleware(log *logrus.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := c.Context()
		ctx = logging.AddLoggerToContext(ctx, log)

		var traceID, spanID string
		if c.HasHeader(TraceIDHeader) {
			traceID = c.GetHeaders()[TraceIDHeader][0]
		}
		if c.HasHeader(SpanIDHeader) {
			spanID = c.GetHeaders()[SpanIDHeader][0]
		}
		ctx = logging.AddTracingToContext(ctx, traceID, spanID)
		c.SetContext(ctx)

		middlewareLog := logging.GetTracedEntry(ctx).WithFields(logrus.Fields{
			"method": c.Method(),
			"URL":    c.OriginalURL(),
			"IP":     c.IP(),
		})

		// initial request log
		middlewareLog.Info("Received a request.")

		t := time.Now()
		handlerErr := c.Next() // run next handler in the chain
		if handlerErr != nil {
			// run error handler
			if err := c.App().ErrorHandler(c, handlerErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
			middlewareLog = middlewareLog.WithError(handlerErr)
		}

		// log after request processing
		middlewareLog.WithFields(logrus.Fields{
			"duration":   time.Since(t).Milliseconds(),
			"statusCode": c.Response().StatusCode(),
		}).Info("Request processed.")
		return nil
	}
}

// NewJwtMiddleware returs a middleware responsible for validating JWTs
func NewJwtMiddleware(publicKey *ecdsa.PublicKey) fiber.Handler {
	cfg := jwtware.Config{
		SigningKey: jwtware.SigningKey{JWTAlg: "ES256", Key: publicKey},
		Claims:     &auth.UserClaims{},
	}
	return jwtware.New(cfg)
}

// NewDefaultContentTypeMiddleware returns a middleware that sets default content type for requests that do not have it already.
func NewDefaultContentTypeMiddleware(contentType string) fiber.Handler {
	return func(c fiber.Ctx) error {
		switch c.Method() {
		case fiber.MethodPost, fiber.MethodPut, fiber.MethodPatch:
			if c.Get(fiber.HeaderContentType) == "" {
				c.Set(fiber.HeaderContentType, contentType)
			}
			fallthrough
		default:
			return c.Next()
		}
	}
}
