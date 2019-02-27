package pub

import (
	"context"
  "net/http"

	"github.com/hiroaki-yamamoto/go-gql-sample/backend/models"
  "github.com/hiroaki-yamamoto/go-gql-sample/backend/auth"
  "github.com/hiroaki-yamamoto/go-gql-sample/backend/middleware"
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
	user := &models.User{}
  resp, ok := ctx.Value(middleware.ResponseKey).(*http.ResponseWriter)
  if !ok {
    panic(
      "The response couldn't be loaded. " +
      "Check whether ContextReqRespMiddleware is loaded",
    )
  }
  user.Authenticate(r.Db, username, password)
  auth.Login(resp, user)
  return transferUser(user), nil
}
func (r *pubMResolver) Signup(ctx context.Context, username string, password string, email string, firstName *string, lastName *string) (UserAndError, error) {
  user := &models.User{
    Username: username,
    Password: password,
    Email: email,
    FirstName: firstName,
    LastName: lastName,
  }
  user.Create(r.Db)

  // Do you have any idea to assign fields from models.User in smart way??
  // If you have, send me PR.
  return transferUser(user), nil
}

type pubQResolver struct{ *Resolver }

func (r *pubQResolver) Country(ctx context.Context) ([]*string, error) {
	panic("not implemented")
}
