package pub

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) PubM() PubMResolver {
	return &pubMResolver{r}
}
func (r *Resolver) PubQ() PubQResolver {
	return &pubQResolver{r}
}

type pubMResolver struct{ *Resolver }

func (r *pubMResolver) Login(ctx context.Context, username string, password string) (User, error) {
	panic("not implemented")
}
func (r *pubMResolver) Signup(ctx context.Context, username string, password string, email string, firstName *string, lastName *string) (*Error, error) {
	panic("not implemented")
}

type pubQResolver struct{ *Resolver }

func (r *pubQResolver) Country(ctx context.Context) ([]*string, error) {
	panic("not implemented")
}
