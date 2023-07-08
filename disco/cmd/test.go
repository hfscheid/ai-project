package cmd

import (
	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/spf13/cobra"
)

// newTestCmd defines the 'test' command for Disco CLI
func (d *Disco) newTestCmd() *cobra.Command {
    return cmd.NewCmd("test").
        WithDescription("Command for managing tests").
        WithExample("Creates a test", "test create").
        AddSubCommand(d.newTestCreateCmd()).
        AddSubCommand(d.newTestDeleteCmd()).
        AddSubCommand(d.newTestListCmd()).
        AddSubCommand(d.newTestRunCmd()).
        AddSubCommand(d.newTestSelectCmd()).
        AddSubCommand(d.newTestStopCmd()).
        AddSubCommand(d.newTestDescribeCmd()).
        Super()
}

// func (gb *Global) routerStartErr(_ context.Context, cmd *cobra.Command) error {
//     // Create container configs and start it
//     absPath, err := filepath.Abs("../")
//     if err != nil {
//         return fmt.Errorf("Failed to get valid absolute path: %q", err)
//     }
//     configdir, err := cmd.Flags().GetString("config-dir")
//     if err != nil {
//         return fmt.Errorf("Failed to get config directory path: %q", err)
//     }
//     routerType, err := cmd.Flags().GetString("router")
//     if err != nil {
//         return fmt.Errorf("Failed to get router type: %q", err)
//     }
//     containerInfo, err := buildContainerInfo(routerType, absPath+configdir)
//     if err != nil {
//         return fmt.Errorf("Failed configure container: %q", err)
//     }
//     cID, err := gb.dockerClient.RunContainer(global.ctx, containerInfo)
//     if err != nil {
//         return fmt.Errorf("Failed to start container: %q", err)
//     }
//     global.containerPool = append(global.containerPool, cID)
//     return nil
// }
// 
// func buildContainerInfo(routerType, configpath string) (docker.ContainerInfo, error) {
//         switch routerType {
//         case "frr":
//             name := fmt.Sprintf("frr-%v", global.count["frr"])
//             global.count["frr"]++
//             return docker.ContainerInfo {
//                 ContainerName: name,
//                 BaseImage: "quay.io/frrouting/frr",
//                 ImageVersion: "8.5.1",
//                 VolumeSource: configpath,
//                 VolumeTarget: "/etc/frr",
//             }, nil
//         case "bird":
//             name := fmt.Sprintf("bird-%v", global.count["bird"])
//             global.count["bird"]++
//             return docker.ContainerInfo{
//                 ContainerName: name,
//                 BaseImage: "",
//                 ImageVersion: "",
//                 VolumeSource: configpath,
//                 VolumeTarget: "",
//             }, nil
//         default:
//             return docker.ContainerInfo{},  fmt.Errorf("Could not recognize router type.")
//         }
// }
