package client

import (
	"log"
	"vague/internal/client"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Send count to server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res struct {
			Count int `json:"count"`
		}

		if err := conn.Execute(cmd.Context(), "function_count", nil, &res); err != nil {
			return err
		}

		log.Println(res.Count)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(countCmd)
}
