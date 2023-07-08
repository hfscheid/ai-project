package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/spf13/cobra"
)

// newTestDeleteCmd defines the 'test delete' command for Disco CLI
func (d *Disco) newTestDeleteCmd() *cobra.Command {
    return cmd.NewCmd("delete").
        WithDescription("Command to delete selected test").
        WithExample("delete selected test", "test delete").
        NoArgs(d.deleteTest)
}

func (d *Disco) deleteTest(ctx context.Context, c *cobra.Command) error {
    currTest := d.selectedTest
    if currTest == nil {
        return fmt.Errorf("No test selected, run 'disco test select <test_name>'")
    }

    delete(d.tests.TestCases, currTest.Name)
    d.selectedTest = nil
    err := config.WriteToConfigFile(d.tests)
    if err != nil {
        return fmt.Errorf("Failed to remove test from config file: %v", err)
    }

    err = d.dockerC.RemoveContainer(ctx, currTest.Name)
    if err != nil {
        return fmt.Errorf("Unable to remove container: %v", err)
    }
    return nil
}
