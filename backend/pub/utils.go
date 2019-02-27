package pub

// helper utils for public API

import "github.com/hiroaki-yamamoto/go-gql-sample/backend/models"

func transferUser(from *models.User) User {
  return User{
    Username: from.Username,
    Password: from.Password,
    Email: from.Email,
    FirstName: from.FirstName,
    LastName: from.LastName,
  }
}
