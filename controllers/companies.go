package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"jobfair.uet.vnu.edu.vn/models"
)

type CompanyInfo struct {
	ID   int                `json:"id"`
	Name string             `json:"name"`
	Type models.CompanyType `json:"type"`
	Logo string             `json:"logo"`
}

// FindCompanies ...
// GET /companies
func FindCompanies(c *gin.Context) {
	var companies []models.Company
	models.DB.Find(&companies)

	res := make([]CompanyInfo, 0)
	for _, c := range companies {
		res = append(res, CompanyInfo{
			ID:   c.ID,
			Name: c.Name,
			Type: c.Type,
			Logo: c.Logo,
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

type CreateCompanyInput struct {
	Name        string             `json:"name" binding:"required"`
	Logo        string             `json:"logo" binding:"required"`
	Type        models.CompanyType `json:"companyType" binding:"required"`
	Description string             `json:"description"`
	PhoneNumber string             `json:"phoneNumber"`
	Email       string             `json:"email"`
	Address     string             `json:"address"`
	Website     string             `json:"website"`
	Facebook    string             `json:"facebook"`
	Linkedin    string             `json:"linkedin"`
}

// CreateCompany ...
// POST /companies
// Create a company
func CreateCompany(c *gin.Context) {
	// Validate input
	var input CreateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new company
	company := models.Company{
		Name:        input.Name,
		Logo:        input.Logo,
		Type:        input.Type,
		Description: input.Description,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Address:     input.Address,
		Website:     input.Website,
		Facebook:    input.Facebook,
		Linkedin:    input.Linkedin,
	}

	err := models.DB.Create(&company).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}

// FindCompany ...
// Get /companies/:id
// Find a company
func FindCompany(c *gin.Context) {
	var company models.Company
	if err := models.DB.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}

type UpdateCompanyInput struct {
	Name        string             `json:"name"`
	Logo        string             `json:"logo"`
	Type        models.CompanyType `json:"companyType"`
	Description string             `json:"description"`
	PhoneNumber string             `json:"phoneNumber"`
	Email       string             `json:"email"`
	Address     string             `json:"address"`
	Website     string             `json:"website"`
	Facebook    string             `json:"facebook"`
	Linkedin    string             `json:"linkedin"`
}

// UpdateCompany ...
// PATCH /companies/:id
// Update a company
func UpdateCompany(c *gin.Context) {
	// Get model if exist
	var company models.Company
	if err := models.DB.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Validate input
	var input UpdateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&company).Updates(input)
}

// DeleteCompany ...
// DELETE /companies/:id
func DeleteCompany(c *gin.Context) {
	// Get company if exist
	var company models.Company
	if err := models.DB.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := models.DB.Delete(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
