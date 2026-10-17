package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/process"

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

		var res string

		par := process.ExecuteParams{
			Name: "ping",
			Args: []string{},
		}

		if err := conn.Call(cmd.Context(), process.ActionExecute, par, &res); err != nil {
			return err
		}

		log.Println("Successfully got res", res)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(pingCmd)
}
