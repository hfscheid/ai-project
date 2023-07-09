package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Controller struct {
    cli             *client.Client
    containerPool   map[string]ContainerInfo
    nwPool          map[string]NetworkInfo
}


func NewController(ctx context.Context) (*Controller, error) {
    controller := Controller{}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
        return nil, err
	}
    controller.cli = cli
    controller.containerPool = map[string]ContainerInfo{}
    controller.nwPool = map[string]NetworkInfo{}

    err = controller.RegisterContainers(ctx)
    if err != nil {
        return nil, err
    }
    err = controller.RegisterNetwork(ctx)
    if err != nil {
        return nil, err
    }

    return &controller, nil
}

func (c *Controller) Shutdown() error {
    return c.cli.Close()
}

func (c *Controller) RegisterNetwork(ctx context.Context) error {
    netList, err := c.cli.NetworkList(ctx, types.NetworkListOptions{})
    if err != nil {
        return fmt.Errorf("Unable to list networks: %v", err)
    }
    for _, net := range netList {
        if strings.HasPrefix(net.Name, "disco-") {
            if len(net.IPAM.Config) != 1 {
                return fmt.Errorf("Error in network %q: currently only 1 IPAM config is supported", net.Name)
            }
            nInfo := NetworkInfo{
                ID: net.ID,
                NetworkName: net.Name,
                Subnet: net.IPAM.Config[0].Subnet,
                Gateway: net.IPAM.Config[0].Gateway,
            }
            c.nwPool[net.Name] = nInfo
        }
    }
    return nil
}

func (c *Controller) RegisterContainers(ctx context.Context) error {
    contList, err := c.cli.ContainerList(ctx, types.ContainerListOptions{ All: true })
    if err != nil {
        return fmt.Errorf("Unable to list containers: %v", err)
    }
    for _, cont := range contList {
        for _, name := range cont.Names {
            if strings.HasPrefix(name, "disco-") {
                var netName string
                imageData := strings.Split(cont.Image, ":")
                if len(imageData) != 2 {
                    return fmt.Errorf("Invalid image, must be '<image name>:<version>'")
                }
                if len(cont.Mounts) == 0 {
                    return fmt.Errorf("No volumes defined for container %q", name)
                }
                vols := []VolumeInfo{}
                for _, vol := range cont.Mounts {
                    vols = append(vols, VolumeInfo{
                        VolumeSource: vol.Source,
                        VolumeTarget: vol.Destination,
                    })
                }
                if len(cont.NetworkSettings.Networks) != 1 {
                    return fmt.Errorf("Error in container %q: currently only 1 network is supported", name)
                }
                for key := range cont.NetworkSettings.Networks {
                    if strings.HasPrefix(key, "disco-") {
                        netName = key
                        break
                    }
                }
                cInfo := ContainerInfo{
                    ID: cont.ID,
                    BaseImage: imageData[0],
                    ImageVersion: imageData[1],
                    ContainerName: name,
                    Volumes: vols,
                    NetworkName: netName,
                }
                c.containerPool[name] = cInfo
            }
        }
    }
    return nil
}
