package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"vague/internal/client"
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

		var res json.RawMessage
		if err := conn.Execute(ctx, name, rest, &res); err != nil {
			return err
		}

		fmt.Println(res)
	}
}
