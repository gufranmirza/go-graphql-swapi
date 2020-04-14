package types

import "github.com/graphql-go/graphql"

var episodeEnum = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Episode",
	Description: "One of the films in the Star Wars Trilogy",
	Values: graphql.EnumValueConfigMap{
		"NEWHOPE": &graphql.EnumValueConfig{
			Value:       4,
			Description: "Released in 1977.",
		},
		"EMPIRE": &graphql.EnumValueConfig{
			Value:       5,
			Description: "Released in 1980.",
		},
		"JEDI": &graphql.EnumValueConfig{
			Value:       6,
			Description: "Released in 1983.",
		},
	},
})

var HumanType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Human",
	Description: "A humanoid creature in the Star Wars universe.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The id of the human.",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "The name of the human.",
		},
		"appearsIn": &graphql.Field{
			Type:        graphql.NewList(episodeEnum),
			Description: "Which movies they appear in.",
		},
		"homePlanet": &graphql.Field{
			Type:        graphql.String,
			Description: "The home planet of the human, or null if unknown.",
		},
	},
})
