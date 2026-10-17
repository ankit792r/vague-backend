package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/handler"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Send ping to server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res handler.PingResponse

		if err := conn.Execute(cmd.Context(), "ping", nil, &res); err != nil {
			return err
		}

		log.Println(res.Message)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(pingCmd)
}
