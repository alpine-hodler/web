# Development

## Dependencies
To develop locally you need to install three dependencies:

### Install Bazel

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

To install, following the intructions [here](https://docs.bazel.build/versions/4.2.2/bazel-overview.html#how-do-i-use-bazel)

If you're on macOS, [you can install Bazel via Homebrew](https://docs.bazel.build/versions/4.2.2/install-os-x.html#step-2-install-bazel-via-homebrew):

```
brew install bazel
```

### Install Docker

https://docs.docker.com/get-docker/

We will primarily use docker for generating go and graphqls files from script/meta/schema files.

### Install Go

https://go.dev/doc/install


## Build

To build simply run

```
./Makefile
```

## Setting up Auth .env file

Somewhere on your local machine you'll need to add a `.env` file for auth between the different API.

### Coinbase

The coinbase .env file should look like this:
```.env
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

The usage looks like this:
```go
// initialize the client connection
client := coinbase.NewClientEnv("/path/to/auth/.env")

// then use it for something
accounts := client.Accounts()
fmt.Printf("%+v\n", accounts)
```

## Running GraphQL Server

To start the graphql server go to the `cmd/graphql` directory and generate the graphql binary with go build. You'll also need to make sure the proper auth data is in some /path/to/auth/.env, see coinbase section for example. Then run

```
./graphql start --port=8080 --env=/path/to/auth/.env
```

and visit http://localhost:8080/

## Extending the API

To extend the api, you'll need to add a new schema to the `scripts/meta/schema` directory, following the template layed out in `scripts/meta/schema/schema.json`.  Then simply run the generate method:

```
docker-compose -f "meta.docker-compose.yaml" run generate
```

This will add data to the following:

- protomodel: A new protomodel will be created for the data, this is a purely generated go model
- model: A new model will the created for the data.  This model extends the protomodel struct and is open to extension
- model accessors: A set of functions to access struct and non-native go types, primiarly used by the graphql resolver.  The file structure lives in the pkg/model directory.
- internal/graph/schema: A new graphql type will be added as a 1-1 relationship with the generated protomodel
- endpoint accessors: A set of endpoint accessor will be added to the pkg/{api} package

You'll also need to run

```
gqlgen generate
```

to generate the graphql structures.

### Testing

You can test the meta package with

```
docker-compose -f "meta.docker-compose.yaml" run test_generate
```
