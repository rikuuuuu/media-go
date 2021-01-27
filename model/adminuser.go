package model

import (
	// "time"
)

type AdminUser struct {
	ID		string		`db:"id" form:"id" json:"id"`
	Email	string		`db:"email" form:"email" json:"email"`
	Password	string	`db:"password" form:"password" json:"password"`
}

type CreateAdminUser struct {
	Id 			string `firestore:"id,omitempty"`
	Email      	string   `firestore:"email,omitempty"`
	Password    string   `firestore:"password,omitempty"`
	Name    	string   `firestore:"name,omitempty"`
}