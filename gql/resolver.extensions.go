package gql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

const (
	AuthorizationHeader = "Authorization"
)

func (r *Resolver) GetAuthTokenFromContext(ctx context.Context) string {
	oc := graphql.GetOperationContext(ctx)
	var authToken string
	if len(oc.Headers[AuthorizationHeader]) > 0 {
		authToken = oc.Headers[AuthorizationHeader][0]
	}
	return authToken
}
