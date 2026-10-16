package commands

import (
	"github.com/spf13/cobra"

	"vague/commands/client"
	"vague/commands/server"
)

var RootCmd = &cobra.Command{
	Use:           "vague",
	Short:         "A Vim-inspired text editor",
	Args:          cobra.ArbitraryArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(client.ClientCmd)
}
