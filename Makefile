chmod +x ./scripts/build_bazel.sh
chmod +x ./scripts/build_meta.sh

go mod tidy

./scripts/build_meta.sh
./scripts/build_bazel.sh

gqlgen generate
