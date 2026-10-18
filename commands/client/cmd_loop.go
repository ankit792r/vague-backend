package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"vague/internal/client"

	"github.com/spf13/cobra"
)

func init() {
	ClientCmd.AddCommand(loopCmd)
}

var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Start Command Loop",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := client.Connect()
		if err != nil {
			return err
		}

		defer conn.Close()

		return startLoop(cmd.Context(), conn, cmd.InOrStdin(), cmd.OutOrStdout(), cmd.ErrOrStderr())
	},
}

func startLoop(ctx context.Context, conn *client.Client, in io.Reader, out, errOut io.Writer) error {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(errOut, "> ")

		if !scanner.Scan() {
			fmt.Fprintln(errOut)
			return scanner.Err()
		}

		if err := ctx.Err(); err != nil {
			return err
		}

		line := strings.TrimSpace(scanner.Text())

		name, rest, _ := strings.Cut(line, " ")
		rest = strings.TrimSpace(rest)

		var args any
		if rest != "" {
			if !json.Valid([]byte(rest)) {
				fmt.Fprintln(errOut, "error: args must be a JSON object")
				continue
			}
			args = json.RawMessage(rest)
		}

		var res json.RawMessage
		if err := conn.Execute(ctx, name, args, &res); err != nil {
			fmt.Fprintln(errOut, "error:", err)
			continue
		}

		fmt.Println(string(res))
	}
}
