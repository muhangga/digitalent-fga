package entity

import "time"

type Book struct {
	ID        int64     `json:"id"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type BookRequest struct {
	ID       int64  `json:"id"`
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}
