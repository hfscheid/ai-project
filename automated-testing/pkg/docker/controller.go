package docker

import "github.com/docker/docker/client"

type Controller struct {
    cli *client.Client
}

// baseImage, imageVersion, containerName string
type ContainerInfo struct {
    BaseImage string
    ImageVersion string
    ContainerName string
    VolumeSource string
    VolumeTarget string
}

func NewController() (*Controller, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
        return nil, err
	}

    return &Controller{cli}, nil
}

func (c *Controller) Shutdown() error {
    return c.cli.Close()
}
