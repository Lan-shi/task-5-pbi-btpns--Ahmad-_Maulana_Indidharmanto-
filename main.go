package main

import (
	"fmt"
	"net/http"

	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/database"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/models"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/router"
)

func main() {
	db := database.InitDb()

	// Automatically create the tables based on models
	db.AutoMigrate(&models.User{}, &models.Photo{})

	r := router.SetupRouter()

	fmt.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
