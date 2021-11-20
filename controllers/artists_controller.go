package controllers

import (
	"ndfy/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateArtistInput struct {
	Name      string     `form:"name" json:"name" binding:"required"`
	Permalink string     `form:"permalink" json:"permalink" binding:"required"`
	Bio       string     `form:"bio" json:"bio"`
	FormedAt  *time.Time `form:"formed_at" time_format:"2006-01-02" json:"formed_at"`
}

type UpdateArtistInput struct {
	Name      string `form:"name" json:"name"`
	Permalink string `form:"permalink" json:"permalink"`
	Bio       string `form:"bio" json:"bio"`
}

func FindArtists(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var artists []models.Artist

	db.Find(&artists)

	c.JSON(http.StatusOK, gin.H{"data": artists})
}

func CreateArtist(c *gin.Context) {
	var input CreateArtistInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	artist := models.Artist{
		Name:      input.Name,
		Permalink: input.Permalink,
		Bio:       input.Bio,
		FormedAt:  input.FormedAt,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&artist)

	c.JSON(http.StatusOK, gin.H{"data": artist})
}

func FindArtist(c *gin.Context) {
	var artist models.Artist
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("permalink = ?", c.Param("permalink")).First(&artist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": artist})
}

func UpdateArtist(c *gin.Context) {
	var artist models.Artist

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("permalink = ?", c.Param("permalink")).First(&artist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not Found"})
		return
	}

	var input UpdateArtistInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	db.Model(&artist).Updates(models.Artist{
		Name:      input.Name,
		Permalink: input.Permalink,
		Bio:       input.Bio,
	})

	c.JSON(http.StatusOK, gin.H{"data": artist})
}

func DeleteArtist(c *gin.Context) {
	var artist models.Artist
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("permalink = ?", c.Param("permalink")).First(&artist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not Found"})
		return
	}

	db.Delete(&artist)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
