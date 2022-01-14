package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alpine-hodler/sdk/graph"
	"github.com/alpine-hodler/sdk/graph/generated"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	startCmd.Flags().StringVarP(&PORT, "port", "p", "", "Port to open the server on")

	startCmd.MarkFlagRequired("port")

	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the graphql server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.DebugLevel)
		startServer()
	},
}

func startServer() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
