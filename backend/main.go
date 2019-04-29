package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gbrlsnchs/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	gauth "github.com/hiroaki-yamamoto/gauth/core"
	guathMid "github.com/hiroaki-yamamoto/gauth/middleware"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prv"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/pub"
)

const defaultPort = "8080"

func findUser(fcon interface{}, username string) (interface{}, error) {
	con := fcon.(*prisma.Client)
	return con.User(prisma.UserWhereUniqueInput{Username: &username}).Exec(
		context.TODO(),
	)
}

func main() {
	con := prisma.New(nil)
	config := gauth.Config{
		Signer:   jwt.NewHS256("test"),
		Audience: "Test Audience",
		Issuer:   "Test Issuer",
		Subject:  "Test Subject",
	}
	headerMiddleware := guathMid.HeaderMiddleware(
		"Auth", con, findUser, &config,
	)
	headerRequired := guathMid.HeaderLoginRequired(
		"Auth", con, findUser, &config,
	)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Handle("/", handler.Playground("GraphQL playground", "/pub"))
	router.Handle("/pub", headerMiddleware(
		handler.GraphQL(pub.NewExecutableSchema(
			pub.Config{Resolvers: &pub.Resolver{Db: con, TokConf: &config}}),
		),
	))
	router.Handle("/prv", headerRequired(
		handler.GraphQL(prv.NewExecutableSchema(
			prv.Config{Resolvers: &prv.Resolver{Db: con}}),
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
