package serverr

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// ServErr represents a server error
type ServErr struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// Error to implement golang error interface
func (s ServErr) Error() string {
	return s.Message
}

// New creates a new ServErr
func New(code int, message string) ServErr {
	return ServErr{code, message}
}

// New500 creates a new ServErr with 500 Internal server error code
func New500(message string, err ...error) ServErr {
	if len(err) == 0 {
		return ServErr{
			Code:    fiber.StatusInternalServerError,
			Message: message,
		}
	}
	return ServErr{
		fiber.StatusInternalServerError,
		fmt.Sprintf("%s\nError %v", message, err[0]),
	}
}

// Handler custom error handler for fiber
func Handler(c *fiber.Ctx, e error) error {
	servErr, ok := e.(ServErr)
	if ok {
		c.Status(servErr.Code)
		return c.JSON(servErr)
	}
	if e != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}
