package paper

import (
	"log"

	"github.com/rydelll/papermc/internal/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Paper server",
	Long: `Download a Minecraft Paper server. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadPaper,
}

func init() {
	paperCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadPaper(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	err := download.Download(download.Paper, version)
	if err != nil {
		log.Fatal(err)
	}
}
