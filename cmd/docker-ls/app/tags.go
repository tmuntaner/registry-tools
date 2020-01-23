package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var tagsCmd = &cobra.Command{
	Use:   "tags [repository]",
	Short: "List tags for a repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		registryUrl, err := cmd.Flags().GetString("registry")
		image, err := registryParser.GunToImage(args[0], registryUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if registryUrl != "" {
			image.Host = registryUrl
		}

		tags, err := registryClient.TagList(image)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, name := range tags {
			fmt.Println(name)
		}
	},
}
