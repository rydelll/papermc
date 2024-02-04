package cli

import (
	"github.com/rydelll/papermc"
	"github.com/spf13/cobra"
)

var travertineCmd = &cobra.Command{
	Use:   "travertine",
	Short: "Minecraft Travertine proxy",
	Long: `Travertine is a Minecraft proxy forked from Waterfall with additional
protocols that aims to support older client versions.`,
}

var travertineDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Travertine proxy",
	Long: `Download a Minecraft Travertine proxy. The latest version will be
downloaded if not specified.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := papermc.NewClient()
		var err error

		if version == "" {
			version, err = c.Travertine.Version.GetLatest()
			if err != nil {
				return err
			}
		}

		var info papermc.ProjectInfo
		if build == 0 {
			info, err = c.Travertine.Build.GetLatest(version)
		} else {
			info, err = c.Travertine.Build.Get(version, build)
		}
		if err != nil {
			return err
		}

		err = c.Travertine.JAR.Download(info.Version, info.Build, info.JAR)
		if err != nil {
			return err
		}

		err = c.Travertine.JAR.ValidateChecksum(info.Checksum)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(travertineCmd)
	travertineCmd.AddCommand(travertineDownloadCmd)

	travertineDownloadCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "project version")
	travertineDownloadCmd.PersistentFlags().IntVarP(&build, "build", "b", 0, "project version build number")
}
