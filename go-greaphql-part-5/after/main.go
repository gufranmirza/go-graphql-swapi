package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/after/schema"
)

func main() {

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/graphql", graphqlHandler)
	http.Handle("/", fs)

	fmt.Println("server is started at: http://localhost:8080/")
	fmt.Println("graphql api server is started at: http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		type GraphQLPostBody struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}

		var graphQLPostBody GraphQLPostBody
		err = json.Unmarshal(body, &graphQLPostBody)
		if err != nil {
			panic(err)
		}

		token := r.Header.Get("token")

		result := graphql.Do(graphql.Params{
			Schema:         schema.Schema,
			RequestString:  graphQLPostBody.Query,
			VariableValues: graphQLPostBody.Variables,
			OperationName:  graphQLPostBody.OperationName,
			Context:        context.WithValue(context.Background(), "token", token),
		})
		json.NewEncoder(w).Encode(result)

	default:
		fmt.Fprintf(w, "Sorry, only POST method are supported.")
	}
}
