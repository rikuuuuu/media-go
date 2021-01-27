package model

import (
	"time"
)

type Article struct {
	ID      int       `db:"id" form:"id" json:"id"`
	Title   string    `db:"title" form:"title" validate:"required,max=50" json:"title"`
	Body    string    `db:"body" form:"body" validate:"required" json:"body"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

type CreateArticle struct {
	Id       		string   `firestore:"id,omitempty"`
	UserId      	string   `firestore:"userId,omitempty"`
	Title    		string   `firestore:"title,omitempty"`
	Description    	string   `firestore:"description,omitempty"`
	CreatedAt 		string   `firestore:"createdAt,omitempty"`
	ThumbnailURL    string 	 `firestore:"thumbnailURL,omitempty"`
}