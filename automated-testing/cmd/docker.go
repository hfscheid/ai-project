package cmd

import (
//    "context"
//    "fmt"

    "github.com/hfscheid/ai-project/automated-testing/pkg/cmd"
    "github.com/spf13/cobra"
)

// NewStartCmd defines the create command for Disco CLI
func NewDockerCmd() *cobra.Command {
    dockerOn := cmd.NewCmd("on").
        WithDescription("Initialize Docker client").
        WithExample("Initialize Docker client", "on").
        NoArgs(routerStartErr)

    dockerOff := cmd.NewCmd("off").
        WithDescription("Terminate Docker client").
        WithExample("Terminate Docker client", "off").
        NoArgs(routerStartErr)

    return cmd.NewCmd("docker").
        WithDescription("Initializes/terminates Docker the Docker client").
        WithExample("Initialize Docker client", "docker on").
        AddSubCommand(dockerOn).
        AddSubCommand(dockerOff).
        Super()
}
