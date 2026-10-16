package client

import (
	"fmt"
	"log"
	"vague/internal/client"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the vague server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Connecting to vague server...")
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		log.Println("Connected successfully", conn)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(connectCmd)
}
