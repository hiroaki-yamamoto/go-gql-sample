package models

import "encoding/base64"

import (
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/argon2"
)

// User Model
type User struct {
  gorm.Model
	Username  string  `gorm:"vchar(256);unique_index"`
	Email     string
	Password  string
	FirstName *string
	LastName  *string
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

// Authenticate finds user with specified username and password
func (user *User) Authenticate(
  db *gorm.DB, username string, password string,
) *User {
  pwHash := base64.StdEncoding.EncodeToString(argon2.IDKey(
    []byte(password),
    []byte(username),
    50, 128*1024, 9, 32,
  ))
  db.Where("username = ? AND password = ?", username, pwHash).Take(user)
  return user
}
