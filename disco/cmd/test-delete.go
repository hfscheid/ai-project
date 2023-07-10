package cmd

import (
	"context"
	"errors"
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
        fmt.Println("No test selected, run 'disco test select <test_name>'")
        return nil
    }

    delete(d.tests.TestCases, currTest.Name)
    d.selectedTest = nil
    err := config.WriteToConfigFile(d.tests)
    if err != nil {
        fmt.Printf("Failed to remove test from config file: %v\n", err)
        return nil
    }
    errs := []error{}
    for _, container := range currTest.Containers {
        contName := fmt.Sprintf("/disco-%s", container.Name)
        err = d.dockerC.RemoveContainer(ctx, contName)
        if err != nil {
            errs = append(errs, err)
        }
    }
    if err := errors.Join(errs...);
    err != nil {
        fmt.Printf("Unable to remove containers: %v\n", err)
        return nil
    }
    netName := fmt.Sprintf("disco-%s", currTest.Network.Name)
    err = d.dockerC.RemoveNetwork(ctx, netName)
    if err != nil {
        fmt.Printf("Unable to remove network: %v\n", err)
    }
    return nil
}
