package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	// flags
	PORT string

	rootCmd = &cobra.Command{
		Use: "graphsrv",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
