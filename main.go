package main

import (
	_ "runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/routes"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.GET("/ping", routes.GetPing)

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
