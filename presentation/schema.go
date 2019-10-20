package presentation

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}}

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
