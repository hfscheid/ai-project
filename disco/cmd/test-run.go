package cmd

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/hfscheid/ai-project/disco/pkg/docker"
	"github.com/spf13/cobra"
)

// newTestRunCmd defines the 'test run' command for Disco CLI
func (d *Disco) newTestRunCmd() *cobra.Command {
    flag := cmd.Flag {
        Name:           "watch",
        Shorthand:      "w",
        Usage:          "disco test-run --watch",
        Value:          &cmd.DiscoOptions.Watch,
        DefValue:       false,
        FlagAddMethod:  "BoolVar",
        DefinedOn:      []string{"run"},
    }
    return cmd.NewCmd("run").
        WithDescription("Command for running selected test").
        WithExample("Run selected test", "test run").
        WithFlags(&flag).
        NoArgs(d.runTest)
}

func (d *Disco) runTest(ctx context.Context, c *cobra.Command) error {
    watch := cmd.DiscoOptions.Watch
    test := d.selectedTest
    if test == nil {
        fmt.Println("No test selected, run 'disco test select <test_name>'")
        return nil
    }

    var wg sync.WaitGroup
    nwInfo := dockerTranslateNw(test.Network)
    if nwId, err := d.dockerC.CreateNetwork(ctx, nwInfo);
    err != nil {
        fmt.Printf("Network %s: %q\n", nwId, err)
        return nil
    }

    cntInfos, err := dockerTranslateContainers(test.Containers, nwInfo.NetworkName)
    if err != nil {
        return err
    }
    wg.Add(len(cntInfos))
    fmt.Println("Starting containers...")
    for _, c := range cntInfos {
        cont := c
        go func(ctx context.Context, cont *docker.ContainerInfo, wg *sync.WaitGroup) {
            _, err := d.dockerC.RunContainer(ctx, *cont, watch)
            if err != nil {
                fmt.Printf("%s: %q\n", cont.ContainerName, err)
            }
            wg.Done()
        }(ctx, &cont, &wg)
    }
    wg.Wait()
    return nil
}

func dockerTranslateNw(nw *config.Network) docker.NetworkInfo {
    return docker.NetworkInfo{
        NetworkName: nw.Name,
        Subnet: nw.Subnet,
        Gateway: nw.Gateway,
    }
}

func dockerTranslateContainers(cs []*config.Container, nwName string) ([]docker.ContainerInfo, error) {
    cInfos := make([]docker.ContainerInfo, len(cs))
    currDir, err := filepath.Abs(".")
    if err != nil {
        return nil, fmt.Errorf("Error getting absolute path to curr dir: %v", err)
    }
    for i, c := range(cs) {
        vols := []docker.VolumeInfo{}
        for _, p := range c.ConfigPaths {
            dirs := strings.Split(p, ":")
            if len(dirs) != 2 {
                return nil, fmt.Errorf("Config paths must be in the format '<host path>:<container path>'")
            }
            vols = append(vols, docker.VolumeInfo{
                VolumeSource: filepath.Join(currDir, dirs[0]),
                VolumeTarget: dirs[1],
            })
        }
        exposedPort := ""
        if c.ExposedPort > 0 {
            exposedPort = fmt.Sprintf("%v",c.ExposedPort)
        }
        cInfo := docker.ContainerInfo {
            ContainerName: c.Name,
            NetworkName: nwName,
            BaseImage: c.Image.Name,
            ImageVersion: c.Image.Version,
            Volumes: vols,
            ContainerIp: c.IP,
            ExposePort: exposedPort,
        }
        cInfos[i] = cInfo
    }
    return cInfos, nil
}
