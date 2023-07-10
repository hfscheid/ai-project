package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/spf13/cobra"
)

// newAnnounceCmd defines the 'annouce' command for Disco CLI
func (d *Disco) newAnnounceCmd() *cobra.Command {
    return cmd.NewCmd("announce").
        WithDescription("Command for announcing BGP routes").
        WithExample("Command syntax", "announce <exabgp_container_name> <announce_command>").
        WithExample("Announcing a route", "announce exabgp1 'announce route 100.10.0.0/24 next-hop self'").
        ExactArgs(2, d.sendAnnouncement)
}

func (d *Disco) sendAnnouncement(ctx context.Context, args []string) error {
    currTest := d.selectedTest
    if currTest == nil {
        return fmt.Errorf("No test selected, run 'disco test select <test_name>'")
    }

    containerName := args[0]
    annoucement := args[1]
    var container *config.Container
    for _, c := range currTest.Containers {
        if c.Name == containerName {
            container = c
            break
        }
    }
    if container == nil {
        return fmt.Errorf("Unable to find container %q in current test", containerName)
    }
    if container.ExposedPort == 0 {
        return fmt.Errorf("Container %q doesn`t have exposed ports", containerName)
    }

    _, err := http.PostForm(
        fmt.Sprintf("http://127.0.0.1:%d", container.ExposedPort),
        url.Values{ "command": []string{ annoucement } },
    )
    return err
}
