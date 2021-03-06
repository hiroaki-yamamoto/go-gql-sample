package pub

import (
	"context"
	"errors"

	gauthConf "github.com/hiroaki-yamamoto/gauth/config"
	gauth "github.com/hiroaki-yamamoto/gauth/core"
	gauthMid "github.com/hiroaki-yamamoto/gauth/middleware"

	"github.com/hiroaki-yamamoto/go-gql-sample/backend/hash"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
)

// Resolver is a struct for pub-api reolver
type Resolver struct {
	Db      *prisma.Client
	TokConf *gauthConf.Config
}

// PubM indicates a public mutation
func (r *Resolver) PubM() PubMResolver {
	return &pubMResolver{r}
}

// PubQ indicates a public query
func (r *Resolver) PubQ() PubQResolver {
	return &pubQResolver{r}
}

// Subscription indicates a public subscription (what??)
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type pubMResolver struct{ *Resolver }

func (r *pubMResolver) Login(
	ctx context.Context,
	username string,
	password string,
) (SessionAndStatus, error) {
	if gauthMid.GetUser(ctx) != nil {
		return nil, errors.New("You are already logged in")
	}
	pwHash := hash.Argon2(password, username)
	users, err := r.Db.Users(&prisma.UsersParams{
		Where: &prisma.UserWhereInput{
			Username: &username,
			Password: &pwHash,
		},
	}).Exec(context.TODO())
	if err != nil {
		return nil, err
	}
	if len(users) < 1 {
		return nil, errors.New("Username and / or password are mismatch")
	}
	// In the most case, this code is enought to tokenize:
	tok, err := gauth.ComposeID(users[0].ID, r.TokConf)
	// However, if you want to use in advanced mode,
	// ComposeToken is also useful.
	// now := time.Now().UTC()
	// tok, err := gauth.ComposeToken(&jwt.JWT{
	// 	Issuer:         r.TokConf.Issuer,
	// 	Subject:        r.TokConf.Subject,
	// 	Audience:       r.TokConf.Audience,
	// 	ExpirationTime: now.Add(2 * time.Hour).Unix(),
	// 	IssuedAt:       now.Unix(),
	// 	ID:             users[0].Username,
	// }, r.TokConf.Signer)
	if err != nil {
		return nil, err
	}
	return Session{string(tok)}, nil
}

func (r *pubMResolver) Signup(
	ctx context.Context,
	username string,
	password string,
) (UserAndStatus, error) {
	return r.Db.CreateUser(prisma.UserCreateInput{
		Username: username,
		Password: hash.Argon2(password, username),
	}).Exec(context.TODO())
}

type pubQResolver struct{ *Resolver }

func (r *pubQResolver) Country(ctx context.Context) ([]*string, error) {
	return nil, errors.New("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) User(ctx context.Context, where *UserSubscriptionWhereInput) (<-chan *UserSubscriptionPayload, error) {
	return nil, errors.New("not implemented")
}
