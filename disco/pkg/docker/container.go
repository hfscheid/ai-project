package docker

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
)

// RunContainer starts a new container with the informed docker image and name, and returns the container's ID if successful
func (c *Controller) RunContainer(ctx context.Context, info ContainerInfo) (string, error) {
    dockerImage := info.BaseImage + ":" + info.ImageVersion
	err := c.EnsureImage(ctx, dockerImage)
	if err != nil {
        return "", err
	}

	resp, err := c.cli.ContainerCreate(
        ctx,
        &container.Config{
            Image: dockerImage,
            Tty:   false,
        },
        &container.HostConfig{
            Privileged: false,
            CapAdd: []string{"CAP_NET_ADMIN", "CAP_NET_RAW", "CAP_SYS_ADMIN"},
            Mounts: []mount.Mount{
                {
                    Type: mount.TypeBind,
                    Source: info.VolumeSource,
                    Target: info.VolumeTarget,
                    ReadOnly: false,
                }, 
            },
        },
        nil, nil, info.ContainerName)
	if err != nil {
        return "", err
	}

	if err := c.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        return "", err
	}

    statusCode, err := c.ContainerWait(ctx, resp.ID)
    if err != nil {
        return "", fmt.Errorf("%d: %s\n", statusCode, err.Error())
    }

    return resp.ID, nil
}

func (c *Controller) RemoveContainer(ctx context.Context, id string) error {
    err := c.cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
    if err != nil {
       return fmt.Errorf("Unable to remove container %q: %q\n", id, err)
    }
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