package cmd

import (
	"github.com/rydelll/papermc/client"
	"github.com/spf13/cobra"
)

var paperCmd = &cobra.Command{
	Use:   "paper",
	Short: "Minecraft Paper server",
	Long: `Paper is a Minecraft game server based on Spigot, designed to greatly improve
performance and offer more advanced features and API.`,
}

var paperDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a Minecraft Paper server",
	Long: `Download a Minecraft Paper server. The latest version will be
downloaded if not specified.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.New()
		var err error

		if version == "" {
			version, err = c.Paper.Version.GetLatest()
			if err != nil {
				return err
			}
		}

		var info client.ProjectInfo
		if build == 0 {
			info, err = c.Paper.Build.GetLatest(version)
		} else {
			info, err = c.Paper.Build.Get(version, build)
		}
		if err != nil {
			return err
		}

		err = c.Paper.JAR.Download(info.Version, info.Build, info.JAR)
		if err != nil {
			return err
		}

		err = c.Paper.JAR.ValidateChecksum(info.Checksum)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(paperCmd)
	paperCmd.AddCommand(paperDownloadCmd)

	paperDownloadCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "project version")
	paperDownloadCmd.PersistentFlags().IntVarP(&build, "build", "b", 0, "project version build number")
}
