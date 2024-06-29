package main

import (
	"github.com/gorilla/mux"
	"github.com/rayspock/go-graphql-impostor/gql"
	"github.com/rayspock/go-graphql-impostor/pkg/gqlutil"
	"net/http"

	"github.com/rs/cors"
	"log"
	"os"

	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/viper"
)

const DefaultPort = "9002"

const GraphQLEndpoint = "https://example.com/v1/graphql/"

type config struct {
	stubFields []string
}

func readConfig() config {
	return config{
		stubFields: viper.GetStringSlice("stub_fields"),
	}
}

func init() {
	viper.SetConfigFile("env.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read config: %s\n", err)
	}
}

func main() {
	// load configuration
	cfg := readConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	resolver := gql.NewResolver(GraphQLEndpoint)

	srv := gqlHandler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: resolver}))
	srv.AroundOperations(gqlutil.GetStubMiddleware(cfg.stubFields, GraphQLEndpoint))

	router := mux.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql/"))
	router.Handle("/graphql/", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
