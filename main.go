package main

import (
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/database"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/models"
)

func main() {
	db := database.InitDb()

	// Automatically create the tables based on models
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
