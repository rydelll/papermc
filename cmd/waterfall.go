package cmd

import (
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
	Long: `Download a Minecraft Waterfall proxy. The latest version will be
downloaded if not specified.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		var err error

		if version == "" {
			version, err = c.Waterfall.Version.GetLatest()
			if err != nil {
				return err
			}
		}

		var info client.ProjectInfo
		if build == 0 {
			info, err = c.Waterfall.Build.GetLatest(version)
		} else {
			info, err = c.Waterfall.Build.Get(version, build)
		}
		if err != nil {
			return err
		}

		err = c.Waterfall.JAR.Download(info.Version, info.Build, info.JAR)
		if err != nil {
			return err
		}

		err = c.Waterfall.JAR.ValidateChecksum(info.Checksum)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(waterfallCmd)
	waterfallCmd.AddCommand(waterfallDownloadCmd)

	waterfallDownloadCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "project version")
	waterfallDownloadCmd.PersistentFlags().IntVarP(&build, "build", "b", 0, "project version build number")
}
