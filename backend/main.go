package main

import (
	"log"
  "net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

  "github.com/hiroaki-yamamoto/go-gql-sample/backend/pub"
  "github.com/hiroaki-yamamoto/go-gql-sample/backend/auth"
)

const defaultPort = "8080"

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic(err)
  }
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  router := chi.NewRouter()
  router.Use(middleware.RequestID)
  router.Use(middleware.RealIP)
  router.Use(middleware.Logger)
  router.Use(middleware.Recoverer)
  router.Use(auth.AuthenticationMiddleware(db))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(pub.NewExecutableSchema(
		pub.Config{Resolvers: &pub.Resolver{ Db: db }}),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
  if err != nil {
		panic(err)
	}
}
