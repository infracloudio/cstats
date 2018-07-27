package src

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func GetStats(name string) {
	cli, err := client.NewEnvClient()
	checkError(err)

	args := filters.NewArgs()
	args.Add("name", name)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Filters: args})
	checkError(err)

	for _, container := range containers {
		stats, err := cli.ContainerStats(context.Background(), container.ID, true)
		checkError(err)
		io.Copy(os.Stdout, stats.Body)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
