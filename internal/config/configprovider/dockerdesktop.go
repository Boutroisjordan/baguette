package configprovider

import (
	"baguette/internal/config"
	"errors"
	"fmt"
)

type DockerDesktopConfig struct {
	Config config.Config
}

func (dc *DockerDesktopConfig) Create(client config.ClientType, target config.Target, socketPath string) (*config.Config, error) {
	config := dc.Config
	config.Target = target
	config.ClientType = client
	config.SocketPath = socketPath
	config.Selected = false
	return &config, nil
}

func (dc *DockerDesktopConfig) GetSocketPath(target string) (string, error) {
	var path string
	errMsg := "%s isn't supported \n"

	switch target {
	case "linux":
	case "darwin":
		path = "~/.docker/run/docker.sock"
	case "windows":
		msg := fmt.Sprintf(errMsg, target)
		return "", errors.New(msg)
	}

	return path, nil
}
