package cmd

import (
	"context"
	"fmt"

	"github.com/hfscheid/ai-project/disco/pkg/cmd"
	"github.com/hfscheid/ai-project/disco/pkg/config"
	"github.com/spf13/cobra"
)

const testCaseTemplate =
`Name: %s
Network:
%s
Containers:
%s
`

const networkTemplate =
`   Name: %s
    Subnet: %s
    Gateway: %s
`

const containerTemplate =
`   Name: %s
    Type: %s
    ConfigPath: %s
    IP: %s
`

// newTestDescribeCmd defines the 'test describe' command for Disco CLI
func (d *Disco) newTestDescribeCmd() *cobra.Command {
    return cmd.NewCmd("describe").
        WithDescription("Command to describe selected test info").
        WithExample("Describe currently selected test info", "test describe").
        NoArgs(d.describeTest)
}

func (d *Disco) describeTest(_ context.Context, c *cobra.Command) error {
    currTest := d.selectedTest
    if currTest == nil {
        return fmt.Errorf("No test selected, run 'disco test select <test_name>'")
    }
    netDesc := generateNetworkDescription(currTest.Network)
    containerDesc := ""
    for _, c := range currTest.Containers {
        containerDesc += generateContainerDescription(c)
    }
    testDesc := fmt.Sprintf(
        testCaseTemplate,
        currTest.Name,
        netDesc,
        containerDesc,
    )
    fmt.Println(testDesc)
    return nil
}

func generateNetworkDescription(net *config.Network) string {
    if net == nil {
        return fmt.Sprintf(networkTemplate, "-", "-", "-")
    }
    return fmt.Sprintf(
        networkTemplate,
        net.Name,
        net.Subnet,
        net.Gateway,
    )
}

func generateContainerDescription(c *config.Container) string {
    if c == nil {
        return fmt.Sprintf(containerTemplate, "-", "-", "-", "-")
    }
    return fmt.Sprintf(
        containerTemplate,
        c.Name,
        c.Type.String(),
        c.ConfigPath,
        c.IP,
    )
}
