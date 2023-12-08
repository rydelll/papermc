package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "papermc",
	Short: "download, install, and setup Minecraft PaperMC products.",
	Long: `Download, install, and setup Minecraft PaperMC products. 
PaperMC products include Minecraft servers and proxies. By default the latest
version will be installed unless. A specific version can be selected as well.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
