package main

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router"
)

func main() {
	r := gin.Default()
	router.MainRouter(r)
	r.Run()
}
