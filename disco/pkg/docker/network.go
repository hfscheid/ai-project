package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"
)

func (c *Controller) GetNetworkId(networkName string) (string, error) {
    id, ok := c.nwPool[networkName]
    if !ok {
        return "", fmt.Errorf("Could not find network %v", networkName)
    }
    return id, nil
}

func (c *Controller) CreateNetwork(ctx context.Context, networkName string) error {
    resp, err := c.cli.NetworkCreate(
        ctx,
        networkName,
        types.NetworkCreate{},
    )
    if err != nil {
        return err
    }
    c.nwPool[networkName] = resp.ID
    return nil
}
