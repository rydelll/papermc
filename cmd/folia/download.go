package folia

import (
	"log"

	"github.com/rydelll/papermc/pkg/papermc"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Folia server",
	Long: `Download a Minecraft Folia server. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadFolia,
}

func init() {
	foliaCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadFolia(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	c := papermc.NewClient()
	var err error

	if version == "latest" {
		version, err = c.ProjectVersion.GetLatest(papermc.Folia)
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.ProjectBuild.GetLatest(papermc.Folia, version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ProjectDownload.Download(papermc.Folia, info)
	if err != nil {
		log.Fatal(err)
	}
}
