package loadbench

import (
	"baguette/internal/docker"
	"bufio"
	"fmt"
	"os"

	"github.com/moby/moby/client"
)

func Load() {
	conn, err := docker.GetClient()

	if err != nil {
		panic("Failed to get docker client")
	}

	ChooseContainer(conn)
}

// Todo: Create PickAContainer and call ListContainers + Display + Input
func ChooseContainer(conn *client.Client) {
	var containerId string

	containers := docker.ListContainers(conn)
	docker.DisplayContainerList(containers)

	fmt.Println("Choose container")
	fmt.Printf("> ")
	if _, err := fmt.Scan(&containerId); err != nil {
		panic("Something went wrong")
		return
	}

	fmt.Printf("You choose %s\n", containerId)
}

func MeasureMetrics() {
	panic("Not implemented yet")
}

func DisplayReport() {
	panic("Not implemented yet")
}

// OBSELETE for now
func getDockerSocketFromIO() {

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	var inputPath string

	messages := []string{
		"Something went wrong whilst loading docker socket",
		"Could you enter the path of your docker.sock please?",
	}

	for _, message := range messages {
		fmt.Println(message)
	}
	fmt.Printf("> ")
	_, err := fmt.Scan(&inputPath)
	if err != nil {
		return
	}
	fmt.Println(inputPath)
}
