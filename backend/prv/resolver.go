package prv

import (
	"context"

	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
)

type Resolver struct {
	Db *prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) PrvQuery() PrvQueryResolver {
	return &prvQueryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, data UserCreateInput) (prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, data UserUpdateInput, where UserWhereUniqueInput) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateManyUsers(ctx context.Context, data UserUpdateManyMutationInput, where *UserWhereInput) (BatchPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertUser(ctx context.Context, where UserWhereUniqueInput, create UserCreateInput, update UserUpdateInput) (prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, where UserWhereUniqueInput) (*prisma.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteManyUsers(ctx context.Context, where *UserWhereInput) (BatchPayload, error) {
	panic("not implemented")
}

type prvQueryResolver struct{ *Resolver }

func (r *prvQueryResolver) Me(ctx context.Context) (prisma.User, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) User(ctx context.Context, where *UserSubscriptionWhereInput) (<-chan *UserSubscriptionPayload, error) {
	panic("not implemented")
}
