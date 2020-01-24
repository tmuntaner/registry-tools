package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tmuntaner/registry-tools/internal/parser"
	"github.com/tmuntaner/registry-tools/internal/registry"
	"os"
)

var tagsCmd = &cobra.Command{
	Use:   "tags [repository]",
	Short: "List tags for a repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		registryURL, err := cmd.Flags().GetString("registry")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		image, err := parser.GunToImage(args[0], registryURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if registryURL != "" {
			image.Host = registryURL
		}

		tags, err := registry.TagList(image)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range tags {
			fmt.Println(name)
		}
	},
}
