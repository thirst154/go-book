package models

import "time"

type Post struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Text          string    `json:"text"`
	Title         string    `json:"title"`
	DatePublished time.Time `json:"date_published"`
	AuthorID      uint      `json:"author_id"`
}
