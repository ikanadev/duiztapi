package services_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vmkevv/duiztapi/serverr"
	"github.com/vmkevv/duiztapi/services"
)

func TestHeaders(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	app.Use(services.WithHeaders)
	app.Get("/", func(c *fiber.Ctx) error { return c.JSON("hello") })
	t.Run("Should set the correct Headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		resp, _ := app.Test(req)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	})
}

func TestAuthMiddleware(t *testing.T) {
	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	app.Use(services.WithAuth)
	app.Get("/", func(c *fiber.Ctx) error {
		tokenStr := fmt.Sprintf("%v", c.Locals("token"))
		return c.SendString(tokenStr)
	})

	t.Run("Should set ID as context value", func(t *testing.T) {
		token, _ := services.GenerateToken("secret", 1)
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", token)

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		respString, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, token, string(respString))
	})
	t.Run("Should return 401 error when there is no token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("Should return 401 error when token is empty", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

}
