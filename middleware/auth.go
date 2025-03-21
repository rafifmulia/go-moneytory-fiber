package middleware

import (
	"restfulapi/helper"

	"github.com/gofiber/fiber/v2"
)

// Authorization Middleware
func AuthMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey != "secret" {
		return helper.RespUnauthorized(c, "")
	}
	return c.Next()
}
