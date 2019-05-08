package prv

import (
	"context"
	"errors"

	gauthConf "github.com/hiroaki-yamamoto/gauth/config"
	gauthMid "github.com/hiroaki-yamamoto/gauth/middleware"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
)

// Resolver is a resolver to resolve private api.
type Resolver struct {
	Db      *prisma.Client
	TokConf *gauthConf.Config
}

// PrvM solves private-mutation.
func (r *Resolver) PrvM() PrvMResolver {
	return &prvMResolver{r}
}

// PrvQuery solves private-query.
func (r *Resolver) PrvQuery() PrvQueryResolver {
	return &prvQueryResolver{r}
}

// Subscription solves subscription-query.
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type prvMResolver struct{ *Resolver }

func (r *prvMResolver) Logout(ctx context.Context) (Status, error) {
	return Status{Ok: true}, nil
}

type prvQueryResolver struct{ *Resolver }

func (r *prvQueryResolver) Me(ctx context.Context) (prisma.User, error) {
	user := gauthMid.GetUser(ctx).(*prisma.User)
	// Want omitempty, but prisma.User is auto-generated model... :()
	user.Password = ""
	return *user, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) User(ctx context.Context, where *UserSubscriptionWhereInput) (<-chan *UserSubscriptionPayload, error) {
	return nil, errors.New("not implemented")
}
