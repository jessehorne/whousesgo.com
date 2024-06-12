package routes

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/database"
	"github.com/jessehorne/whousesgo.com/database/models"
	"github.com/jessehorne/whousesgo.com/util"
	"net/http"
)

type PostCompanyRequest struct {
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Website     string `binding:"required"`
	GitHub      string `json:"github" binding:"required"`
	LinkedIn    string `json:"linkedin" binding:"required"`
}

var ErrorInvalidSlugLength = errors.New("invalid slug length")

func PostCompanyRequestToCompany(pcr PostCompanyRequest) (*models.Company, error) {
	slug := util.GetSlugFromName(pcr.Name)

	if slug == "" {
		return nil, ErrorInvalidSlugLength
	}

	return &models.Company{
		Name:        pcr.Name,
		Slug:        util.GetSlugFromName(pcr.Name),
		Description: pcr.Description,
		Location:    pcr.Location,
		Website:     pcr.Website,
		GitHub:      pcr.GitHub,
		LinkedIn:    pcr.LinkedIn,
	}, nil
}

func PostCompany(c *gin.Context) {
	// get and validate data from request
	var newCompany PostCompanyRequest
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !util.IsValidToken(c.Request.Header.Get("Authorization")) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	companyModel, err := PostCompanyRequestToCompany(newCompany)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if a record with the slug already exists, return duplicate error
	var found models.Company
	foundResult := database.GDB.Where("slug = ?", companyModel.Slug).First(&found)
	if foundResult.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "duplicate"})
		return
	}

	result := database.GDB.Create(&companyModel)

	if result.RowsAffected == 0 {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
