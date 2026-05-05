package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wgmsg",
	Short: "WireGuard-backed terminal messenger",
	Long:  "wgmsg is a peer-to-peer terminal messenger that sends messages over a WireGuard network.",
}

func Execute() error {
	fmt.Println("Starting WireGuard Messenger")
	return rootCmd.Execute()
}
