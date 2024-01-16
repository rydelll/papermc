package cmd

import (
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
	Long: `Download a Minecraft Waterfall proxy. By default the latest version will be
installed unless. A specific version can be selected as well.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	// TODO
}
