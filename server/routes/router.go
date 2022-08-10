package routes

import (
	"github.com/andrade-felipe-dev/first-api-in-go/controllers"
	"github.com/andrade-felipe-dev/first-api-in-go/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("user")
		{
			users.GET("/:id", controllers.CreateUser)
			users.GET("/", controllers.ShowAllUsers)
			users.POST("/", controllers.CreateUser)
			users.PUT("/", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}

		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowAllBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		main.POST("login", controllers.Login)
	}

	return router
}
