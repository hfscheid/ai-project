package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestListCmd defines the 'test list' command for Disco CLI
func (d *Disco) newTestListCmd() *cobra.Command {
    return cmd.NewCmd("list").
        WithDescription("Command for listing available tests").
        WithExample("List available tests", "test list").
        NoArgs(d.listTests)
}

func (d *Disco) listTests(_ context.Context, c *cobra.Command) error {
    tests := ""
    for _, test := range d.tests.TestCases {
        tests += fmt.Sprintf("%s\n", test.Name)
    }
    if len(tests) == 0 {
        fmt.Println("No tests found")
        return nil
    }
    fmt.Println(tests)
    return nil
}
