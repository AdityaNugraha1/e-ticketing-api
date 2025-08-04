package main

import (
	"e-ticketing-api/config"
	"e-ticketing-api/internal/database"
	"e-ticketing-api/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	cfg := config.AppConfig

	database.ConnectDB()
	defer database.DB.Close()

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", handlers.LoginHandler)

		authorized := api.Group("/")
		authorized.Use(handlers.AuthMiddleware())
		{
			authorized.POST("/terminals", handlers.CreateTerminalHandler)
		}
	}
	log.Printf("Server berjalan di port %s", cfg.ServerPort)
	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}