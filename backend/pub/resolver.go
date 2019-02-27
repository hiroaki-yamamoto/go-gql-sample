package pub

import (
	"context"

	"github.com/hiroaki-yamamoto/go-gql-sample/backend/models"
	"github.com/jinzhu/gorm"
)

type Resolver struct {
	Db *gorm.DB
}

func (r *Resolver) PubM() PubMResolver {
	return &pubMResolver{r}
}
func (r *Resolver) PubQ() PubQResolver {
	return &pubQResolver{r}
}

type pubMResolver struct{ *Resolver }

func (r *pubMResolver) Login(ctx context.Context, username string, password string) (UserAndError, error) {
	panic("not implemented")
}
func (r *pubMResolver) Signup(ctx context.Context, username string, password string, email string, firstName *string, lastName *string) (UserAndError, error) {
  user := models.User{
    Username: username,
    Password: password,
    Email: email,
    FirstName: firstName,
    LastName: lastName,
  }
  (&user).Create(r.Db)

  // Do you have any idea to assign fields from models.User in smart way??
  // If you have, send me PR.
  return User{
    Username: user.Username,
    Password: user.Password,
    Email: user.Email,
    FirstName: user.FirstName,
    LastName: user.LastName,
  }, nil
}

type pubQResolver struct{ *Resolver }

func (r *pubQResolver) Country(ctx context.Context) ([]*string, error) {
	panic("not implemented")
}
