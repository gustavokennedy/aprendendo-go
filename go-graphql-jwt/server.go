package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/graphql-go/http"
	"gitlab.com/pragmaticreviews/graphql-go/middleware"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	server := gin.Default()

	server.Use()

	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", middleware.AuthorizeJWT(), http.GraphQLHandler())

	server.Run(defaultPort)
}
