package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/models"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

// get all songs
func GetAllSongs(context *gin.Context) {
	var song []models.Song
	models.DB.Find(&song)

	context.JSON(http.StatusOK, gin.H{"songs": song})
}

// Get track by id
func GetSong(context *gin.Context) {
	// need to check do we have that song..
	var song models.Song
	err := models.DB.Where("id = ?", context.Param("id")).First(&song).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "we haven't this song :("})
		return
	}

	context.JSON(http.StatusOK, gin.H{"songs": song})
}

// It's for track creating
func CreateSong(context *gin.Context) {
	var input CreateTrackInput
	//cheking our input
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song := models.Song{Artist: input.Artist, Title: input.Title}
	models.DB.Create(&song)

	context.JSON(http.StatusOK, gin.H{"tracks": song})
}

// DELETE
func DeleteSong(context *gin.Context) {
	// need to check data before deleting
	var track models.Song
	err := models.DB.Where("id = ?", context.Param("id")).First(&track).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "We haven't this song"})
		return
	}

	models.DB.Delete(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": true})
}
