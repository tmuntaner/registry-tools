package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var tagsCmd = &cobra.Command{
	Use:   "tags [repo]",
	Short: "List tags for a repository",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		image, err := registryParser.GunToImage(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
