package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vague/commands"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := commands.RootCmd.ExecuteContext(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
