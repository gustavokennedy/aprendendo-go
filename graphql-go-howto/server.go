package main

import (
	"github.com/glyphack/go-graphql-hackernews/internal/auth"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	hackernews "github.com/glyphack/go-graphql-hackernews"
	"github.com/glyphack/go-graphql-hackernews/internal/pkg/db/mysql"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
