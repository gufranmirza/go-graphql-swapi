package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/after/resolvers"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-6/after/types"
)

var (
	Schema graphql.Schema
)

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
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
					return resolvers.GetHuman(id), nil
				},
			},
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createHuman": &graphql.Field{
				Type:        types.HumanType,
				Description: "Update existing human",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Description: "new name of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"homePlanet": &graphql.ArgumentConfig{
						Description: "new home planet of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"appearsIn": &graphql.ArgumentConfig{
						Description: "new episodes of the human",
						Type:        graphql.NewList(graphql.Int),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// marshall and cast the argument value
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)
					homePlanet, _ := params.Args["homePlanet"].(string)
					appearsIn, _ := params.Args["appearsIn"].([]interface{})

					// type assertion to convert []interface to []int
					appearsin := make([]int, len(appearsIn))
					for i := range appearsin {
						appearsin[i] = appearsIn[i].(int)
					}

					char := resolvers.CreatePerson(id, name, appearsin, homePlanet)
					HumanPublisher()

					return char, nil
				},
			},
		},
	})

	Subscription := graphql.NewObject(graphql.ObjectConfig{
		Name: "Subscription",
		Fields: graphql.Fields{
			"newHuman": &graphql.Field{
				Type: graphql.NewList(types.HumanType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolvers.HumanData, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        Query,
		Mutation:     Mutation,
		Subscription: Subscription,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
