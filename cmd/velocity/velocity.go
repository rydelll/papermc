package velocity

import (
	"github.com/rydelll/papermc/cmd"
	"github.com/spf13/cobra"
)

var velocityCmd = &cobra.Command{
	Use:   "velocity",
	Short: "Minecraft Velocity proxy",
	Long: `Velocity is the modern, high-performance proxy for Minecraft. Designed with
performance and stability in mind, it’s a full alternative to Waterfall with
its own diverse plugin ecosystem.`,
}

func init() {
	cmd.RootCmd.AddCommand(velocityCmd)
}
