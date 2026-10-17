package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"vague/internal/client"
	"vague/internal/runtime"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered functions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res json.RawMessage
		if err := conn.Execute(cmd.Context(), "function_list", nil, &res); err != nil {
			return err
		}

		return printJSON(cmd.OutOrStdout(), res)
	},
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count the registered functions",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		var res json.RawMessage
		if err := conn.Execute(cmd.Context(), "function_count", nil, &res); err != nil {
			return err
		}

		return printJSON(cmd.OutOrStdout(), res)
	},
}

var existsCmd = &cobra.Command{
	Use:   "exists <name>",
	Short: "Check whether a function is registered",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		fargs := runtime.FunctionNameArgs{Name: args[0]}

		var res json.RawMessage
		if err := conn.Execute(cmd.Context(), "function_exist", fargs, &res); err != nil {
			return err
		}

		return printJSON(cmd.OutOrStdout(), res)
	},
}

var describeCmd = &cobra.Command{
	Use:   "describe <name>",
	Short: "Describe a function and its arguments",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		fargs := runtime.FunctionNameArgs{Name: args[0]}

		var res json.RawMessage
		if err := conn.Execute(cmd.Context(), "function_describe", fargs, &res); err != nil {
			return err
		}

		return printJSON(cmd.OutOrStdout(), res)
	},
}

func printJSON(w io.Writer, raw json.RawMessage) error {
	var out bytes.Buffer
	if err := json.Indent(&out, raw, "", "  "); err != nil {
		return fmt.Errorf("format result: %w", err)
	}

	_, err := fmt.Fprintln(w, out.String())
	return err
}

func init() {
	ClientCmd.AddCommand(listCmd)
	ClientCmd.AddCommand(countCmd)
	ClientCmd.AddCommand(existsCmd)
	ClientCmd.AddCommand(describeCmd)
}
