package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestRunCmd defines the 'test run' command for Disco CLI
func (d *Disco) newTestRunCmd() *cobra.Command {
    return cmd.NewCmd("run").
        WithDescription("Command for running selected test").
        WithExample("Run selected test", "test run").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
