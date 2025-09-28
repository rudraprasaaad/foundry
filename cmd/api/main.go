package main

import (
	"foundry/internal/api"
	"foundry/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %V", err)
	}

	router := gin.Default()
	api.SetupRoutes(router)

	log.Printf("Starting Foundry server on port %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}
