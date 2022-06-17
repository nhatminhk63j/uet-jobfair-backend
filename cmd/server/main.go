package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"jobfair.uet.vnu.edu.vn/controllers"
	"jobfair.uet.vnu.edu.vn/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "welcome to uet job fair 2022!"})
	})

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.Default())

	models.ConnectDatabase()

	r.GET("/companies", controllers.FindCompanies)
	r.POST("/companies", controllers.CreateCompany)
	r.GET("/companies/:id", controllers.FindCompany)
	r.PATCH("/companies/:id", controllers.UpdateCompany)
	r.DELETE("/companies/:id", controllers.DeleteCompany)

	err := r.Run()
	if err != nil {
		log.Fatal("serve server error ....")
		return
	}
}
