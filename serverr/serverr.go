package serverr

import "github.com/gofiber/fiber/v2"

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
