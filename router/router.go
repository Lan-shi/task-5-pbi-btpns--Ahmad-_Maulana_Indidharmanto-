package router

import (
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/controllers"
	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.PUT("/users/update", controllers.UpdateUser)
	r.DELETE("/users/delete", controllers.DeleteUser)

	r.POST("/photos/upload", controllers.UploadPhoto)
	r.GET("/photos/:id", controllers.GetPhoto)
	r.PUT("/photos/update", controllers.UpdatePhoto)
	r.DELETE("/photos/:id", controllers.DeletePhoto)

	authenticated := r.Group("/")
	authenticated.Use(middlewares.AuthMiddleware())
	{
		authenticated.PUT("/users/update", controllers.UpdateUser)
		authenticated.DELETE("/users/delete", controllers.DeleteUser)
		authenticated.POST("/photos/upload", controllers.UploadPhoto)
		authenticated.GET("/photos/:id", controllers.GetPhoto)
		authenticated.PUT("/photos/update", controllers.UpdatePhoto)
		authenticated.DELETE("/photos/:id", controllers.DeletePhoto)
	}

	return r
}
