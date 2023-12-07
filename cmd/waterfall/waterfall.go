package waterfall

import (
	"github.com/rydelll/papermc/cmd"
	"github.com/spf13/cobra"
)

var waterfallCmd = &cobra.Command{
	Use:   "waterfall",
	Short: "Minecraft Waterfall proxy",
	Long: `Waterfall is an upgraded BungeeCord Minecraft proxy, offering full
compatibility with improvements to performance and stability.`,
}

func init() {
	cmd.RootCmd.AddCommand(waterfallCmd)
}
