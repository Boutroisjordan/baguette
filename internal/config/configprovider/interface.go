package configprovider

import "baguette/internal/config"

type ConfigProvider interface {
	Create(client config.ClientType, target config.Target, socketPath string) (*config.Config, error)
	GetSocketPath(target string) (string, error)
}
