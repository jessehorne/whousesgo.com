package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/jessehorne/whousesgo.com/util"
	"net/http"
)

func GetCompanyList(c *gin.Context) {
	companies, err := util.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"companies": companies,
	})
}
