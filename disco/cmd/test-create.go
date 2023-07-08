package cmd

import (
	"context"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestCreateCmd defines the 'test create' command for Disco CLI
func (d *Disco) newTestCreateCmd() *cobra.Command {
    return cmd.NewCmd("create").
        WithDescription("Command for creating tests").
        WithExample("Creates a test", "test create").
        NoArgs(func(ctx context.Context, c *cobra.Command) error { return nil })
}
