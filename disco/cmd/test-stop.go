package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestStopCmd defines the 'test stop' command for Disco CLI
func (d *Disco) newTestStopCmd() *cobra.Command {
    return cmd.NewCmd("stop").
        WithDescription("Command to stop running selected test").
        WithExample("Stop selected test", "test stop").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
