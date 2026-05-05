package cli

import (
	"github.com/awaleha/wg-messenger/internal/transport"
	"github.com/spf13/cobra"
)

var sendAddr string

var sendCmd = &cobra.Command{
	Use:   "send \"message\"",
	Short: "Send a message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		msg := args[0]
		transport.Client(sendAddr, msg)

	},
}

func init() {
	sendCmd.Flags().StringVar(
		&sendAddr,
		"addr",
		"127.0.0.1:7777",
		"target address to send the message to",
	)

	rootCmd.AddCommand(sendCmd)
}
