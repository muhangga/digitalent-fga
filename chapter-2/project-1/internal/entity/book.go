package entity

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primary_key"`
	NameBook  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create, default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update, default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type BookDTO struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}

type Error struct {
	Message string `json:"message"`
}
