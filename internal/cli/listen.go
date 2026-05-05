package cli

import (
	"github.com/awaleha/wg-messenger/internal/transport"
	"github.com/spf13/cobra"
)

var listenAddr string

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {

		transport.Server(listenAddr)

	},
}

func init() {
	listenCmd.Flags().StringVar(
		&listenAddr,
		"addr",
		"127.0.0.1:7777",
		"address to listen for messages",
	)

	rootCmd.AddCommand(listenCmd)
}
