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

	//email := "fahrianafdholi077gmail.com"
	//password := "12345"
	//
	//err := make(chan error)
	//
	//go helper.ValidationEmail(email, err)
	//go helper.ValidationPassword(password, err)
	//
	//fmt.Println(<-err)
}
