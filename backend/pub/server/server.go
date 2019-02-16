package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/hiroaki-yamamoto/go-gql-sample/backend/pub"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const defaultPort = "8080"

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(pub.NewExecutableSchema(
		pub.Config{Resolvers: &pub.Resolver{
			db: &db,
		}}),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
