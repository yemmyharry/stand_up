package main

import (
	"os"
	adapter "stand_up/internal/adapter/api/resource"
	"stand_up/internal/core/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("stand_up.env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	router := gin.Default()
	handler := adapter.NewHTTPHandler()
	handler.Routes(router)
	logger.Info("Starting server on port 8088")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}
	router.Run(":" + port)
}
