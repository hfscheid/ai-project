package main

import (
    "context"
    "fmt"
    "path/filepath"

    "github.com/hfscheid/ai-project/automated-testing/docker"
)


func main() {
    ctx := context.Background()
    // Connect to Docker client
    client, err := docker.NewController()
    if err != nil {
        fmt.Printf("Failed to initialize Docker client: %q", err)
        return
    }
    defer func() {
        if err := client.Shutdown(); err != nil {
            fmt.Printf("Failed to close Docker connection: %q", err)
            return
        }
    }()

    // Create container configs and start it
    absPath, err := filepath.Abs("../")
    if err != nil {
        fmt.Printf("Failed to get valid absolute path: %q", err)
        return
    }
    absPath += "/routers/frr/confs"
    containerInfo := docker.ContainerInfo{
        ContainerName: "frr",
        BaseImage: "quay.io/frrouting/frr",
        ImageVersion: "8.5.1",
        VolumeSource: absPath,
        VolumeTarget: "/etc/frr",
    }
    cID, err := client.RunContainer(ctx, containerInfo)
    if err != nil {
        fmt.Printf("Failed to start container: %q", err)
        return
    }
    
    // See container logs
    cLogs, err := client.GetContainerLogs(ctx, cID)
    if err != nil {
        fmt.Printf("Failed to get container %q logs: %q", cID, err)
        return
    }
    fmt.Println(cLogs)

    // Shutdown container
    if err := client.RemoveContainer(ctx, cID); err != nil {
        fmt.Printf("Failed to close Docker connection: %q", err)
        return
    }
}
