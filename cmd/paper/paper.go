package paper

import (
	"github.com/rydelll/papermc/cmd"
	"github.com/spf13/cobra"
)

var paperCmd = &cobra.Command{
	Use:   "paper",
	Short: "Minecraft Paper server",
	Long: `Paper is a Minecraft game server based on Spigot, designed to greatly improve
performance and offer more advanced features and API.`,
}

func init() {
	cmd.RootCmd.AddCommand(paperCmd)
}
