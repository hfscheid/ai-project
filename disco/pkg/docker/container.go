package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

type VolumeInfo struct {
    VolumeSource    string
    VolumeTarget    string
}

// Information required to build and run the container,
// plus its ID after creation
type ContainerInfo struct {
    BaseImage       string
    ImageVersion    string
    ContainerName   string
    ContainerIp     string
    Volumes         []VolumeInfo
    ExposePort      string
    NetworkName     string
    ID              string
}

// RunContainer starts a new container with the informed docker image and name, and returns the container's ID if successful
func (c *Controller) RunContainer(ctx context.Context, info ContainerInfo, watch bool) (string, error) {
    containerID := ""
    containerName := ""
    if cont, ok := c.containerPool[fmt.Sprintf("/disco-%s", info.ContainerName)]; ok {
        containerID = cont.ID
        containerName = cont.ContainerName
    } else {
        dockerImage := info.BaseImage + ":" + info.ImageVersion
        networkId, err := c.GetNetworkId(info.NetworkName)
        endpt := &network.EndpointSettings{
            NetworkID: networkId,
            IPAddress: info.ContainerIp,
        }
        if err != nil {
            return "", err
        }
        err = c.EnsureImage(ctx, dockerImage)
        if err != nil {
            return "", err
        }

        vols := []mount.Mount{}
        for _, vol := range info.Volumes {
            vols = append(vols, mount.Mount{
                Type: mount.TypeBind,
                Source: vol.VolumeSource,
                Target: vol.VolumeTarget,
                ReadOnly: false,
            })
        }

        containerName = fmt.Sprintf("disco-%s", info.ContainerName)
        containerCfg := &container.Config{
            Image: dockerImage,
            Tty:   false,
        }
        hostCfg := &container.HostConfig{
            Privileged: false,
            CapAdd: []string{"CAP_NET_ADMIN", "CAP_NET_RAW", "CAP_SYS_ADMIN"},
            Mounts: vols,
        }
        if info.ExposePort != "" {
            port := nat.Port(info.ExposePort)  
            containerCfg.ExposedPorts = nat.PortSet{
                port: struct{}{},
            }
            hostCfg.PortBindings = nat.PortMap{
                port: []nat.PortBinding{
                    {
                        HostIP: "0.0.0.0",
                        HostPort: info.ExposePort,
                    },
                },
            }
        }
        resp, err := c.cli.ContainerCreate(
            ctx,
            containerCfg,
            hostCfg,
            &network.NetworkingConfig{
                EndpointsConfig: map[string]*network.EndpointSettings{
                    fmt.Sprintf("disco-%s", info.NetworkName): endpt,
                },
            }, nil, containerName,
        )
        if err != nil {
            return "", err
        }
        containerID = resp.ID
        info.ID = containerID
        c.containerPool[info.ContainerName] = info
    }
    tst := strings.NewReader(fmt.Sprintf("Running container [%s] %s\n", containerID, containerName))
    _, _ = io.Copy(os.Stdout, tst)
    if err := c.cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
        return "", err
    }

    if watch {
        statusCode, err := c.ContainerWait(ctx, containerID)
        if err != nil {
            return "", fmt.Errorf("%d: %s\n", statusCode, err.Error())
        }
        logString, err := c.GetContainerLogs(ctx, containerID)
        if err != nil {
            return "", fmt.Errorf("Failed to get container logs for container %s: %q\n",
            info.ContainerName,
            err.Error())
        }
        logStream := strings.NewReader(fmt.Sprintf("[%s]: %s\n", info.ContainerName, logString)) 
        _, _ = io.Copy(os.Stdout, logStream)
    }

    return containerID, nil
}

func (c *Controller) StopContainer(ctx context.Context, containerName string) error {
    if _, ok := c.containerPool[containerName]; !ok {
        return fmt.Errorf("Unable to stop container %q: container not found", containerName)
    }
    id := c.containerPool[containerName].ID
    err := c.cli.ContainerStop(ctx, id, container.StopOptions{})
    if err != nil {
       return fmt.Errorf("Unable to stop container %q: %q\n", id, err)
    }
    return nil
}

func (c *Controller) RemoveContainer(ctx context.Context, containerName string) error {
    if _, ok := c.containerPool[containerName]; !ok {
        return nil
    }
    id := c.containerPool[containerName].ID
    err := c.cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
    if err != nil {
       return fmt.Errorf("Unable to remove container %q: %q\n", id, err)
    }
    delete(c.containerPool, containerName)
    return nil
}

// ContainerWait waits until the given container has been stopped
func (c *Controller) ContainerWait(ctx context.Context, id string) (int64, error) {
	statusCh, errCh := c.cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        return 0, err
    case result := <-statusCh:
        return result.StatusCode, nil
    }
}

func (c *Controller) GetContainerLogs(ctx context.Context, id string) (string, error) {
    logCtx, cancel := context.WithTimeout(ctx, time.Second * 10)
    defer cancel()

	out, err := c.cli.ContainerLogs(logCtx, id, types.ContainerLogsOptions{
        ShowStdout: true,
        ShowStderr: true,
    })
	if err != nil {
        return "", err
	}

    buffer, err := io.ReadAll(out)
    return string(buffer), err
}
