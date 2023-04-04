package app

import (
	"database/sql"
	"session-3/challange-3/internal/delivery"
	"session-3/challange-3/internal/repository"
	"session-3/challange-3/internal/usecase"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	delivery := delivery.NewBookDelivery(bookUsecase)

	api := router.Group("/api")

	api.GET("/books", delivery.GetAllBooks)
	api.GET("/book/:id", delivery.GetBookByID)
	api.POST("/book", delivery.CreateBook)
	api.PUT("/book", delivery.UpdateBook)
	api.DELETE("/book/:id", delivery.DeleteBook)

	return router
}
