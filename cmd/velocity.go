package cmd

import (
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
	Long: `Download a Minecraft Velocity proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	// TODO
}
