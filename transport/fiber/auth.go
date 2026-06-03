package fiber

import (
	"github.com/Zyvik/common/auth"
	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
)

func GetUserClaims(c fiber.Ctx) (*auth.UserClaims, error) {
	token := jwtware.FromContext(c)
	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		return nil, ErrorMissingTokenClaims
	}
	return claims, nil
}
