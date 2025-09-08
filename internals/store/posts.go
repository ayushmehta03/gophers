package store

import (
	"context"
	"database/sql"
)


type Post struct{
	ID int64         `json:"id"`
	Conetnt string  `json:"content"`
	Title string	`json:"title"`
	UserID string	`json:"user_id"`
	Tags []string	`json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type PostsStore struct{
	db *sql.DB

}

func (s *PostsStore) Create(ctx context.Context, post *Post) error{
	return nil
}