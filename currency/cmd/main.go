package main

import (
	"currency/internal/rates"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/rates", rates.GetRates)
	router.GET("/history", rates.GetHistory)

	log.Println("Currency service running on :8087")
	router.Run(":8087")
}
