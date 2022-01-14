# the mod cache can cause some goofy warnings if its not cleaned before we build
go clean -modcache

bazel clean
bazel run //:gazelle

# Whenever we add new go dependencies we need to update the repos through bazel
bazel run //:gazelle-update-repos

bazel build //...

# # regenerate graphql stuff
# go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate
