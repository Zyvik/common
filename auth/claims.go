package auth

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.Claims
	UserID    int      `json:"userID"`
	UserRoles []string `json:"userRoles"`
}

func (u *UserClaims) ValidateRole(role string) bool {
	for _, userRole := range u.UserRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
