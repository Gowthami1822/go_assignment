package routes

import (
	"book-api/controllers"
	"book-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	bookRoutes := r.Group("/books")
	bookRoutes.Use(middleware.AuthMiddleware())
	{
		bookRoutes.GET("", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("", controllers.AddBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
}
