package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/duiztapi/mocks"
	"github.com/vmkevv/duiztapi/serverr"
)

// UserServices contains all user rest services
type UserServices struct {
	actions   mocks.UserActions
	validator *validator.Validate
}

// SetupUserServices create a new instace of UserServices
func SetupUserServices(actions mocks.UserActions, validator *validator.Validate) UserServices {
	return UserServices{
		actions:   actions,
		validator: validator,
	}
}

// ServeRoutes serve the routes defined
func (us UserServices) ServeRoutes(app fiber.Router) {
	app.Post("/user", us.register())
	app.Post("/email", us.sendEmail())
}

func (us UserServices) register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqData := RegisterReq{}
		if err := c.BodyParser(&reqData); err != nil {
			return serverr.New500("Can't read name and email from body", err)
		}

		err := us.validator.Struct(reqData)
		if err != nil {
			return serverr.New(fiber.StatusBadRequest, "name must be at least 2 characters lenght and you need to type a valid email")
		}

		if us.actions.ExistsEmail(reqData.Email) {
			return serverr.New(fiber.StatusBadRequest, "There is already an account with this email.")
		}

		savedUser, err := us.actions.Register(reqData.Name, reqData.Email)
		if err != nil {
			return serverr.New500("There was a problem saving user", err)
		}

		tokenStr, err := us.actions.GenerateToken(savedUser.ID)
		if err != nil {
			return serverr.New500("Error generating token", err)
		}

		return c.JSON(RegisterRes{
			User:  *savedUser,
			Token: tokenStr,
		})
	}
}

func (us UserServices) sendEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqData := SendEmailReq{}
		if err := c.BodyParser(&reqData); err != nil {
			return serverr.New500("Can't read email from body", err)
		}

		if !us.actions.ExistsEmail(reqData.Email) {
			return serverr.New(fiber.StatusBadRequest, "The provided email does not exists in database")
		}

		err := us.actions.SendEmailToken(reqData.Email)
		if err != nil {
			return serverr.New500("Error while sending email", err)
		}

		return c.JSON(SendEmailRes{
			Message: "Magic link have been send to email address",
		})
	}
}
