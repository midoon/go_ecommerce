package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoo/go_ecommerce/configs"
)

func StartServer(config configs.AppConfig) {
	app := fiber.New()

	app.Listen(config.ServerPort)
}
