package cmd

import (
	"log"

	"github.com/rydelll/papermc/client"
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
	Run: downloadPaper,
}

func init() {
	RootCmd.AddCommand(paperCmd)
	paperCmd.AddCommand(paperDownloadCmd)
	paperDownloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadPaper(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	c := client.NewClient()
	var err error

	if version == "latest" {
		version, err = c.Paper.LatestVersion()
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.Paper.LatestBuild(version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Paper.Download(info)
	if err != nil {
		log.Fatal(err)
	}
}
