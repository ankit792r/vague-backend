package client

import (
	"log"
	"vague/internal/client"
	"vague/internal/runtime"

	"github.com/spf13/cobra"
)

var listFunctionsCmd = &cobra.Command{
	Use:   "list",
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

		if err := conn.Execute(cmd.Context(), "function_list", nil, &fnamesResponse); err != nil {
			return err
		}

		log.Println(fnamesResponse.Names)

		return nil
	},
}

var existsFunctionCmd = &cobra.Command{
	Use:   "exists <name>",
	Short: "Send exists to server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res struct {
			Exist bool `json:"exist"`
		}

		fargs := runtime.FunctionNameArgs{Name: args[0]}

		if err := conn.Execute(cmd.Context(), "function_exist", fargs, &res); err != nil {
			return err
		}

		log.Println(res.Exist)

		return nil
	},
}

var descFunctionCmd = &cobra.Command{
	Use:   "desc <name>",
	Short: "Describe a function and its arguments",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res runtime.HandlerMeta

		fargs := runtime.FunctionNameArgs{Name: args[0]}

		if err := conn.Execute(cmd.Context(), "function_describe", fargs, &res); err != nil {
			return err
		}

		// log.Println(res)

		for _, a := range res.Args {
			log.Printf("  %s (%s, required=%t): %s", a.Name, a.Type, a.Required, a.Desc)
		}

		return nil
	},
}

func init() {
	ClientCmd.AddCommand(listFunctionsCmd)
	ClientCmd.AddCommand(existsFunctionCmd)
	ClientCmd.AddCommand(descFunctionCmd)
}
