package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestListCmd defines the 'test list' command for Disco CLI
func (d *Disco) newTestListCmd() *cobra.Command {
    return cmd.NewCmd("list").
        WithDescription("Command for listing available tests").
        WithExample("List available tests", "test list").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
