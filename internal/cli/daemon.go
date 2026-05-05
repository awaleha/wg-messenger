package cli

import (
	"fmt"
	"os"

	"github.com/awaleha/wg-messenger/internal/daemon"
	"github.com/spf13/cobra"
)

var listenAddr string

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the daemon",
	Run: func(cmd *cobra.Command, args []string) {

		d, err := daemon.NewDaemon()
		if err != nil {
			fmt.Println("Failed create new Daemon")
			os.Exit(1)
		}
		if err := daemon.Run(d, listenAddr); err != nil {
			fmt.Println("Failed to run listener")
			os.Exit(1)
		}

	},
}

func init() {
	daemonCmd.Flags().StringVar(
		&listenAddr,
		"listen",
		"127.0.0.1:7777",
		"address to listen for messages",
	)

	rootCmd.AddCommand(daemonCmd)
}
