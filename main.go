package main

import (
	"github.com/jessehorne/whousesgo.com/database"
	"github.com/jessehorne/whousesgo.com/util"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	_ "runtime/debug"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if err := database.InitDB(); err != nil {
		panic(err)
	}

	howOften := os.Getenv("ADMIN_TOKEN_UPDATE_INTERVAL")
	howOftenInt, err := strconv.Atoi(howOften)
	if err != nil {
		panic(err)
	}

	util.StartUpdateAdminTokenJob(time.Duration(howOftenInt) * time.Second)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Use(gin.Recovery())
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	// WEB ROUTES

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "WhoUsesGo.com | A thorough list of companies that use Golang",
		})
	})

	// API ROUTES

	api := r.Group("/api")
	api.GET("/ping", routes.GetPing)
	api.POST("/company", routes.PostCompany)

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
