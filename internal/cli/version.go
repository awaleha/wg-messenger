package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wgmsg",
	Long:  `All software has versions. This is WG-Messenger's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WG-Messenger V.0.1")
	},
}
