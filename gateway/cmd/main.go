package main

import (
	"gateway/internal/auth"
	"gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// Middleware for JWT validation
	router.Use(auth.JWTMiddleware())

	// Routes
	router.GET("/currency", proxy.GetRates)
	router.GET("/currency/history", proxy.GetHistory)

	log.Println("Gateway running on :8086")
	router.Run(":8086")
}
