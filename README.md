# GraphQL Impostor

## Description

This tool is used to generate placeholders for GraphQL API calls that have the same method signature as the real API calls. It is valuable for testing and development purposes when certain API calls are unavailable or when you want to test the application without sending real API requests. Despite this, you can still maintain the ability to interact with the actual GraphQL server.

## Features

- [x] Acts as a reverse proxy for the actual GraphQL server without stubbing implemented fields. 
- [x] Configurable stub creation for specific root fields.


## Setup environment
1. To set up the environment, you need to have Go installed on your machine. You could install Go using Homebrew on MacOS.
    ```shell
    # Install Go
    $ brew install go

    # Sync the go dependencies 
    $ go mod tidy 
    ```

1. In the root directory of the project, create a copy of env.yaml.example and name it env.yaml. Fill in the values for the environment variables in the file.

## Run the server
```shell
$ make run
```
Once the server is running, you can access the GraphQL playground at `http://localhost:9002/` to test the API calls.
You could also point your application to `http://localhost:9002/graphql/` to make the API calls.

## Configure stubs

To configure the stubs, you could simply make change to `gql/schema.resolvers.go` file. The file contains the stubs for the GraphQL API calls. You could change the return values of the functions to whatever you want to return.

## Sync the latest schema from GraphQL 
```shell
# Generate the Go code from the GraphQL schema. 
$ make generate
```

## Format the code
```shell
$ make fmt
```

## Run the tests
```shell
$ make test
```
