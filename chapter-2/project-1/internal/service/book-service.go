package service

import (
	"github.com/muhangga/internal/config"
	"github.com/muhangga/internal/entity"
)

func InsertBook(bookDTO entity.BookDTO) (entity.Book, error) {

	db := config.DBConn()
	defer config.CloseConn(db)

	var book entity.Book

	book.NameBook = bookDTO.NameBook
	book.Author = bookDTO.Author

	db.Begin()
	save := db.Save(&book).Error
	if save != nil {
		db.Rollback()
		return book, save
	}
	db.Commit()

	return book, nil
}

func GetAllBook() ([]entity.Book, error) {

	db := config.DBConn()
	defer config.CloseConn(db)

	var books []entity.Book

	findBooks := db.Table("books").Find(&books)
	if findBooks.Error != nil {
		return books, findBooks.Error
	}

	return books, nil
}

func GetBookByID(id int) (entity.Book, error) {

	db := config.DBConn()
	defer config.CloseConn(db)

	var book entity.Book

	findBook := db.Table("books").Where("id = ?", id).First(&book)
	if findBook.Error != nil {
		return book, findBook.Error
	}

	return book, nil
}

func UpdateBook(id int, bookDTO entity.BookDTO) (entity.Book, error) {

	db := config.DBConn()
	defer config.CloseConn(db)

	var book entity.Book

	findBook := db.Table("books").Where("id = ?", id).First(&book)
	if findBook.Error != nil {
		return book, findBook.Error
	}

	book.NameBook = bookDTO.NameBook
	book.Author = bookDTO.Author

	db.Begin()
	err := db.Save(&book).Error
	if err != nil {
		db.Rollback()
		return book, err
	}
	db.Commit()

	return book, nil
}

func DeleteBook(id int) (entity.Book, error) {

	db := config.DBConn()
	defer config.CloseConn(db)

	var book entity.Book

	findBook := db.Table("books").Where("id = ?", id).First(&book)
	if findBook.Error != nil {
		return book, findBook.Error
	}

	db.Begin()
	delete := db.Delete(&book).Error
	if delete != nil {
		db.Rollback()
		return book, delete
	}
	db.Commit()

	return book, nil
}
