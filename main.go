package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/env", func(c *gin.Context) {
		env := os.Environ()
		c.JSON(http.StatusOK, gin.H{
			"env": env,
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
