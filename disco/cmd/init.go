package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/hfscheid/ai-project/disco/pkg/docker"
	"github.com/spf13/cobra"
)

type Disco struct {
    dockerC *docker.Controller
    cli *cobra.Command
    tests *config.Tests
    selectedTest *config.TestCase
}

func CreateDisco(ctx context.Context) (*Disco, error) {
    disco := &Disco{}
    dockerC, err := docker.NewController(ctx)
    if err != nil {
        return nil, fmt.Errorf("Failed to start Docker client: %v", err)
    }
    disco.dockerC = dockerC
    
    tests, err := config.ReadConfigFile()
    if err != nil {
        return nil, fmt.Errorf("Failed to read/create config file: %v", err)
    }
    disco.tests = tests
    disco.selectedTest = disco.tests.SelectedTest
    disco.createCLI()    
    return disco, nil
}

func (d *Disco) Execute() error {
    if err := d.cli.Execute(); err != nil {
        return fmt.Errorf("Error while running disco: %v", err)
    }
    return nil
}
