package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := NewRootCmd().ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]: %v\n", err)
		os.Exit(1)
	}
}
