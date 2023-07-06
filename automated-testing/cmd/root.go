// based on: https://github.com/b-nova-techhub/jamctl
package cmd

import (
    "github.com/spf13/cobra"
    "github.com/hfscheid/ai-project/automated-testing/pkg/cmd"
// might be interesting for frr and bird configs
//    "github.com/spf13/viper"
)

func NewMainCmd(version string) *cobra.Command {
    rootCmd := cmd.NewCmd("disco").
        Version(version).
        WithDescription("disco - tool for creating, configuring and testing software routers").
        WithLongDescription("disco - tool for creating, configuring and testing software routers. Use 'disco help' to list all available commands").
        AddSubCommand(
            NewStartCmd(),
        ).
        Super()

    return rootCmd
}
