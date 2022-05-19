package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermesoaresmarques/go-api-jwt-books/controllers"
	"github.com/guilhermesoaresmarques/go-api-jwt-books/services/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("users")
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		main.POST("login", controllers.Login)
		main.GET("login", controllers.ShowAllUsers)
		main.DELETE("login/:id", controllers.DeleteUser)
	}

	return router
}
