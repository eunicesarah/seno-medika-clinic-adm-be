package main

import (
	"net/http"
	"seno-medika.com/config/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		q, err := db.DB.Exec("INSERT INTO users VALUES(1,'1','1','1','1','1')")
		if err != nil {
			// Handle the error if the database operation fails
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": q,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
