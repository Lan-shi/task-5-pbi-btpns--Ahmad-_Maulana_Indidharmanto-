package controllers

import (
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/database"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/models"
	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	db := database.InitDb()

	file, _ := c.FormFile("file")
	path := "photos/" + file.Filename
	c.SaveUploadedFile(file, path)

	photo := models.Photo{Path: path}
	db.Create(&photo)

	c.JSON(200, gin.H{"message": "Photo uploaded successfully"})
}

func GetPhoto(c *gin.Context) {
	db := database.InitDb()

	var photo models.Photo
	db.First(&photo, c.Param("id"))

	c.JSON(200, photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.InitDb()

	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&photo)

	c.JSON(200, gin.H{"message": "Photo updated successfully"})
}

func DeletePhoto(c *gin.Context) {
	db := database.InitDb()

	var photo models.Photo
	db.First(&photo, c.Param("id"))
	db.Delete(&photo)

	c.JSON(200, gin.H{"message": "Photo deleted successfully"})
}
