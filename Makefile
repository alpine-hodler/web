PKGS=$(shell scripts/list_pkgs.sh ./pkg)

default:
	go mod tidy
	scripts/build_meta.sh
	scripts/build_bazel.sh

build-bazel:
	scripts/build_bazel.sh

build-meta:
	scripts/build_meta.sh
	gqlgen generate

list_pkgs:
	echo $(PKGS)

start-graphql:
	docker-compose -f "graphql.docker-compose.yaml" up -d --build
