package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"seno-medika.com/router"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	r.Use(cors.New(config))

	router.MainRouter(r)
	r.Run()
}
