package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tmuntaner/registry-tools/internal/registry"
	"os"
)

var reposCmd = &cobra.Command{
	Use:   "repositories [registry]",
	Short: "List repositories in registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		repositories, err := registry.RepositoryList(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range repositories {
			fmt.Println(name)
		}
	},
}
