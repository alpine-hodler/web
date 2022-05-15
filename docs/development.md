# Development

- [Dependencies](#dependencies)
	- [Bazel](#bazel)
	- [Docker](#docker)
	- [Go](#go)
- [Build](#build)
- [Packages](#packages)
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

## Packages

- [Coinbase Cloud API (Coinbase Pro)](https://github.com/alpine-hodler/sdk/blob/main/pkg/coinbasepro/README.md)

## Extending the API

To extend the api, you'll need to add a new schema to the `scripts/meta/schema` directory, following the template layed out in `scripts/meta/schema/schema.json`.  Then build the metadata:
```
make build-meta
```
