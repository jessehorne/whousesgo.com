package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/util"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	_ "runtime/debug"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	allCompanies, err := util.GetAllCompanies()
	if err != nil {
		panic(err)
	}
	companiesCount := len(allCompanies)

	// WEB ROUTES

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":          "WhoUsesGo.com | A thorough list of companies that use Golang",
			"companies":      allCompanies,
			"companiesCount": companiesCount,
		})
	})

	// END WEB ROUTES

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
