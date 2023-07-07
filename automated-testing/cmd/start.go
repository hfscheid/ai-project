package cmd

import (
	"context"
	"fmt"
    "path/filepath"
	"github.com/hfscheid/ai-project/automated-testing/pkg/cmd"
	"github.com/hfscheid/ai-project/automated-testing/pkg/docker"
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

func routerStartErr(_ context.Context, cmd *cobra.Command) error {
    // Create container configs and start it
    absPath, err := filepath.Abs("../")
    if err != nil {
        return fmt.Errorf("Failed to get valid absolute path: %q", err)
    }
    configdir, err := cmd.Flags().GetString("config-dir")
    if err != nil {
        return fmt.Errorf("Failed to get config directory path: %q", err)
    }
    routerType, err := cmd.Flags().GetString("router")
    if err != nil {
        return fmt.Errorf("Failed to get router type: %q", err)
    }
    containerInfo, err := buildContainerInfo(routerType, absPath+configdir)
    if err != nil {
        return fmt.Errorf("Failed configure container: %q", err)
    }
    cID, err := global.dockerClient.RunContainer(global.ctx, containerInfo)
    if err != nil {
        return fmt.Errorf("Failed to start container: %q", err)
    }
    global.containerPool = append(global.containerPool, cID)
    return nil
}

func buildContainerInfo(routerType, configpath string) (docker.ContainerInfo, error) {
        switch routerType {
        case "frr":
            name := fmt.Sprintf("frr-%v", global.count["frr"])
            global.count["frr"]++
            return docker.ContainerInfo {
                ContainerName: name,
                BaseImage: "quay.io/frrouting/frr",
                ImageVersion: "8.5.1",
                VolumeSource: configpath,
                VolumeTarget: "/etc/frr",
            }, nil
        case "bird":
            name := fmt.Sprintf("bird-%v", global.count["bird"])
            global.count["bird"]++
            return docker.ContainerInfo{
                ContainerName: name,
                BaseImage: "",
                ImageVersion: "",
                VolumeSource: configpath,
                VolumeTarget: "",
            }, nil
        default:
            return docker.ContainerInfo{},  fmt.Errorf("Could not recognize router type.")
        }
}
