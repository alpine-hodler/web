# Development

- [Dependencies](#dependencies)
	- [Bazel](#bazel)
	- [Docker](#docker)
	- [Go](#go)
- [Build](#build)
- [Setting up Auth .env file](#setting-up-auth-env-file)
- [Extending the API](#extending-the-api)

## Dependencies
To develop locally you need to install three dependencies:

### Bazel

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

To install, following the intructions [here](https://docs.bazel.build/versions/4.2.2/bazel-overview.html#how-do-i-use-bazel)

If you're on macOS, [you can install Bazel via Homebrew](https://docs.bazel.build/versions/4.2.2/install-os-x.html#step-2-install-bazel-via-homebrew):

```
brew install bazel
```

### Docker

https://docs.docker.com/get-docker/

### Go

https://go.dev/doc/install


## Build

To build simply run

```
make
```

## Setting up Auth .env file

Somewhere on your local machine you'll need to add a `.env` file for auth between the different API.

### Coinbase

You will need to create an account for the [Coinbase Pro Sandbox]("https://api-public.sandbox.exchange.coinbase.com") and setup a new API key for the `Default Portfolio` with `Vilew/Trade/Transfer` permissions.  Then populate the following data in `pkg/coinbase/.simple-test.env`:
```.env
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

Note that `pkg/coinbase/.simple-test.env` is an ignored file and should not be commitable to the repository.  The Coinbase Pro Sandbox can be accessed [here](https://public.sandbox.pro.coinbase.com).

## Extending the API

To extend the api, you'll need to add a new schema to the `scripts/meta/schema` directory, following the template layed out in `scripts/meta/schema/schema.json`.  Then build the metadata:
```
make build-meta
```
