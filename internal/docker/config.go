package docker

import (
	"os/user"
	"runtime"

	"github.com/moby/moby/client"
)

func GetClient() (*client.Client, error) {
	path, err := getDockerSocketFromOs(runtime.GOOS)
	if err != nil {
		panic(err)
	}

	conn, err := Connect(path)
	defer CloseConnection(conn)
	if err != nil {
		panic(err)
	}

	return conn, nil
}

func getDockerSocketFromOs(os string) (string, error) {
	var path string

	switch os {
	case "linux":
		path = "/var/run/docker.sock"
	case "windows":
		panic("not yet supported")
	case "darwin":
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		path = u.HomeDir + "/.docker/run/docker.sock"
	default:
		panic("No os matched")
	}

	return path, nil
}
