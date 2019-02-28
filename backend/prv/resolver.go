package prv

import (
	"context"
)

import (
  "github.com/hiroaki-yamamoto/go-gql-sample/backend/auth"

)

type Resolver struct{}

func (r *Resolver) PrvQuery() PrvQueryResolver {
	return &prvQueryResolver{r}
}

type prvQueryResolver struct{ *Resolver }

func (r *prvQueryResolver) Me(ctx context.Context) (User, error) {
  user := auth.GetUser(ctx)
  return transferUser(&user), nil
}
