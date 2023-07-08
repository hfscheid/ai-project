package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
)

func (c *Controller) EnsureImage(ctx context.Context, image string) error {
	reader, err := c.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
        return err
	}

	defer reader.Close()
    _, err = io.Copy(os.Stdout, reader)
    return err
}
