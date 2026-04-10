package main

import (
	"fmt"
	"os"

	"scout-cli/internal/usecase"
)

func main() {
	application := usecase.NewApp(os.Stdout, os.Stderr)

	if err := application.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}