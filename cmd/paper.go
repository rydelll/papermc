package cmd

import (
	"github.com/spf13/cobra"
)

var paperCmd = &cobra.Command{
	Use:   "paper",
	Short: "Minecraft Paper server",
	Long: `Paper is a Minecraft game server based on Spigot, designed to greatly improve
performance and offer more advanced features and API.`,
}

var paperDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Paper server",
	Long: `Download a Minecraft Paper server. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	// TODO
}
