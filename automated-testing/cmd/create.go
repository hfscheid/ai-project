package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
    "os"
)

var (
    createcmd = &cobra.Command {
        Use:    "create",
        Short:  "creates a software router",
        Run:    create,
    }
)

func create(ccmd *cobra.Command, args []string) {
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr,
                     "at least one type router type must be specified")
        return
    }
    for _, arg := range args {
        fmt.Println("Adding router of type", arg, "to pool")
    }
}
