package usecase

import (
	_ "database/sql"
	"fmt"
	"session-3/challange-3/internal/entity"
	"session-3/challange-3/internal/repository"
	"time"
)

type BookUsecase interface {
	GetBookByID(id int) (entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	CreateBook(req entity.BookRequest) (entity.Book, error)
	UpdateBook(req entity.BookRequest) (entity.Book, error)
	DeleteBook(id int) error
}

func (u *bookUsecase) GetBookByID(id int) (entity.Book, error) {

	book, err := u.repository.FindById(int64(id))
	if err != nil {
		return book, err
	}

	return book, nil
}

func (u *bookUsecase) GetAllBooks() ([]entity.Book, error) {

	books, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) CreateBook(req entity.BookRequest) (entity.Book, error) {

	var book entity.Book
	book.NameBook = req.NameBook
	book.Author = req.Author
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	saveBook, err := u.repository.Save(book)
	if err != nil {
		return book, err
	}

	return saveBook, nil
}

func (u *bookUsecase) UpdateBook(req entity.BookRequest) (entity.Book, error) {
	books, err := u.repository.FindById(req.ID)
	if err != nil {
		return books, fmt.Errorf("book not found")
	}

	var book entity.Book
	book.ID = req.ID
	book.NameBook = req.NameBook
	book.Author = req.Author
	book.UpdatedAt = time.Now()

	book, err = u.repository.Update(book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (u *bookUsecase) DeleteBook(id int) error {

	book, err := u.repository.FindById(int64(id))
	if err != nil {
		return fmt.Errorf("Book not found")
	}

	err = u.repository.Delete(book)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(repository repository.BookRepository) *bookUsecase {
	return &bookUsecase{
		repository: repository,
	}
}
