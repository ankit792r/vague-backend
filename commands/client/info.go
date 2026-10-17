package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/process"

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

		par := process.ExecuteParams{
			Name: "info",
			Args: []string{},
		}

		if err := conn.Call(cmd.Context(), process.ActionExecute, par, &res); err != nil {
			return err
		}

		log.Println(res)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(infoCmd)
}
