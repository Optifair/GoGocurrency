package proxy

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func GetRates(c *gin.Context) {
	resp, _ := http.Get("http://currency:8087/rates")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

func GetHistory(c *gin.Context) {
	resp, _ := http.Get("http://currency:8087/history")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}
