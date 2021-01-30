package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/duiztapi/serverr"
)

// WithHeaders set headers for the application
func WithHeaders(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return c.Next()
}

// WithAuth reads the Authorization token and set the ID in context
func WithAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return serverr.New(fiber.StatusUnauthorized, "Access denied")
	}
	c.Locals("token", token)
	return c.Next()
}
