package schema

import (
	"errors"
	"strconv"

	"github.com/graphql-go/graphql/gqlerrors"

	"github.com/graphql-go/graphql"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/after/auth"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/after/resolvers"
	"github.com/gufranmirza/go-graphql-swapi/go-greaphql-part-5/after/types"
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
					// validate token
					isValid, err := auth.ValidateToken(p.Context.Value("token").(string))
					if err != nil {
						return nil, err
					}
					if !isValid {
						return nil, gqlerrors.FormatError(errors.New("Invalid token"))
					}

					// get person id
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

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createToken": &graphql.Field{
				Type:        graphql.String,
				Description: "creates a new  JWT token ",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Description: "username",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Description: "password",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// marshall and cast the argument value
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)

					token, err := auth.CreateToken(username, password)
					if err != nil {
						return nil, err
					}

					return token, nil

				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		panic(err)
	}
	Schema = schema
}
