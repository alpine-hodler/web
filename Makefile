PKGS=$(shell scripts/list_pkgs.sh ./pkg)

default:
	scripts/build_meta.sh
	scripts/build_bazel.sh

build-bazel:
	scripts/build_bazel.sh

build-meta:
	scripts/build_meta.sh

generate:
	docker-compose -f "meta.docker-compose.yaml" run test_generate
	docker-compose -f "meta.docker-compose.yaml" run generate
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

list_pkgs:
	echo $(PKGS)

start-graphql:
	docker-compose -f "graphql.docker-compose.yaml" up -d --build
