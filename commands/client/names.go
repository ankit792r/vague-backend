package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/process"

	"github.com/spf13/cobra"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Send name to server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res struct {
			Names []string `json:"names"`
		}

		par := process.ExecuteParams{
			Name: "function_names",
			Args: []string{},
		}

		if err := conn.Call(cmd.Context(), process.ActionExecute, par, &res); err != nil {
			return err
		}

		log.Println(res.Names)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(nameCmd)
}
