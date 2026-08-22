package pathResolver

import (
	"baguette/internal/config/configprovider"
	"errors"
	"fmt"
)

// Todo replace by GetFactory
// Send facotry and then use getPath
func GetPathFromTargetAndClientType(target string, clientType writer.ClientType) (string, error) {

	var socketPath string
	var err error
	conf := writer.Config{}
	var provider writer.ConfigProvider

	switch clientType.String() {
	case "docker-desktop":
		provider = &configprovider.DockerDesktopConfig{
			Config: conf,
		}
		socketPath, err = provider.GetSocketPath(target)
	case "rancher-desktop", "podman-desktop", "docker-engine", "orb-stack":
		err = errors.New("client not supported yet")
	default:
		msg := fmt.Sprintf("unknown client type %s", clientType.String())
		err = errors.New(msg)
	}

	if err != nil {
		return "", err
	}

	return socketPath, nil
}
