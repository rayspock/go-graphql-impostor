//go:generate go run github.com/99designs/gqlgen -config ../api/gqlgen.yml

package gql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	graphqlEndpoint string
}

func NewResolver(graphqlEndpoint string) *Resolver {
	return &Resolver{
		graphqlEndpoint: graphqlEndpoint,
	}
}
