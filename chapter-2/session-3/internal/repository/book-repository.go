package repository

import (
	"database/sql"
	"fmt"
	"session-3/challange-3/internal/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindById(id int64) (entity.Book, error)
	Save(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) error
}

func (r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	rows, err := r.db.Query("SELECT * FROM public.books ORDER BY id")
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book entity.Book

		err = rows.Scan(&book.ID, &book.NameBook, &book.Author, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *bookRepository) FindById(id int64) (entity.Book, error) {
	var book entity.Book

	err := r.db.QueryRow("SELECT * FROM public.books WHERE id = $1", id).Scan(&id, &book.NameBook, &book.Author, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Save(book entity.Book) (entity.Book, error) {
	_, err := r.db.Exec("INSERT INTO public.books (name_book, author, created_at, updated_at) VALUES ($1, $2, $3, $4)", book.NameBook, book.Author, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return book, err
	}

	row := r.db.QueryRow("SELECT id, name_book, author FROM public.books ORDER BY id DESC LIMIT 1")
	err = row.Scan(&book.ID, &book.NameBook, &book.Author)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Update(book entity.Book) (entity.Book, error) {
	res, err := r.db.Exec("UPDATE public.books SET name_book = $2, author = $3, updated_at = $4 WHERE id = $1", book.ID, book.NameBook, book.Author, book.UpdatedAt)
	if err != nil {
		return book, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return book, err
	}
	fmt.Println(count)

	return book, nil
}

func (r *bookRepository) Delete(book entity.Book) error {
	_, err := r.db.Exec("DELETE FROM public.books WHERE id = $1", book.ID)
	if err != nil {
		return err
	}

	return nil
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{
		db: db,
	}
}
