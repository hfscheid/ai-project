// based on: https://github.com/b-nova-techhub/jamctl
package cmd

import (
    "github.com/spf13/cobra"
// might be interesting for frr and bird configs
//    "github.com/spf13/viper"
)

var (
    tstrouters = &cobra.Command {
        Use: "tstrouters",
        Short: "tstrouters - tool for creating, configuring and testing software routers",
        Version: "0.1.0"
        SilenceErrors: true,
        SilenceUsage: true,
    }
)

func Execute() error {
    return tstrouters.Execute()
}
