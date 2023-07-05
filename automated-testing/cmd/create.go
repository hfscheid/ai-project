package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var (
    create = &cobra.Command {
        Use:    "create",
        Short:  "creates a software router",
        Run:    create,
    }
)

func add(ccmd *cobra.Command, args []string) {
    if len(args) == 0 {
        fmt.
    }
