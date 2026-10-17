package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/process"

	"github.com/spf13/cobra"
)

var listFunctionsCmd = &cobra.Command{
	Use:   "fn_list",
	Short: "Send List fn to server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var fnamesResponse struct {
			Names []string `json:"names"`
		}

		par := process.ExecuteParams{
			Name: "function_list",
			Args: []string{},
		}

		if err := conn.Call(cmd.Context(), process.ActionExecute, par, &fnamesResponse); err != nil {
			return err
		}

		log.Println(fnamesResponse.Names)

		return nil
	},
}

var existsFunctionCmd = &cobra.Command{
	Use:   "exists",
	Short: "Send exists to server",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res struct {
			Exist bool `json:"exist"`
		}

		par := process.ExecuteParams{
			Name: "function_exist",
			Args: args,
		}

		if err := conn.Call(cmd.Context(), process.ActionExecute, par, &res); err != nil {
			return err
		}

		log.Println(res.Exist)

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(listFunctionsCmd)
	ClientCmd.AddCommand(existsFunctionCmd)
}
