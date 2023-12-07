package velocity

import (
	"log"

	"github.com/rydelll/papermc/internal/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Velocity proxy",
	Long: `Download a Minecraft Velocity proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: downloadVelocity,
}

func init() {
	velocityCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadVelocity(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	err := download.Download(download.Velocity, version)
	if err != nil {
		log.Fatal(err)
	}
}
