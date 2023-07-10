package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/spf13/cobra"
)

// newTestSelectCmd defines the 'test select' command for Disco CLI
func (d *Disco) newTestSelectCmd() *cobra.Command {
    return cmd.NewCmd("select").
        WithDescription("Command for selecting a test").
        WithExample("Select a test", "test select <test name>").
        ExactArgs(1, d.selectTest)
}

func (d *Disco) selectTest(ctx context.Context, s []string) error {
    testName := s[0]
    if _, ok := d.tests.TestCases[testName]; !ok {
        fmt.Printf("Unable to find test %q\n", testName)
        return nil
    }
    d.tests.SelectedTest = d.tests.TestCases[testName]
    err := config.WriteToConfigFile(d.tests)
    if err != nil {
        fmt.Println(err)
    }
    return nil
}
