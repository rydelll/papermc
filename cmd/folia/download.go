package folia

import (
	"log"

	"github.com/rydelll/papermc/internal/download"
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

	err := download.Download(download.Folia, version)
	if err != nil {
		log.Fatal(err)
	}
}
