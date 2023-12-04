package controllers

import (
	"time"

	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/database"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	db := database.InitDb()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.Create(&user)

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	db := database.InitDb()

	var user models.User
	var incomingUser models.User
	if err := c.ShouldBindJSON(&incomingUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", incomingUser.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Email or password is incorrect"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(incomingUser.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Email or password is incorrect"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("your_secret_key"))

	c.JSON(200, gin.H{"token": tokenString})
}

func UpdateUser(c *gin.Context) {
	db := database.InitDb()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&user)

	c.JSON(200, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	db := database.InitDb()

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&user)

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
