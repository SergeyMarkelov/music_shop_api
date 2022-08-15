package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/models"
)

func main() {

	basicRouter := gin.Default()

	basicRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Api is working!"})
	})

	//Router
	basicRouter.GET("/tracks", controllers.GetAllSongs)
	basicRouter.POST("/tracks", controllers.CreateSong)
	basicRouter.GET("/tracks/:id", controllers.GetSong)
	basicRouter.DELETE("/tracks/:id", controllers.DeleteSong)

	basicRouter.Run()
	// for connecting to our database
	models.Connect_to_DB()

}
