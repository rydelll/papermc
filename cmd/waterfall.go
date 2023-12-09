package cmd

import (
	"log"

	"github.com/rydelll/papermc/client"
	"github.com/spf13/cobra"
)

var waterfallCmd = &cobra.Command{
	Use:   "waterfall",
	Short: "Minecraft Waterfall proxy",
	Long: `Waterfall is an upgraded BungeeCord Minecraft proxy, offering full
compatibility with improvements to performance and stability.`,
}

var waterfallDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Waterfall proxy",
	Long: `Download a Minecraft Waterfall proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadWaterfall,
}

func init() {
	RootCmd.AddCommand(waterfallCmd)
	waterfallCmd.AddCommand(waterfallDownloadCmd)
	waterfallDownloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadWaterfall(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	c := client.NewClient()
	var err error

	if version == "latest" {
		version, err = c.Waterfall.LatestVersion()
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.Waterfall.LatestBuild(version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Waterfall.Download(info)
	if err != nil {
		log.Fatal(err)
	}
}
