package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/after/schema"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/after/subscriptions"
)

func main() {
	// GraphQL Schema
	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// graphql api  server
	http.Handle("/", h)

	// graphql subscriptions
	http.HandleFunc("/subscriptions", subscriptions.Handler)

	fmt.Println("server is started at: http://localhost:8080/")
	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	fmt.Println("subscriptions api server is started at: http://localhost:8080/subscriptions")
	http.ListenAndServe(":8080", nil)
}
