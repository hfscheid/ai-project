// based on: https://github.com/b-nova-techhub/jamctl
package cmd

import (
	"github.com/hfscheid/ai-project/automated-testing/pkg/cmd"
	"github.com/hfscheid/ai-project/automated-testing/pkg/docker"
	"github.com/spf13/cobra"
    "context"
	// might be interesting for frr and bird configs
	// "github.com/spf13/viper"
)


var global struct {
    dockerClient *docker.Controller
    ctx context.Context
    count map[string]int
    containerPool []string
}

func initGlobal(client *docker.Controller) {
    global.dockerClient = client
    global.ctx = context.Background()
    global.count["frr"] = 0
    global.count["bird"] = 0
}

func NewMainCmd(version string, client *docker.Controller) *cobra.Command {
    initGlobal(client)
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
