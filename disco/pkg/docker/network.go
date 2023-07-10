package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	// "github.com/docker/docker/client"
)

// Information required to create the network,
// plus its ID after creation
type NetworkInfo struct {
    NetworkName string
    Subnet      string
    Gateway     string
    ID          string
}

func (c *Controller) GetNetworkId(networkName string) (string, error) {
    nwInfo, ok := c.nwPool[networkName]
    if !ok {
        return "", fmt.Errorf("Could not find network %v", networkName)
    }
    return nwInfo.ID, nil
}

func (c *Controller) CreateNetwork(ctx context.Context, info NetworkInfo) (string, error) {
    if info.NetworkName == "bridge" || info.NetworkName == "host" || info.NetworkName == "none" {
        return "", fmt.Errorf("Network can't be named 'bridge', 'host' or 'none'")
    }
    resp, err := c.cli.NetworkCreate(
        ctx,
        fmt.Sprintf("disco-%s", info.NetworkName),
        types.NetworkCreate{},
    )
    if err != nil {
        return "", err
    }
    info.ID = resp.ID
    c.nwPool[info.NetworkName] = info
    return resp.ID, nil
}

func (c *Controller) RemoveNetwork(ctx context.Context, networkName string) error {
    if _, ok := c.nwPool[networkName]; !ok {
        return nil
    }
    id := c.nwPool[networkName].ID
    err := c.cli.NetworkRemove(ctx, id)
    if err != nil {
       return fmt.Errorf("Unable to remove network %q: %q\n", id, err)
    }
    delete(c.nwPool, networkName)
    return nil
}
