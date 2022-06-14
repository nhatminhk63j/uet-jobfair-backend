package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"jobfair.uet.vnu.edu.vn/controllers"
	"jobfair.uet.vnu.edu.vn/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "welcome to uet job fair 2022!"})
	})

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
