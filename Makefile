.PHONY: default
default:
	go mod tidy
	scripts/build_meta.sh
	scripts/build_bazel.sh

build-bazel:
	scripts/build_bazel.sh

build-meta:
	scripts/build_meta.sh
	gqlgen generate

start-graphql:
	docker-compose -f "graphql.docker-compose.yaml" up -d --build
