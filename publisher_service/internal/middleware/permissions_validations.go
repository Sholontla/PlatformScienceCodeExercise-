package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (pub SecurityService) ValidateRolesAndPermissions(roles []string, permissions []string) func(*fiber.Ctx) error {
	secretKey := pub.conf.KeyConfig()
	claims := ClaimsWithScope{}
	roles = []string{"user"}
	return func(c *fiber.Ctx) error {
		c.Request().Header.Set("authorization", c.Cookies("jwt"))
		//c.Set("authorization", c.Cookies("jwt"))

		// Check if the "Authorization" header is present
		authorizationHeader := c.Get("authorization")
		if authorizationHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing authorization header. Please include the 'Authorization' header with a valid token"})
		}

		// Parse the JWT token
		token, err := jwt.ParseWithClaims(c.Cookies("jwt"), &claims, func(token *jwt.Token) (interface{}, error) {
			// Replace with your own key or key provider
			return []byte(secretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization token"})
		}

		// Check if the user has the required roles
		claims, ok := token.Claims.(*ClaimsWithScope)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization claims"})
		}
		for _, role := range roles {
			if !contains(claims.Roles, role) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "user does not have the required role"})
			}
		}

		// Check if the user has the required permissions
		for _, permission := range permissions {
			if !contains(claims.Permissions, permission) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "user does not have the required permission"})
			}
		}

		// If the user has the required roles and permissions, continue to the next middleware or handler
		return c.Next()
	}
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
