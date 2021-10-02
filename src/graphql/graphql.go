package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/rlawnsxo131/madre-server/src/components/user"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": user.UserQueryType,
		},
	},
)

func NewHandler() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	}), nil
}
