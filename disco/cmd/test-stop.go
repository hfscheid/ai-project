package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestStopCmd defines the 'test stop' command for Disco CLI
func (d *Disco) newTestStopCmd() *cobra.Command {
    return cmd.NewCmd("stop").
        WithDescription("Command to stop running selected test").
        WithExample("Stop selected test", "test stop").
        NoArgs(d.stopTest)
}

func (d *Disco) stopTest (ctx context.Context, c *cobra.Command) error {
    // get all active containers and network from current test suite
    currTest := d.selectedTest
    // docker stop all
    errs := []error{}
    for _, container := range currTest.Containers {
        err := d.dockerC.StopContainer(ctx, container.Name)
        if err != nil {
            errs = append(errs, err)
        }
    }
    if err := errors.Join(errs...);
    err != nil {
        return fmt.Errorf("Unable to stop containers: %v", err)
    }
    return nil
}
