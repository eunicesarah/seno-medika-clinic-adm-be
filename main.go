package main

import (
 "github.com/gin-gonic/gin"
 "seno-medika.com/router"
 "github.com/gin-contrib/cors"
 
)

func main() {
 r := gin.Default()
 config := cors.DefaultConfig()
 config.AllowOrigins = []string{"http://localhost:3000"}
 config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
 r.Use(cors.New(config))

 router.MainRouter(r)
 r.Run()
}