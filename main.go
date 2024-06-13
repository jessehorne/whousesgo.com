package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/endpoints"
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

	// WEB ROUTES

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":          "WhoUsesGo.com | A thorough list of companies that use Golang",
			"companies":      []util.Company{allCompanies[0]}, // needed due to listjs needing a template
			"companiesCount": 1,
		})
	})

	// END WEB ROUTES

	// API ROUTES

	r.GET("/api/companies", endpoints.GetCompanyList)

	// END API ROUTES

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		panic(err)
	}
}
