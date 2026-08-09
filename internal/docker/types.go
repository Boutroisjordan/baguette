package docker

type Container struct {
	Id     string
	Name   string
	Image  string
	Up     bool
	Status string
}
