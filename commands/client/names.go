package client

import (
	"log"
	"vague/internal/client"

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

		if err := conn.Execute(cmd.Context(), "function_list", nil, &res); err != nil {
			return err
		}

		log.Println(res.Names)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(nameCmd)
}
