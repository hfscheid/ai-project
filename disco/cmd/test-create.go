package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/spf13/cobra"
)

// newTestCreateCmd defines the 'test create' command for Disco CLI
func (d *Disco) newTestCreateCmd() *cobra.Command {
    return cmd.NewCmd("create").
        WithDescription("Command for creating tests").
        WithExample("Creates a test", "test create /path/to/test.yaml").
        ExactArgs(1, d.createTest)
}

func (d *Disco) createTest(ctx context.Context, args []string) error {
    currDir, err := filepath.Abs(".")
    if err != nil {
        fmt.Printf("Error getting absolute path to curr dir: %v\n", err)
        return nil
    }
    cfgFile := filepath.Join(currDir, args[0])
    testCase, err := config.ReadTestConfig(cfgFile)
    if err != nil {
        return err
    }
    if _, ok := d.tests.TestCases[testCase.Name]; !ok {
        d.tests.TestCases[testCase.Name] = testCase
        d.tests.SelectedTest = d.tests.TestCases[testCase.Name] 
        err := config.WriteToConfigFile(d.tests)
        if err != nil {
            fmt.Println(err)
        }
        return nil
    }
    fmt.Println("Test with the same name already exists, must be unique")
    return nil
}
