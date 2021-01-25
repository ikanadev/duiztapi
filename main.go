package main

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/vmkevv/duiztapi/actions"
	"github.com/vmkevv/duiztapi/config"
	"github.com/vmkevv/duiztapi/ent"
	"github.com/vmkevv/duiztapi/serverr"
	"github.com/vmkevv/duiztapi/services"
)

func start() error {
	config.SetEnvs()
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	client, err := ent.Open("postgres", config.PostgresConnStr())
	if err != nil {
		return err
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: serverr.Handler,
	})
	appV1 := app.Group("/api/v1")

	validator := validator.New()

	context := context.Background()

	userActions := actions.SetupUserActions(context, client)
	services.SetupUserServices(userActions, validator).ServeRoutes(appV1)

	app.Listen(":8000")
	return nil
}

func main() {
	err := start()
	if err != nil {
		log.Fatalf("Error initalizing app: %v", err)
	}
}
