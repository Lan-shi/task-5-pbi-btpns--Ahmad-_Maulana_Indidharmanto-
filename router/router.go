package router

import (
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	return r
}
