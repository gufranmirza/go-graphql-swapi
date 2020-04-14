package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/before/schema"
)

func main() {
	// GraphQL Schema

	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// graphql api  server
	http.Handle("/graphql", h)

	fmt.Println("server is started at: http://localhost:8080/")
	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
