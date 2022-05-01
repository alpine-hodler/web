.PHONY: default
default:
	go mod tidy
	scripts/build_meta.sh
	scripts/build_bazel.sh
