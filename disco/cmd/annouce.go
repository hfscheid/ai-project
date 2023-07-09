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

// newAnnouceCmd defines the 'annouce' command for Disco CLI
func (d *Disco) newAnnouceCmd() *cobra.Command {
    return cmd.NewCmd("annouce").
        WithDescription("Command for annoucing BGP routes").
        WithExample("Command syntax", "annouce <exabgp_container_name> <annouce_command>").
        WithExample("Annoucing a route", "annouce exabgp1 'annouce route 100.10.0.0/24 next-hop self'").
        ExactArgs(2, d.sendAnnoucement)
}

func (d *Disco) sendAnnoucement(ctx context.Context, args []string) error {
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

    _, err := http.PostForm(
        fmt.Sprintf("http://127.0.0.1:%d", container.ExposedPort),
        url.Values{ "command": []string{ annoucement } },
    )
    return err
}
