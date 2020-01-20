package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var reposCmd = &cobra.Command{
	Use:   "repositories [repository]",
	Short: "List repositories in registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		repositories, err := registryClient.RepositoryList(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range repositories {
			fmt.Println(name)
		}
	},
}
