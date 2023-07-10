package cmd

import (
	"github.com/hfscheid/ai-project/disco/pkg/cmd"
)

func (d *Disco) createCLI() {
    rootCmd := cmd.NewCmd("disco").
        Version("0.0.0").
        WithDescription("disco - tool for creating, configuring and testing software routers").
        WithLongDescription("disco - tool for creating, configuring and testing software routers. Use 'disco help' to list all available commands").
        AddSubCommand(
            d.newTestCmd(),
            d.newAnnounceCmd(),
        ).
        Super()
    d.cli = rootCmd
}
