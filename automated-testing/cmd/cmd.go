// based on: https://github.com/b-nova-techhub/jamctl
package cmd

import (
    "github.com/spf13/cobra"
// might be interesting for frr and bird configs
//    "github.com/spf13/viper"
)

var (
    automatedTesting = &cobra.Command {
        Use: "automated-testing",
        Short: "automated-testing - tool for creating, configuring and testing software routers",
        Version: "0.1.0",
        SilenceErrors: true,
        SilenceUsage: true,
    }
)

func Execute() error {
    return automatedTesting.Execute()
}

func init() {
    automatedTesting.AddCommand(createcmd)
}
