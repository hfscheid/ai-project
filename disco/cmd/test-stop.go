package cmd

import (
	"context"
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

func (d *Disco) stopTest(ctx context.Context, c *cobra.Command) error {
    // get all active containers and network from current test suite
    currTest := d.selectedTest
    // docker stop all
    fmt.Println("Stopping containers...")
    for _, container := range currTest.Containers {
        contName := fmt.Sprintf("/disco-%s", container.Name)
        err := d.dockerC.StopContainer(ctx, contName)
        if err != nil {
            fmt.Printf("Error stopping container %s: %v\n", contName, err)
            continue
        }
        fmt.Printf("%s stopped\n", contName)
    }
    return nil
}
