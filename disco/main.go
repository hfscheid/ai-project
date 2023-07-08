package main

import (
    "fmt"

    "github.com/hfscheid/ai-project/disco/cmd"
)

// To run the cli properly, run `go install` from within the /disco directory
func main() {
    disco, err := cmd.CreateDisco()
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := disco.Execute(); err != nil {
        fmt.Println(err)
    }
}

// func _main() {
//     ctx := context.Background()
//     // Create container configs and start it
//     absPath, err := filepath.Abs("../")
//     if err != nil {
//         fmt.Printf("Failed to get valid absolute path: %q", err)
//         return
//     }
//     absPath += "/routers/frr/confs"
//     containerInfo := docker.ContainerInfo{
//         ContainerName: "frr",
//         BaseImage: "quay.io/frrouting/frr",
//         ImageVersion: "8.5.1",
//         VolumeSource: absPath,
//         VolumeTarget: "/etc/frr",
//     }
//     cID, err := client.RunContainer(ctx, containerInfo)
//     if err != nil {
//         fmt.Printf("Failed to start container: %q", err)
//         return
//     }
//     
//     // See container logs
//     cLogs, err := client.GetContainerLogs(ctx, cID)
//     if err != nil {
//         fmt.Printf("Failed to get container %q logs: %q", cID, err)
//         return
//     }
//     fmt.Println(cLogs)
// 
//     // Shutdown container
//     if err := client.RemoveContainer(ctx, cID); err != nil {
//         fmt.Printf("Failed to close Docker connection: %q", err)
//         return
//     }
// }
