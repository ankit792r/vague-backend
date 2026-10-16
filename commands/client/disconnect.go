package client

import (
	"fmt"

	"github.com/spf13/cobra"
)

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from the vague server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Disconnecting from vague server...")
		return nil
	},
}

func init() {
	ClientCmd.AddCommand(disconnectCmd)
}
