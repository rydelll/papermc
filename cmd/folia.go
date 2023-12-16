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
	Long: `Download a Minecraft Folia server. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()
		var err error

		if version == "latest" {
			version, err = c.Folia.LatestVersion()
			if err != nil {
				return err
			}
		}

		info, err := c.Folia.LatestBuild(version)
		if err != nil {
			return err
		}

		err = c.Folia.Download(info)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(foliaCmd)
	foliaCmd.AddCommand(foliaDownloadCmd)
	foliaDownloadCmd.Flags().StringVarP(&version, "version", "v", "latest", "version to download")
}
