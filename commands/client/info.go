package client

import (
	"log"
	"vague/internal/client"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Send info to server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res string

		if err := conn.Execute(cmd.Context(), "info", nil, &res); err != nil {
			return err
		}

		log.Println(res)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(infoCmd)
}
