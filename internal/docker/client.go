package docker

import (
	"context"
	"log"

	"github.com/moby/moby/client"
)

// Connect provides a connection from Docker Daemon API
func Connect(path string) (*client.Client, error) {

	host := "unix://" + path
	conn, err := client.New(client.WithHost(host))

	if err != nil {
		log.Fatal("Failed to create Docker client: ", err.Error())
		return nil, err
	}

	opts := client.PingOptions{}

	if _, err = conn.Ping(context.Background(), opts); err != nil {
		_ = conn.Close()
		log.Fatal("Failed to create Docker client", err.Error())
		return nil, err
	}

	return conn, nil
}

// CloseConnection should be used for Close moby/client
func CloseConnection(client *client.Client) {
	_ = client.Close()
}
