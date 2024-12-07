package rates

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"USD": 1.0,
		"EUR": 0.85,
	})
}

func GetHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"2023-10-01": map[string]float64{"USD": 1.0, "EUR": 0.84},
		"2023-10-02": map[string]float64{"USD": 1.0, "EUR": 0.85},
	})
}
