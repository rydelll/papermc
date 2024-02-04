package cli

import (
	"github.com/rydelll/papermc"
	"github.com/spf13/cobra"
)

var velocityCmd = &cobra.Command{
	Use:   "velocity",
	Short: "Minecraft Velocity proxy",
	Long: `Velocity is the modern, high-performance proxy for Minecraft. Designed with
performance and stability in mind, it's a full alternative to Waterfall with
its own diverse plugin ecosystem.`,
}

var velocityDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Velocity proxy",
	Long: `Download a Minecraft Velocity proxy. The latest version will be
downloaded if not specified.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := papermc.NewClient()
		var err error

		if version == "" {
			version, err = c.Velocity.Version.GetLatest()
			if err != nil {
				return err
			}
		}

		var info papermc.ProjectInfo
		if build == 0 {
			info, err = c.Velocity.Build.GetLatest(version)
		} else {
			info, err = c.Velocity.Build.Get(version, build)
		}
		if err != nil {
			return err
		}

		err = c.Velocity.JAR.Download(info.Version, info.Build, info.JAR)
		if err != nil {
			return err
		}

		err = c.Velocity.JAR.ValidateChecksum(info.Checksum)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(velocityCmd)
	velocityCmd.AddCommand(velocityDownloadCmd)

	velocityDownloadCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "project version")
	velocityDownloadCmd.PersistentFlags().IntVarP(&build, "build", "b", 0, "project version build number")
}
