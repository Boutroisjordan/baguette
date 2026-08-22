package config

type Config struct {
	Id         int        `json:"id"`
	Target     Target     `json:"target"`
	ClientType ClientType `json:"clientType"`
	SocketPath string     `json:"socketPath"`
	Selected   bool       `json:"selected"`
}

type Target int

const (
	Linux Target = iota
	Darwin
	Windows
)

var targetName = map[Target]string{
	Linux:   "linux",
	Darwin:  "macos",
	Windows: "windows",
}

func (t Target) String() string {
	return targetName[t]
}

type ClientType int

const (
	DockerDesktop ClientType = iota
	OrbStack
	PodmanDesktop
	RancherDesktop
	Engine
)

var clientTypeName = map[ClientType]string{
	DockerDesktop:  "docker-desktop",
	OrbStack:       "orb-stack",
	PodmanDesktop:  "podman-desktop",
	RancherDesktop: "rancher-desktop",
	Engine:         "docker-engine",
}

func (ct ClientType) String() string {
	return clientTypeName[ct]
}

func GetAllProviders() []string {
	var providers []string

	for i := 0; i < len(clientTypeName); i++ {
		providers = append(providers, clientTypeName[ClientType(i)])
	}

	return providers
}
