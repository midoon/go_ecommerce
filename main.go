package main

import (
	"log"

	"github.com/midoo/go_ecommerce/configs"
	"github.com/midoo/go_ecommerce/internal/api"
)

func main() {
	cfg, err := configs.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
