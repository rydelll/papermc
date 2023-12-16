package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "papermc",
	Short: "Download, install, and setup Minecraft PaperMC products.",
	Long: `Download, install, and setup Minecraft PaperMC products. 
PaperMC products include Minecraft servers and proxies. By default the latest
version will be installed unless. A specific version can be selected as well.`,
}

func Execute() int {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}

	return 0
}
