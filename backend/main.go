package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/gbrlsnchs/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	gauthConf "github.com/hiroaki-yamamoto/gauth/config"
	guathMid "github.com/hiroaki-yamamoto/gauth/middleware"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prisma"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/prv"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/pub"
)

const defaultPort = "8080"

func findUser(fcon interface{}, username string) (interface{}, error) {
	con := fcon.(*prisma.Client)
	return con.User(prisma.UserWhereUniqueInput{ID: &username}).Exec(
		context.TODO(),
	)
}

func main() {
	con := prisma.New(nil)
	config, err := gauthConf.New(
		jwt.NewHS256("test"),
		"Test Audience", "Test Issuer",
		"Test Subject", 3600*time.Minute,
	)
	if err != nil {
		log.Fatal(err)
	}
	headerMiddleware := guathMid.HeaderMiddleware(
		"Auth", con, findUser, config,
	)
	headerRequired := guathMid.HeaderLoginRequired(
		"Auth", con, findUser, config,
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
			pub.Config{Resolvers: &pub.Resolver{Db: con, TokConf: config}}),
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
