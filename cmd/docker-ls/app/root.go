package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

// Execute runs the command
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
