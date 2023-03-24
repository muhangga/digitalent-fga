package delivery

import (
	"chapter-2/session-2/internal/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

var bookData = []entity.Book{
	{
		ID:     "1",
		Title:  "The Alchemist",
		Author: "Paulo Coelho",
		Desc:   "The Alchemist is a novel by Brazilian author Paulo Coelho",
	},
	{
		ID:     "2",
		Title:  "The Kite Runner",
		Author: "Khaled Hosseini",
		Desc:   "Lorem ipsum sit amet",
	},
}

func SaveBook(c *gin.Context) {
	var book entity.Book

	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
			"code":    405,
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book.ID = fmt.Sprintf("%d", len(bookData)+1)
	c.JSON(200, gin.H{
		"book": book,
	})
}

func GetBooks(c *gin.Context) {

	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
			"code":    405,
		})
	}

	if len(bookData) > 0 {
		c.JSON(200, gin.H{
			"books": bookData,
		})
		return
	}

	c.JSON(404, gin.H{
		"message": "No books found",
	})
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
			"code":    405,
		})
		return
	}

	for _, book := range bookData {
		if book.ID == id {
			c.JSON(200, gin.H{
				"book": book,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Book not found",
	})
}

func UpdateBook(c *gin.Context) {
	var book entity.Book

	if c.Request.Method != "PUT" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
			"code":    405,
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Book updated successfully!"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Method != "DELETE" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
			"code":    405,
		})
		return
	}

	for i, book := range bookData {
		if book.ID == id {
			bookData = append(bookData[:i], bookData[i+1:]...)
			c.JSON(200, gin.H{
				"message": "Book deleted successfully!",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "Book not found",
	})
}
