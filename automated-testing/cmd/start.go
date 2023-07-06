package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/automated-testing/pkg/cmd"
	"github.com/spf13/cobra"
)

// NewStartCmd defines the create command for Disco CLI
func NewStartCmd() *cobra.Command {
    startRouter := cmd.NewCmd("router").
        WithDescription("Start a new router in a container").
        WithExample("Start FRR router", "router --router frr --config-dir /this/dir").
        WithCommonFlags().
        NoArgs(routerStartErr)

    return cmd.NewCmd("start").
        WithDescription("Starts structures in containers").
        WithExample("Starts a router", "start router --router frr --config-dir /this/dir").
        AddSubCommand(startRouter).
        Super()
}

func routerStartErr(_ context.Context) error {
    return fmt.Errorf("Expected no arguments, but received some")
}
