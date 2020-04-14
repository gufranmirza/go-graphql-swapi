package schema

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql/gqlerrors"

	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/before/resolvers"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/before/types"
)

var (
	Schema graphql.Schema
)

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", gqlerrors.FormatError(fmt.Errorf("error in query"))
				},
			},
			"human": &graphql.Field{
				Type: types.HumanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}

					char, err := resolvers.GetHuman(id)
					if err != nil {
						return nil, err
					}

					return char, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
