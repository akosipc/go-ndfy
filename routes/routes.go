package routes

import (
	"ndfy/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "all good on our end"})
	})

	r.GET("/artists", controllers.FindArtists)
	r.POST("/artists", controllers.CreateArtist)
	r.GET("/artists/:permalink", controllers.FindArtist)
	r.PATCH("/artists/:permalink", controllers.UpdateArtist)
	r.DELETE("/artists/:permalink", controllers.DeleteArtist)

	return r
}
