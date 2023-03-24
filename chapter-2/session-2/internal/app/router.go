package app

import (
	"chapter-2/session-2/internal/delivery"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", delivery.SaveBook)
	router.GET("/books", delivery.GetBooks)
	router.GET("/book/:id", delivery.GetBookByID)
	router.PUT("/book/:id", delivery.UpdateBook)
	router.DELETE("/book/:id", delivery.DeleteBook)

	return router
}
