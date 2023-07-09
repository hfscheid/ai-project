package cmd

import (
	"context"
	"errors"
	"fmt"
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
        return fmt.Errorf("No test selected, run 'disco test select <test_name>'")
    }

    var wg sync.WaitGroup
    nwInfo := dockerTranslateNw(test.Network)
    if nwId, err := d.dockerC.CreateNetwork(ctx, nwInfo);
    err != nil {
        return fmt.Errorf("Network %s: %q", nwId, err)
    }

    cntInfos := dockerTranslateContainers(test.Containers, nwInfo.NetworkName)
    wg.Add(len(cntInfos))
    errCh := make(chan error)
    errA := []error{}
    go func(errCh chan error, errA []error) {
        for {
            err, ok := <- errCh
            if ok {
                errA = append(errA, err)
            } else {
                return
            }
        }
    }(errCh, errA)
    for _, c := range cntInfos {
        go func(ctx context.Context, c *docker.ContainerInfo,
                wg *sync.WaitGroup, errCh chan error) {
            id, err := d.dockerC.RunContainer(ctx, *c, watch)
            if err != nil {
                errCh <- fmt.Errorf("%v: %q", id, err)
            }
            wg.Done()
        }(ctx, &c, &wg, errCh)
    }
    wg.Wait()
    close(errCh)
    return errors.Join(errA...)
}

func dockerTranslateNw(nw *config.Network) docker.NetworkInfo {
    return docker.NetworkInfo{
        NetworkName: nw.Name,
        Subnet: nw.Subnet,
        Gateway: nw.Gateway,
    }
}

func dockerTranslateContainers(cs []*config.Container, nwName string) []docker.ContainerInfo {
    cInfos := make([]docker.ContainerInfo, len(cs))
    for i, c := range(cs) {
        vols := []docker.VolumeInfo{}
        for _, p := range c.ConfigPaths {
            dirs := strings.Split(p, ":")
            vols = append(vols, docker.VolumeInfo{
                VolumeSource: dirs[0],
                VolumeTarget: dirs[1],
            })
        }
        cInfo := docker.ContainerInfo {
            ContainerName: c.Name,
            NetworkName: nwName,
            BaseImage: c.Image.Name,
            ImageVersion: c.Image.Version,
            Volumes: vols,
            ContainerIp: c.IP,
            ExposePort: fmt.Sprintf("%v",c.ExposedPort),
        }
        cInfos[i] = cInfo
    }
    return cInfos
}
