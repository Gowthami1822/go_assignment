package routes

import (
	"book-api/internal/book"
	"book-api/internal/jobs"
	"book-api/internal/user"
	"book-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", user.Login)
	r.POST("/jobs", jobs.SubmitJob) // public job submission

	bookRoutes := r.Group("/books")
	bookRoutes.Use(middleware.AuthMiddleware("admin", "user"))
	{
		bookRoutes.GET("", book.GetBooks)
		bookRoutes.GET(":id", book.GetBook)
		bookRoutes.POST("", book.AddBook)
		bookRoutes.PUT(":id", book.UpdateBook)
		bookRoutes.DELETE(":id", book.DeleteBook)
	}
}
