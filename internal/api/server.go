package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoo/go_ecommerce/configs"
	"github.com/midoo/go_ecommerce/internal/api/rest"
	"github.com/midoo/go_ecommerce/internal/api/rest/handlers"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
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
