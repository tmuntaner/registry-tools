package app

import (
	"fmt"
	"github.com/spf13/cobra"
	registry "github.com/tmuntaner/registry-tools/internal/registry"
	"os"
)

var rootCmd = &cobra.Command{}
var registryClient registry.Client

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	tagsCmd.Flags().StringP("registry", "r", "", "URL for registry")
	rootCmd.AddCommand(reposCmd)
	rootCmd.AddCommand(tagsCmd)
}
