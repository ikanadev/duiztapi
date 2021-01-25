package services

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/duiztapi/actions"
	"github.com/vmkevv/duiztapi/ent"
	"github.com/vmkevv/duiztapi/mocks"
	"github.com/vmkevv/duiztapi/serverr"
)

// UserServices contains all user rest services
type UserServices struct {
	actions   mocks.UserActions
	validator *validator.Validate
}

// SetupUserServices create a new instace of UserServices
func SetupUserServices(ctx context.Context, client *ent.Client, validator *validator.Validate) UserServices {
	return UserServices{
		actions:   actions.SetupUserActions(ctx, client),
		validator: validator,
	}
}

// ServeRoutes serve the routes defined
func ServeRoutes(app fiber.Router) {}

func (us UserServices) register() fiber.Handler {
	type response struct {
		User  ent.User `json:"user"`
		Token string   `json:"token"`
	}
	return func(c *fiber.Ctx) error {
		reqData := struct {
			Name  string `json:"name" validate:"required,gte=2"`
			Email string `json:"email" validate:"required,email"`
		}{}
		if err := c.BodyParser(&reqData); err != nil {
			return serverr.New(fiber.StatusInternalServerError, "Can't read name and email from body")
		}

		err := us.validator.Struct(reqData)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
			return err
		}
		return nil
	}
}
