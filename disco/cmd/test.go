package cmd

import (
	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestCmd defines the 'test' command for Disco CLI
func (d *Disco) newTestCmd() *cobra.Command {
    return cmd.NewCmd("test").
        WithDescription("Command for managing tests").
        WithExample("Creates a test", "test create").
        AddSubCommand(d.newTestCreateCmd()).
        AddSubCommand(d.newTestDeleteCmd()).
        AddSubCommand(d.newTestListCmd()).
        AddSubCommand(d.newTestRunCmd()).
        AddSubCommand(d.newTestSelectCmd()).
        AddSubCommand(d.newTestStopCmd()).
        AddSubCommand(d.newTestDescribeCmd()).
        Super()
}
