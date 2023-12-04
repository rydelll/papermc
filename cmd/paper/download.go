package paper

import (
	"log"

	"github.com/rydelll/papermc/internal/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: downloadPaper,
}

func init() {
	paperCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("version", "v", "latest", "version to download")
}

func downloadPaper(cmd *cobra.Command, args []string) {
	version, _ := cmd.Flags().GetString("version")

	err := download.Download(download.Paper, version)
	if err != nil {
		log.Fatal(err)
	}
}
