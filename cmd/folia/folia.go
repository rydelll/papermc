package folia

import (
	"github.com/rydelll/papermc/cmd"
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

func init() {
	cmd.RootCmd.AddCommand(foliaCmd)
}
