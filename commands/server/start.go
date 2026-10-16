package server

import (
	"fmt"
	"log"
	"vague/internal/process"
	intServ "vague/internal/server"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the vague server",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Starting Vague Server")
		lis, err := process.ListenV2()
		if err != nil {
			return err
		}

		log.Printf("Server Listening... \n")
		if err := intServ.New().Serve(cmd.Context(), lis); err != nil {
			return fmt.Errorf("server: %w", err)
		}

		return nil
	},
}

func init() {
	ServerCmd.AddCommand(startCmd)
}
