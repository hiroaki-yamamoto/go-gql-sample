package models

import "encoding/base64"

import (
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/argon2"
)

// User model.
type User struct {
	gorm.Model
	Username  string  `json:"username" gorm:"vchar(256);unique_index"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

// Create stores newly-created user to db in the safe method.
// Note: This function doesn't check if the user already exists in the DB.
func (user *User) Create(db *gorm.DB) *User {
  user.Password = base64.StdEncoding.EncodeToString(argon2.IDKey(
    []byte(user.Password),
    []byte(user.Username),
    50, 128*1024, 9, 32,
  ))
  db.Create(user)
  return user
}
