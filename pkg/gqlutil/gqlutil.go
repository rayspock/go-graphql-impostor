package gqlutil

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rayspock/go-graphql-impostor/pkg/jsonutil"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetStubMiddleware(stubFields []string, fallbackUrl string) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		// Check if the root field in current operation is a stub field
		// if so, return the stub response directly
		// Caveat: the operation is indiscrete in this stage, so we assume that
		// all root fields in the operation are stub fields if any of them are
		rootFields := getRootFieldString(oc)
		fieldStr := strings.Join(rootFields, ",")
		for _, rootField := range rootFields {
			for _, stubField := range stubFields {
				if rootField == stubField {
					log.Printf("execute %s [%s] - Stubs", oc.OperationName, fieldStr)
					return next(ctx)
				}
			}
		}
		log.Printf("execute %s [%s]", oc.OperationName, fieldStr)

		resp, err := forwardGraphqlRequest(oc, fallbackUrl)
		if err != nil {
			log.Printf("Failed to forward request: %s\n", err)
			return next(ctx)
		}

		return func(ctx context.Context) *graphql.Response {
			graphqlResp := graphql.Response{}
			err := json.Unmarshal(resp, &graphqlResp)
			if err != nil {
				log.Printf("Failed to unmarshal response: %s\n", err)
				return nil
			}
			return &graphqlResp
		}
	}
}

func getRootFieldString(reqCtx *graphql.OperationContext) []string {
	cfs := graphql.CollectFields(reqCtx, reqCtx.Operation.SelectionSet, nil)
	var rootFields []string
	for _, cf := range cfs {
		rootFields = append(rootFields, cf.Name)
	}
	return rootFields
}

func forwardGraphqlRequest(oc *graphql.OperationContext, url string) ([]byte, error) {
	payload := struct {
		Query     string               `json:"query"`
		Variables jsonutil.VariableMap `json:"variables"`
	}{
		Query:     oc.RawQuery,
		Variables: oc.Variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %s\n", err)
		return nil, err
	}

	resp, err := sendHttpRequest(oc.Headers, body, url)
	if err != nil {
		log.Printf("Failed to send request: %s\n", err)
		return nil, err
	}
	return resp, nil
}

func sendHttpRequest(headers http.Header, body []byte, url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	// fix failed to unmarshal response: invalid character '\x1f' looking for beginning of value
	// not support compressed responses from the server
	delete(req.Header, "Accept-Encoding")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
