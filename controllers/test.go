package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type createTestInput struct {
}

type updateTestInput struct {
}


// FindTest is get all tests
func FindTest(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	var tests []models.Test
	db.Find(&tests)

	c.JSON(http.StatusOK, tests)
}

// CreateTest is create a test
func CreateTest(c *gin.Context) {

	var input createTestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	test := models.Test{}

	c.JSON(http.StatusOK, test)
}

// FindTest is find a test
func FindTest(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var test models.Test
	if err := db.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, test)
}

// UpdateTest is update a test
func UpdateTest(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var test models.Test
	if err := db.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input updateTestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&test).Updates(input)

	c.JSON(http.StatusOK, test)
}

// DeleteTest is delete a test
func DeleteTest(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var test models.Test
	if err := db.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&test)

	c.JSON(http.StatusOK, test)
}