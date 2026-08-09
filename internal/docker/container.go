package docker

import (
	"baguette/tools/displayTools"
	"context"
	"fmt"
	"strings"

	"github.com/moby/moby/client"
)

// ListContainers get List of container from the moby/client
func ListContainers(conn *client.Client) []Container {

	opts := client.ContainerListOptions{
		All: true,
	}
	containers, err := conn.ContainerList(context.Background(), opts)

	if err != nil {
		panic(err)
	}

	result := mappingItemsToContainer(containers)
	return result
}

func DisplayContainerList(containers []Container) {

	fmt.Printf("%-10s %-30s %-20s %-10s %-20s\n", "ID", "IMAGE", "NAME", "UP", "STATUS")
	fmt.Println(strings.Repeat("-", 100))

	for _, container := range containers {
		fmt.Printf("%-10s %-30q %-20s %-10t %-20s \n", displayTools.TruncateString(container.Id, 12, false), displayTools.TruncateString(container.Name, 30, true), displayTools.TruncateString(container.Image, 20, true), container.Up, displayTools.TruncateString(container.Status, 20, true))
	}
}

func mappingItemsToContainer(result client.ContainerListResult) []Container {

	var containers []Container

	for _, item := range result.Items {
		container := Container{
			Id:     item.ID,
			Name:   item.Names[0],
			Image:  item.Image,
			Up:     item.State == "running",
			Status: item.Status,
		}

		containers = append(containers, container)
	}

	return containers
}
