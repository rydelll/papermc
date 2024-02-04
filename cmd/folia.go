package cmd

import (
	"github.com/rydelll/papermc/client"
	"github.com/spf13/cobra"
)

var foliaCmd = &cobra.Command{
	Use:   "folia",
	Short: "Minecraft Folia server",
	Long: `Folia is a new fork of the Paper Minecraft server that adds regionized
multithreading to the server. Folia is designed to address the constant
bottleneck of the server running on a single thread causing performance issues.
It is not a drop-in replacement for Paper as it breaks most public plugins.`,
}

var foliaDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Folia server",
	Long: `Download a Minecraft Folia server. The latest version will be
downloaded if not specified.`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		var err error

		if version == "" {
			version, err = c.Folia.Version.GetLatest()
			if err != nil {
				return err
			}
		}

		var info client.ProjectInfo
		if build == 0 {
			info, err = c.Folia.Build.GetLatest(version)
		} else {
			info, err = c.Folia.Build.Get(version, build)
		}
		if err != nil {
			return err
		}

		err = c.Folia.JAR.Download(info.Version, info.Build, info.JAR)
		if err != nil {
			return err
		}

		err = c.Folia.JAR.ValidateChecksum(info.Checksum)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(foliaCmd)
	foliaCmd.AddCommand(foliaDownloadCmd)

	foliaDownloadCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "project version")
	foliaDownloadCmd.PersistentFlags().IntVarP(&build, "build", "b", 0, "project version build number")
}
