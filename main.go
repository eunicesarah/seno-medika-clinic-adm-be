package main

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router"
)

func main() {
	r := gin.Default()
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
