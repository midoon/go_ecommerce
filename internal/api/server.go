package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoo/go_ecommerce/configs"
	"github.com/midoo/go_ecommerce/internal/api/rest"
	"github.com/midoo/go_ecommerce/internal/api/rest/handlers"
	"github.com/midoo/go_ecommerce/internal/connection"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	db := connection.GetDbConnection(config)

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRouters(rh)

	app.Listen(config.ServerPort)
}

func setupRouters(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)

	//catalog handler
	//transaction handler
}
