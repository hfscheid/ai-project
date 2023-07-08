package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestSelectCmd defines the 'test select' command for Disco CLI
func (d *Disco) newTestSelectCmd() *cobra.Command {
    return cmd.NewCmd("select").
        WithDescription("Command for selecting a test").
        WithExample("Select a test", "test select").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
