package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/service"
)

// Save Book
// @Summary Save Book
// @Description Save Book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body entity.BookDTO true "Book"
// @Success 201 {object} entity.BookDTO
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /api/book [post]
func SaveBook(c *fiber.Ctx) error {
	var bookDTO entity.BookDTO
	if err := c.BodyParser(&bookDTO); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"errors":  err,
		})
	}

	book, err := service.InsertBook(bookDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"errors":  err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Succussfully Created",
		"status":  "Success",
		"data":    book,
	})
}

// Get All Book
// @Summary Get All Book
// @Description Get All Book
// @Tags Book
// @Accept json
// @Produce json
// @Success 200 {object} entity.Book
// @Failure 500 {object} entity.Error
// @Router /api/books [get]
func GetAllBook(c *fiber.Ctx) error {
	books, err := service.GetAllBook()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"errors":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Get All Books",
		"status":  "Success",
		"data":    books,
	})
}

// Get Book By ID
// @Summary Get Book By ID
// @Description Get Book By ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} entity.Book
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /api/book/{id} [get]
func GetBookByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	// var book entity.Book

	book, err := service.GetBookByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Get Book By ID",
		"status":  "Success",
		"data":    book,
	})

}

// Update Book
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body entity.BookDTO true "Book"
// @Success 200 {object} entity.Book
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /api/book/{id} [put]
func UpdateBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var bookDTO entity.BookDTO
	if err := c.BodyParser(&bookDTO); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"errors":  err,
		})
	}

	book, err := service.UpdateBook(id, bookDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"errors":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Update Book",
		"status":  "success",
		"data":    book,
	})
}

// Delete Book
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} entity.Book
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /api/book/{id} [delete]
func DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	_, err := service.DeleteBook(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Delete Book",
	})
}
