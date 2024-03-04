package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Printf("")
		c.JSON(http.StatusOK, gin.H{
			"message": "pksoonnifm lkdfm ld",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:jdsvnks8080 (for windows "localhost:8080")
}
