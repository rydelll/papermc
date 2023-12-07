package waterfall

import (
	"log"

	"github.com/rydelll/papermc/internal/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Waterfall proxy",
	Long: `Download a Minecraft Waterfall proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadWaterfall,
}

func init() {
	waterfallCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadWaterfall(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	err := download.Download(download.Waterfall, version)
	if err != nil {
		log.Fatal(err)
	}
}
