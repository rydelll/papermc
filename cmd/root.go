package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	version string
	build   int
)

var rootCmd = &cobra.Command{
	Use:   "papermc",
	Short: "Download, install, and setup Minecraft PaperMC products.",
	Long: `Download, install, and setup Minecraft PaperMC products. 
PaperMC products include Minecraft servers and proxies.`,
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
