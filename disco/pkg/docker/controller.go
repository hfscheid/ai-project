package docker

import "github.com/docker/docker/client"

type Controller struct {
    cli             *client.Client
    containerPool          map[string]ContainerInfo
    nwPool   map[string]NetworkInfo
}


func NewController() (*Controller, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
        return nil, err
	}

    return &Controller{cli: cli}, nil
}

func (c *Controller) Shutdown() error {
    return c.cli.Close()
}
