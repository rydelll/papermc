package velocity

import (
	"log"

	"github.com/rydelll/papermc/pkg/papermc"
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

	c := papermc.NewClient()
	var err error

	if version == "latest" {
		version, err = c.ProjectVersion.GetLatest(papermc.Velocity)
		if err != nil {
			log.Fatal(err)
		}
	}

	info, err := c.ProjectBuild.GetLatest(papermc.Velocity, version)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ProjectDownload.Download(papermc.Velocity, info)
	if err != nil {
		log.Fatal(err)
	}
}
