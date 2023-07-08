package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestDeleteCmd defines the 'test delete' command for Disco CLI
func (d *Disco) newTestDeleteCmd() *cobra.Command {
    return cmd.NewCmd("delete").
        WithDescription("Command to delete selected test").
        WithExample("delete selected test", "test delete").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
