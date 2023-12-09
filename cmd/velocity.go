package cmd

import (
	"log"

	"github.com/rydelll/papermc/client"
	"github.com/spf13/cobra"
)

var velocityCmd = &cobra.Command{
	Use:   "velocity",
	Short: "Minecraft Velocity proxy",
	Long: `Velocity is the modern, high-performance proxy for Minecraft. Designed with
performance and stability in mind, it’s a full alternative to Waterfall with
its own diverse plugin ecosystem.`,
}

var velocityDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Velocity proxy",
	Long: `Download a Minecraft Velocity proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadVelocity,
}

func init() {
	RootCmd.AddCommand(velocityCmd)
	velocityCmd.AddCommand(velocityDownloadCmd)
	velocityDownloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadVelocity(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	c := client.NewClient()
	var err error

	if version == "latest" {
		version, err = c.Velocity.LatestVersion()
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.Velocity.LatestBuild(version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Velocity.Download(info)
	if err != nil {
		log.Fatal(err)
	}
}
