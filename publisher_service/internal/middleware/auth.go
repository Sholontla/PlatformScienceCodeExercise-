package middleware

import (
	"net/http"
	"publisher_service/config"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type SecurityService struct {
	conf config.Config
}

type ClaimsWithScope struct {
	jwt.RegisteredClaims
	Scope       string
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func (pub SecurityService) IsAuthenticated(ctx *fiber.Ctx) error {

	secretKey := pub.conf.KeyConfig()

	cookie := ctx.Cookies("jwt")
	claims := ClaimsWithScope{}
	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthenticated"})
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	isSuperAdmin := strings.Contains(ctx.Path(), "/service/publisher/admin")
	if (payLoad.Scope == "user" && isSuperAdmin) || (payLoad.Scope == "super_admin" && !isSuperAdmin) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})
	}

	isAdmin := strings.Contains(ctx.Path(), "/service/publisher")
	if (payLoad.Scope == "super_admin" && isAdmin) || (payLoad.Scope == "user" && !isAdmin) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})
	}

	return ctx.Next()
}

func (pub SecurityService) GetUserLogin(c *fiber.Ctx) (string, error) {
	secretKey := pub.conf.KeyConfig()
	cookie := c.Cookies("jwt")

	claims := ClaimsWithScope{}

	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	return payLoad.Subject, nil
}

func (pub SecurityService) GenerateJWT(adminEmail string, scope string, role []string, permissions []string) (string, error) {
	secretKey := pub.conf.KeyConfig()
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	payLoad := ClaimsWithScope{}
	payLoad.Subject = adminEmail
	payLoad.ExpiresAt = jwt.NewNumericDate(expireTime)
	payLoad.Scope = scope
	payLoad.Roles = role
	payLoad.Permissions = permissions
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad).SignedString([]byte(secretKey))

}
