package client

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"vague/internal/client"
	"vague/internal/process"

	"github.com/spf13/cobra"
)

var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Start command loop",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		return startLoop(cmd.Context(), conn)
	},
}

func startLoop(ctx context.Context, conn *client.Client) error {
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !sc.Scan() {
			break
		}

		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}

		if fields[0] == "exit" || fields[0] == "quit" {
			return nil
		}

		params := process.ExecuteParams{
			Name: fields[0],
			Args: fields[1:],
		}

		var res json.RawMessage
		if err := conn.Execute(ctx, params, &res); err != nil {
			if errors.Is(err, client.ErrClosed) {
				return err
			}

			fmt.Println("error:", err)
			continue
		}

		fmt.Println(res)

		// printResult(res)
	}

	if err := sc.Err(); err != nil {
		return fmt.Errorf("read standard input: %w", err)
	}

	return nil
}

// printResult prints a string result as-is and anything else as indented JSON.
func printResult(res json.RawMessage) {
	if len(res) == 0 {
		return
	}

	var s string
	if err := json.Unmarshal(res, &s); err == nil {
		fmt.Println(s)
		return
	}

	var v any
	if err := json.Unmarshal(res, &v); err != nil {
		fmt.Println(string(res))
		return
	}

	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(string(res))
		return
	}

	fmt.Println(string(pretty))
}

func init() {
	ClientCmd.AddCommand(loopCmd)
}
