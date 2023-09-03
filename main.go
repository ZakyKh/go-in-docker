package main

import (
	"net/http"
	"os"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	appPort := os.Getenv("APP_PORT")
	router := gin.Default()

	router.GET("/healthz", HandleHealthz)
	router.POST("/echo", HandleEcho)

	router.Run(":" + appPort)
}

func HandleHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func HandleEcho(c *gin.Context) {
	req := c.Request
	defer req.Body.Close()
	reqBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Body cannot be read")
	}

	c.String(http.StatusOK, string(reqBytes))
}