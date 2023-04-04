package delivery

import (
	"net/http"
	"session-3/challange-3/internal/entity"
	"session-3/challange-3/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookDelivery struct {
	bookUsecase usecase.BookUsecase
}

func NewBookDelivery(bookUsecase usecase.BookUsecase) *bookDelivery {
	return &bookDelivery{bookUsecase}
}

func (d *bookDelivery) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := d.bookUsecase.GetBookByID(idConv)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (d *bookDelivery) GetAllBooks(c *gin.Context) {
	books, err := d.bookUsecase.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (d *bookDelivery) CreateBook(c *gin.Context) {
	var book entity.BookRequest

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := d.bookUsecase.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (d *bookDelivery) UpdateBook(c *gin.Context) {
	var book entity.BookRequest

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := d.bookUsecase.UpdateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (d *bookDelivery) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.bookUsecase.DeleteBook(idConv)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book successfully deleted"})
}
