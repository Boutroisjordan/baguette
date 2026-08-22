package configprovider

import (
	"baguette/internal/config"
	"fmt"
)

type DockerEngineConfig struct {
	Config config.Config
}

func (dc *DockerEngineConfig) Create(client config.ClientType, target config.Target, socketPath string) (*config.Config, error) {
	config := dc.Config
	config.Target = target
	config.ClientType = client
	config.SocketPath = socketPath
	config.Selected = false
	return &config, nil
}

func (dc *DockerEngineConfig) GetSocketPath(target string) (string, error) {
	var path string
	errMsg := "Docker engine isn't supported on %s \n"

	switch target {
	case "linux":
		path = "/var/run/docker.sock"
	case "darwin":
		message := fmt.Sprintf(errMsg, target)
		return "", fmt.Errorf(message)
	case "windows":
		message := fmt.Sprintf(errMsg, target)
		return "", fmt.Errorf(message)
	}

	return path, nil
}
