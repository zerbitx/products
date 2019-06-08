package main

import (
	"github.com/zerbitx/federation-demo/products/service"
	"log"
	"net/http"
	"os"
	
	"github.com/99designs/gqlgen/handler"
	"github.com/zerbitx/federation-demo/products"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(products.NewExecutableSchema(products.Config{Resolvers: products.New(service.New())})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
