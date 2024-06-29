package gqlutil

import (
	"github.com/shurcooL/graphql"
	"net/http"
)

// authRoundTripper is a http.RoundTripper that adds an authorization header to each request.
type authRoundTripper struct {
	Transport http.RoundTripper
	Token     string
}

// RoundTrip adds the authorization header to the request before passing it to the underlying Transport.
func (a *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", a.Token)
	return a.Transport.RoundTrip(req)
}

// GetGraphqlClient creates a new GraphQL client targeting the specified GraphQL server URL with the provided auth token.
func GetGraphqlClient(url, authToken string) *graphql.Client {
	httpClient := &http.Client{
		Transport: &authRoundTripper{
			Transport: http.DefaultTransport,
			Token:     authToken,
		},
	}
	return graphql.NewClient(url, httpClient)
}
