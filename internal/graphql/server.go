package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
)

func NewServer(resolver *Resolver) *handler.Server {
	srv := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{Resolvers: resolver},
		),
	)
	return srv
}

func PlaygroundHandler() http.Handler {
	return playground.Handler("GraphQL", "/query")
}
