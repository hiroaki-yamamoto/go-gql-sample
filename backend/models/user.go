package models

import "github.com/jinzhu/gorm"

// User model.
type User struct {
	gorm.Model
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
